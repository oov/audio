package wave

import (
	"bytes"
	"encoding/binary"
	"github.com/oov/audio"
	"io"
	"io/ioutil"
	"os"
)

type Writer struct {
	w       io.Writer
	wfext   WaveFormatExtensible
	aw      audio.InterleavedWriter
	body    io.Writer
	head    int64
	written int64
}

func NewWriter(w io.Writer, wfext *WaveFormatExtensible) (*Writer, error) {
	if ws, ok := w.(io.WriteSeeker); ok {
		return newDirectWriter(ws, wfext)
	}

	if wr, err := newTempFileWriter(w, wfext); err == nil {
		return wr, err
	}
	return newTempMemWriter(w, wfext)
}

func newDirectWriter(ws io.WriteSeeker, wfext *WaveFormatExtensible) (*Writer, error) {
	conv, err := wfext.Format.InterleavedFormatConverter()
	if err != nil {
		return nil, err
	}

	head, err := ws.Seek(0, os.SEEK_CUR)
	if err != nil {
		return nil, err
	}

	// insert header margin
	// "RIFF" size "WAVE" "fmt " size body "data" size
	_, err = ws.Write(make([]byte, 4+4+4+4+4+wfext.Size()+4+4))
	if err != nil {
		return nil, err
	}

	return &Writer{
		w:     ws,
		wfext: *wfext,
		aw:    audio.NewInterleavedWriter(conv, ws),
		head:  head,
	}, nil
}

func newTempFileWriter(w io.Writer, wfext *WaveFormatExtensible) (*Writer, error) {
	conv, err := wfext.Format.InterleavedFormatConverter()
	if err != nil {
		return nil, err
	}

	tempfile, err := ioutil.TempFile("", "tempwave")
	if err != nil {
		return nil, err
	}

	return &Writer{
		w:     w,
		wfext: *wfext,
		aw:    audio.NewInterleavedWriter(conv, tempfile),
		body:  tempfile,
	}, nil
}

func newTempMemWriter(w io.Writer, wfext *WaveFormatExtensible) (*Writer, error) {
	conv, err := wfext.Format.InterleavedFormatConverter()
	if err != nil {
		return nil, err
	}

	tempbuf := bytes.NewBufferString("")
	return &Writer{
		w:     w,
		wfext: *wfext,
		aw:    audio.NewInterleavedWriter(conv, tempbuf),
		body:  tempbuf,
	}, nil
}

func (w *Writer) WriteFloat32Interleaved(p [][]float32) (n int, err error) {
	n, err = w.aw.WriteFloat32Interleaved(p)
	w.written += int64(n)
	return
}

func (w *Writer) WriteFloat64Interleaved(p [][]float64) (n int, err error) {
	n, err = w.aw.WriteFloat64Interleaved(p)
	w.written += int64(n)
	return
}

func (w *Writer) Close() error {
	var err error

	ws, isWriteSeeker := w.w.(io.WriteSeeker)
	if isWriteSeeker {
		if _, err = ws.Seek(w.head, os.SEEK_SET); err != nil {
			return err
		}
	}

	dataSize := w.written * int64(w.wfext.Format.BlockAlign)
	// "RIFF" + sz + "WAVE" + "fmt " + sz + fmtbody + "data" + sz + databody
	fileSize := 4 + 4 + 4 + 4 + 4 + int64(w.wfext.Size()) + 4 + 4 + dataSize

	_, err = w.w.Write([]byte("RIFF"))
	if err != nil {
		return err
	}

	err = binary.Write(w.w, binary.LittleEndian, int32(fileSize))
	if err != nil {
		return err
	}

	_, err = w.w.Write([]byte("WAVE"))
	if err != nil {
		return err
	}

	// write "fmt " chunk

	_, err = w.w.Write([]byte("fmt "))
	if err != nil {
		return err
	}

	err = binary.Write(w.w, binary.LittleEndian, int32(w.wfext.Size()))
	if err != nil {
		return err
	}

	_, err = w.wfext.WriteTo(w.w)
	if err != nil {
		return err
	}

	// write "data" chunk

	_, err = w.w.Write([]byte("data"))
	if err != nil {
		return err
	}

	err = binary.Write(w.w, binary.LittleEndian, int32(dataSize))
	if err != nil {
		return err
	}

	if isWriteSeeker {
		// already written
		return nil
	}

	switch t := w.body.(type) {
	case *os.File:
		_, err = t.Seek(0, os.SEEK_SET)
		if err != nil {
			return err
		}

		_, err = io.Copy(w.w, t)
		if err != nil {
			return err
		}

		err = t.Close()
		if err != nil {
			return err
		}
		os.Remove(t.Name())

	case *bytes.Buffer:
		_, err = w.w.Write(t.Bytes())
		if err != nil {
			return err
		}
		t.Reset()
	}
	return nil
}

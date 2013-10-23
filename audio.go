package audio

import (
	"github.com/oov/audio/converter"
	"io"
)

type Reader interface {
	ReadFloat32(p []float32) (n int, err error)
	ReadFloat64(p []float64) (n int, err error)
}

type Writer interface {
	WriteFloat32(p []float32) (n int, err error)
	WriteFloat64(p []float64) (n int, err error)
}

type InterleavedReader interface {
	ReadFloat32Interleaved(p [][]float32) (n int, err error)
	ReadFloat64Interleaved(p [][]float64) (n int, err error)
}

type InterleavedWriter interface {
	WriteFloat32Interleaved(p [][]float32) (n int, err error)
	WriteFloat64Interleaved(p [][]float64) (n int, err error)
}

type reader struct {
	conv converter.FormatConverter
	r    io.Reader
	buf  []byte
}

func NewReader(conv converter.FormatConverter, r io.Reader) Reader {
	return &reader{
		conv: conv,
		r:    r,
	}
}

func (r *reader) ReadFloat32(p []float32) (n int, err error) {
	ln := len(p)
	if ln > len(r.buf) {
		r.buf = make([]byte, ln)
	}

	n, err = r.r.Read(r.buf[:ln])
	r.conv.ToFloat32(r.buf[:n], p)
	return
}

func (r *reader) ReadFloat64(p []float64) (n int, err error) {
	ln := len(p)
	if ln > len(r.buf) {
		r.buf = make([]byte, ln)
	}

	n, err = r.r.Read(r.buf[:ln])
	r.conv.ToFloat64(r.buf[:n], p)
	return
}

type interleavedReader struct {
	conv converter.InterleavedFormatConverter
	r    io.Reader
	buf  []byte
}

func NewInterleavedReader(conv converter.InterleavedFormatConverter, r io.Reader) InterleavedReader {
	return &interleavedReader{
		conv: conv,
		r:    r,
	}
}

func (r *interleavedReader) ReadFloat32Interleaved(p [][]float32) (n int, err error) {
	ln := len(p) * len(p[0])
	if ln > len(r.buf) {
		r.buf = make([]byte, ln)
	}

	n, err = r.r.Read(r.buf[:ln])
	r.conv.ToFloat32Interleaved(r.buf[:n], p)
	return
}

func (r *interleavedReader) ReadFloat64Interleaved(p [][]float64) (n int, err error) {
	ln := len(p) * len(p[0])
	if ln > len(r.buf) {
		r.buf = make([]byte, ln)
	}

	n, err = r.r.Read(r.buf[:ln])
	r.conv.ToFloat64Interleaved(r.buf[:n], p)
	return
}

type writer struct {
	w    io.Writer
	buf  []byte
	conv converter.FormatConverter
}

func NewWriter(conv converter.FormatConverter, w io.Writer) Writer {
	return Writer(&writer{w: w, conv: conv})
}

func (w *writer) WriteFloat32(p []float32) (n int, err error) {
	ln := len(p)
	if ln > len(w.buf) {
		w.buf = make([]byte, ln)
	}

	w.conv.FromFloat32(p, w.buf[:ln])
	n, err = w.w.Write(w.buf[:ln*w.conv.Size()])
	return
}

func (w *writer) WriteFloat64(p []float64) (n int, err error) {
	ln := len(p)
	if ln > len(w.buf) {
		w.buf = make([]byte, ln)
	}

	w.conv.FromFloat64(p, w.buf[:ln])
	n, err = w.w.Write(w.buf[:ln*w.conv.Size()])
	return
}

type interleavedWriter struct {
	w    io.Writer
	buf  []byte
	conv converter.InterleavedFormatConverter
}

func NewInterleavedWriter(conv converter.InterleavedFormatConverter, w io.Writer) InterleavedWriter {
	return &interleavedWriter{
		w:    w,
		conv: conv,
	}
}

func (w *interleavedWriter) WriteFloat32Interleaved(p [][]float32) (n int, err error) {
	ln := len(p) * len(p[0])
	if ln > len(w.buf) {
		w.buf = make([]byte, ln)
	}

	w.conv.FromFloat32Interleaved(p, w.buf[:ln])
	n, err = w.w.Write(w.buf[:ln*w.conv.Size()])
	return
}

func (w *interleavedWriter) WriteFloat64Interleaved(p [][]float64) (n int, err error) {
	ln := len(p) * len(p[0])
	if ln > len(w.buf) {
		w.buf = make([]byte, ln)
	}

	w.conv.FromFloat64Interleaved(p, w.buf[:ln])
	n, err = w.w.Write(w.buf[:ln*w.conv.Size()])
	return
}
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
	conv converter.Converter
	r    io.Reader
	buf  []byte
}

func NewReader(conv converter.Converter, r io.Reader) Reader {
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
	conv converter.InterleavedConverter
	r    io.Reader
	buf  []byte
}

func NewInterleavedReader(conv converter.InterleavedConverter, r io.Reader) InterleavedReader {
	return &interleavedReader{
		conv: conv,
		r:    r,
	}
}

func (r *interleavedReader) ReadFloat32Interleaved(p [][]float32) (n int, err error) {
	ln := len(p) * len(p[0]) * r.conv.SampleSize()
	if ln > len(r.buf) {
		r.buf = make([]byte, ln)
	}

	n, err = r.r.Read(r.buf[:ln])
	r.conv.ToFloat32Interleaved(r.buf[:n], p)
	n /= len(p) * r.conv.SampleSize()
	return
}

func (r *interleavedReader) ReadFloat64Interleaved(p [][]float64) (n int, err error) {
	ln := len(p) * len(p[0]) * r.conv.SampleSize()
	if ln > len(r.buf) {
		r.buf = make([]byte, ln)
	}

	n, err = r.r.Read(r.buf[:ln])
	r.conv.ToFloat64Interleaved(r.buf[:n], p)
	n /= len(p) * r.conv.SampleSize()
	return
}

type writer struct {
	w    io.Writer
	buf  []byte
	conv converter.Converter
}

func NewWriter(conv converter.Converter, w io.Writer) Writer {
	return Writer(&writer{w: w, conv: conv})
}

func (w *writer) WriteFloat32(p []float32) (n int, err error) {
	ln := len(p)
	if ln > len(w.buf) {
		w.buf = make([]byte, ln)
	}

	w.conv.FromFloat32(p, w.buf[:ln])
	n, err = w.w.Write(w.buf[:ln*w.conv.SampleSize()])
	return
}

func (w *writer) WriteFloat64(p []float64) (n int, err error) {
	ln := len(p)
	if ln > len(w.buf) {
		w.buf = make([]byte, ln)
	}

	w.conv.FromFloat64(p, w.buf[:ln])
	n, err = w.w.Write(w.buf[:ln*w.conv.SampleSize()])
	return
}

type interleavedWriter struct {
	w    io.Writer
	buf  []byte
	conv converter.InterleavedConverter
}

func NewInterleavedWriter(conv converter.InterleavedConverter, w io.Writer) InterleavedWriter {
	return &interleavedWriter{
		w:    w,
		conv: conv,
	}
}

func (w *interleavedWriter) WriteFloat32Interleaved(p [][]float32) (n int, err error) {
	ln := len(p) * len(p[0]) * w.conv.SampleSize()
	if ln > len(w.buf) {
		w.buf = make([]byte, ln)
	}

	w.conv.FromFloat32Interleaved(p, w.buf[:ln])
	n, err = w.w.Write(w.buf[:ln])
	n /= len(p) * w.conv.SampleSize()
	return
}

func (w *interleavedWriter) WriteFloat64Interleaved(p [][]float64) (n int, err error) {
	ln := len(p) * len(p[0]) * w.conv.SampleSize()
	if ln > len(w.buf) {
		w.buf = make([]byte, ln)
	}

	w.conv.FromFloat64Interleaved(p, w.buf[:ln])
	n, err = w.w.Write(w.buf[:ln])
	n /= len(p) * w.conv.SampleSize()
	return
}

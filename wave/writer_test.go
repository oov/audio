package wave

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

var wfext = &WaveFormatExtensible{
	Format: WaveFormatEx{
		FormatTag:      WAVE_FORMAT_PCM,
		Channels:       2,
		SamplesPerSec:  48000,
		AvgBytesPerSec: 96000,
		BlockAlign:     4,
		BitsPerSample:  16,
		ExtSize:        0,
	},
}
var samples = [][]float64{
	[]float64{-1, 0, 1},
	[]float64{1, 0, -1},
}
var golden = []byte("RIFF\x38\x00\x00\x00WAVEfmt \x10\x00\x00\x00\x01\x00\x02\x00\x80\xbb\x00\x00\x00\x77\x01\x00\x04\x00\x10\x00data\x0c\x00\x00\x00\x01\x80\xff\x7f\x00\x00\x00\x00\xff\x7f\x01\x80")

func TestDirectWriter(t *testing.T) {
	f, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Error(err)
		return
	}

	defer f.Close()
	defer os.Remove(f.Name())

	w, err := newDirectWriter(f, wfext)
	if err != nil {
		t.Error(err)
		return
	}

	n, err := w.WriteFloat64Interleaved(samples)
	if err != nil {
		t.Error(err)
		return
	}
	if n != 3 {
		t.Log("invalid written size:", n)
		t.Fail()
		return
	}

	err = w.Close()
	if err != nil {
		t.Error(err)
		return
	}

	b, err := ioutil.ReadFile(f.Name())
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(golden, b) {
		t.Log("golden:", golden)
		t.Log("invalid output:", b)
		t.Fail()
		return
	}
}

func TestTempFileWriter(t *testing.T) {
	buf := bytes.NewBufferString("")
	w, err := newTempFileWriter(buf, wfext)
	if err != nil {
		t.Error(err)
		return
	}

	n, err := w.WriteFloat64Interleaved(samples)
	if err != nil {
		t.Error(err)
		return
	}

	if n != 3 {
		t.Log("invalid written size:", n)
		t.Fail()
		return
	}

	err = w.Close()
	if err != nil {
		t.Error(err)
		return
	}
	b := buf.Bytes()
	if !bytes.Equal(golden, b) {
		t.Log("golden:", golden)
		t.Log("invalid output:", b)
		t.Fail()
		return
	}
}

func TestTempMemWriter(t *testing.T) {
	buf := bytes.NewBufferString("")
	w, err := newTempMemWriter(buf, wfext)
	if err != nil {
		t.Error(err)
		return
	}

	n, err := w.WriteFloat64Interleaved(samples)
	if err != nil {
		t.Error(err)
		return
	}

	if n != 3 {
		t.Log("invalid written size:", n)
		t.Fail()
		return
	}

	err = w.Close()
	if err != nil {
		t.Error(err)
		return
	}

	b := buf.Bytes()
	if !bytes.Equal(golden, b) {
		t.Log("golden:", golden)
		t.Log("invalid output:", b)
		t.Fail()
		return
	}
}

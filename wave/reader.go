// Package wave provides support for reading and wrting RIFF Waveform Audio Format File.
package wave

import (
	"encoding/binary"
	"errors"
	"github.com/oov/audio"
	"io"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// NewReader returns an audio.InterleavedReader which waveform audio data from r.
func NewReader(r io.Reader) (audio.InterleavedReader, *WaveFormatExtensible, error) {
	var chunk [4]byte
	var err error
	if _, err = io.ReadFull(r, chunk[:]); err != nil {
		return nil, nil, err
	}
	if string(chunk[:4]) != "RIFF" {
		return nil, nil, errors.New("wave: invalid header")
	}

	var ln int32
	if err = binary.Read(r, binary.LittleEndian, &ln); err != nil {
		return nil, nil, err
	}

	lr := io.LimitReader(r, int64(ln))

	if _, err = io.ReadFull(lr, chunk[:]); err != nil {
		return nil, nil, err
	}
	if string(chunk[:4]) != "WAVE" {
		return nil, nil, errors.New("wave: invalid header")
	}

	var n int
	var ignored [4096]byte

	// find "fmt " chunk

	ln = 0
	for string(chunk[:4]) != "fmt " {
		n, err = io.ReadFull(lr, ignored[:min(len(ignored), int(ln))])
		if err != nil {
			return nil, nil, err
		}
		ln -= int32(n)
		if ln != 0 {
			continue
		}

		if _, err = io.ReadFull(lr, chunk[:]); err != nil {
			return nil, nil, err
		}
		if err = binary.Read(r, binary.LittleEndian, &ln); err != nil {
			return nil, nil, err
		}
	}

	if ln < 16 {
		return nil, nil, errors.New("wave: fmt chunk too small")
	}

	// read "fmt " chunk

	var wf WaveFormatExtensible
	var rd int64
	if rd, err = wf.Format.ReadFrom(r); err != nil {
		return nil, nil, err
	}

	// ignore unsupported chunk data

	ln -= int32(rd)
	for ln > 0 {
		n, err = io.ReadFull(lr, ignored[:min(len(ignored), int(ln))])
		if err != nil {
			return nil, nil, err
		}
		ln -= int32(n)
	}

	// find "data" chunk

	ln = 0
	for string(chunk[:4]) != "data" {
		n, err = io.ReadFull(lr, ignored[:min(len(ignored), int(ln))])
		if err != nil {
			return nil, nil, err
		}
		ln -= int32(n)
		if ln != 0 {
			continue
		}

		if _, err = io.ReadFull(lr, chunk[:]); err != nil {
			return nil, nil, err
		}
		if err = binary.Read(r, binary.LittleEndian, &ln); err != nil {
			return nil, nil, err
		}
	}

	conv, err := wf.Format.InterleavedConverter()
	if err != nil {
		return nil, nil, err
	}
	return audio.NewInterleavedReader(conv, io.LimitReader(r, int64(ln))), &wf, nil
}

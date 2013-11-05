package wave

import (
	"math"
	"os"
	"testing"
)

type TestFile struct {
	filename string
	wf       WaveFormatExtensible
}

var testfiles = []TestFile{
	TestFile{
		filename: "48kHz1ch8bit.wav",
		wf: WaveFormatExtensible{
			Format: WaveFormatEx{
				FormatTag:      WAVE_FORMAT_PCM,
				Channels:       1,
				SamplesPerSec:  48000,
				AvgBytesPerSec: 48000,
				BlockAlign:     1,
				BitsPerSample:  8,
			},
		},
	},
	TestFile{
		filename: "48kHz1ch16bit.wav",
		wf: WaveFormatExtensible{
			Format: WaveFormatEx{
				FormatTag:      WAVE_FORMAT_PCM,
				Channels:       1,
				SamplesPerSec:  48000,
				AvgBytesPerSec: 48000 * 2,
				BlockAlign:     2,
				BitsPerSample:  16,
			},
		},
	},
	TestFile{
		filename: "48kHz1ch24bit.wav",
		wf: WaveFormatExtensible{
			Format: WaveFormatEx{
				FormatTag:      WAVE_FORMAT_PCM,
				Channels:       1,
				SamplesPerSec:  48000,
				AvgBytesPerSec: 48000 * 3,
				BlockAlign:     3,
				BitsPerSample:  24,
			},
		},
	},
	TestFile{
		filename: "48kHz1ch32bit.wav",
		wf: WaveFormatExtensible{
			Format: WaveFormatEx{
				FormatTag:      WAVE_FORMAT_PCM,
				Channels:       1,
				SamplesPerSec:  48000,
				AvgBytesPerSec: 48000 * 4,
				BlockAlign:     4,
				BitsPerSample:  32,
			},
		},
	},
	TestFile{
		filename: "48kHz1ch32bitFloat.wav",
		wf: WaveFormatExtensible{
			Format: WaveFormatEx{
				FormatTag:      WAVE_FORMAT_IEEE_FLOAT,
				Channels:       1,
				SamplesPerSec:  48000,
				AvgBytesPerSec: 48000 * 4,
				BlockAlign:     4,
				BitsPerSample:  32,
			},
		},
	},
	TestFile{
		filename: "48kHz1ch64bitFloat.wav",
		wf: WaveFormatExtensible{
			Format: WaveFormatEx{
				FormatTag:      WAVE_FORMAT_IEEE_FLOAT,
				Channels:       1,
				SamplesPerSec:  48000,
				AvgBytesPerSec: 48000 * 8,
				BlockAlign:     8,
				BitsPerSample:  64,
			},
		},
	},
	TestFile{
		filename: "48kHz2ch16bit.wav",
		wf: WaveFormatExtensible{
			Format: WaveFormatEx{
				FormatTag:      WAVE_FORMAT_PCM,
				Channels:       2,
				SamplesPerSec:  48000,
				AvgBytesPerSec: 48000 * 2 * 2,
				BlockAlign:     4,
				BitsPerSample:  16,
			},
		},
	},
}

func isSameWaveFormatEx(a, b *WaveFormatEx) bool {
	if a.FormatTag != b.FormatTag {
		return false
	}
	if a.Channels != b.Channels {
		return false
	}
	if a.SamplesPerSec != b.SamplesPerSec {
		return false
	}
	if a.AvgBytesPerSec != b.AvgBytesPerSec {
		return false
	}
	if a.BlockAlign != b.BlockAlign {
		return false
	}
	if a.BitsPerSample != b.BitsPerSample {
		return false
	}
	if a.ExtSize != b.ExtSize {
		return false
	}
	return true
}

func isValidSamples(p []float64) bool {
	if len(p) != 12 {
		return false
	}
	for _, s := range p[0:4] {
		if math.Abs(s-1.0) > 0.01 {
			return false
		}
	}
	for _, s := range p[4:8] {
		if math.Abs(s) > 0.01 {
			return false
		}
	}
	for _, s := range p[8:12] {
		if math.Abs(s+1.0) > 0.01 {
			return false
		}
	}
	return true
}

func TestReader(t *testing.T) {
	for _, tf := range testfiles {
		f, err := os.Open(tf.filename)
		if err != nil {
			t.Error(err)
			return
		}
		r, wf, err := NewReader(f)
		if err != nil {
			t.Error(err)
			return
		}

		if !isSameWaveFormatEx(&wf.Format, &tf.wf.Format) {
			t.Fail()
			return
		}

		samples := make([][]float64, 0)
		for i := 0; i < int(wf.Format.Channels); i++ {
			samples = append(samples, make([]float64, 12))
		}
		n, err := r.ReadFloat64Interleaved(samples)
		if err != nil {
			t.Error(err)
			return
		}
		if n != 12 {
			t.Fail()
			return
		}
		if !isValidSamples(samples[0]) {
			t.Log("invalid samples on ch1", tf.filename, samples[0])
			t.Fail()
			return
		}
		if wf.Format.Channels > 1 {
			if !isValidSamples(append(append(samples[1][8:12], samples[1][4:8]...), samples[1][0:4]...)) {
				t.Log("invalid samples on ch2", tf.filename, samples[1])
				t.Fail()
				return
			}
		}
	}
}

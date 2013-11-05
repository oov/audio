// +build ignore

//resample.go is a audio resampler.
//
//Usage:
//
//  go run resample.go [options] infile
//
package main

import (
	"flag"
	"fmt"
	"github.com/oov/audio/resampler"
	"github.com/oov/audio/saturator"
	"github.com/oov/audio/wave"
	"io"
	"os"
)

var (
	freq    = flag.Float64("f", 0.8, "Frequency ratio")
	quality = flag.Int("q", 5, "Resampling quality(0-10)")
)

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("resample.go is a audio resampler.")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println()
		fmt.Println("  go run resample.go [options] infile")
		fmt.Println()
		fmt.Println("The options are:")
		fmt.Println()
		flag.PrintDefaults()
		return
	}

	f, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	f2, err := os.Create(flag.Arg(0) + `.out.wav`)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f2.Close()

	ar, wfext, err := wave.NewReader(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	aw, err := wave.NewWriter(f2, wfext)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer aw.Close()

	fmt.Println("Input file:")
	fmt.Println("  Path:", flag.Arg(0))
	fmt.Println("  Samplerate:", wfext.Format.SamplesPerSec)
	fmt.Println("  Channels:", wfext.Format.Channels)
	fmt.Println("  Bits:", wfext.Format.BitsPerSample)
	fmt.Println()
	fmt.Println("Output file:")
	fmt.Println("  Path:", flag.Arg(0)+`.out.wav`)
	fmt.Println("  Samplerate:", wfext.Format.SamplesPerSec)
	fmt.Println("  Channels:", wfext.Format.Channels)
	fmt.Println("  Bits:", wfext.Format.BitsPerSample)
	fmt.Println()

	infreq := int(wfext.Format.SamplesPerSec)
	outfreq := int(float64(wfext.Format.SamplesPerSec) * *freq)
	fmt.Println("resampling:")
	fmt.Printf("  %dHz to %dHz\n", infreq, outfreq)

	inBuf, outBuf := [][]float64{}, [][]float64{}
	for i := 0; i < int(wfext.Format.Channels); i++ {
		inBuf = append(inBuf, make([]float64, wfext.Format.SamplesPerSec))
		outBuf = append(outBuf, make([]float64, wfext.Format.SamplesPerSec))
	}
	in, out := make([][]float64, wfext.Format.Channels), make([][]float64, wfext.Format.Channels)

	rs := resampler.NewWithSkipZeros(int(wfext.Format.Channels), infreq, outfreq, *quality)
	var n int
	var rerr error
	for rerr != io.EOF {
		// read
		n, rerr = ar.ReadFloat64Interleaved(inBuf)
		if rerr != nil && rerr != io.EOF {
			fmt.Println(err)
			return
		}
		for i, inbuf := range inBuf {
			in[i] = inbuf[:n]
		}

		for len(in[0]) > 0 {
			// resample
			for i, inbuf := range in {
				rn, wn := rs.ProcessFloat64(i, inbuf, outBuf[i])
				in[i] = inbuf[rn:]
				out[i] = outBuf[i][:wn]
				saturator.Saturate64Slice(out[i], out[i])
			}

			// write
			_, err = aw.WriteFloat64Interleaved(out)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

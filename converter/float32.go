package converter

import (
	"math"
)

var (
	Float32 Float32Converter
)

type Float32Converter float32

func (c Float32Converter) SampleSize() int {
	return 4
}

func (c Float32Converter) ToFloat32(input []byte, output []float32) {
	for i, o, ln := 0, 0, len(input); i < ln; i += c.SampleSize() {
		output[o] = ByteToFloat32(input[i], input[i+1], input[i+2], input[i+3])
		o++
	}
}

func (c Float32Converter) ToFloat64(input []byte, output []float64) {
	for i, o, ln := 0, 0, len(input); i < ln; i += c.SampleSize() {
		output[o] = Float32ToFloat64(ByteToFloat32(input[i], input[i+1], input[i+2], input[i+3]))
		o++
	}
}

func (c Float32Converter) FromFloat32(input []float32, output []byte) {
	o := 0
	for _, s := range input {
		output[o], output[o+1], output[o+2], output[o+3] = Float32ToByte(s)
		o += c.SampleSize()
	}
}

func (c Float32Converter) FromFloat64(input []float64, output []byte) {
	o := 0
	for _, s := range input {
		output[o], output[o+1], output[o+2], output[o+3] = Float32ToByte(Float64ToFloat32(s))
		o += c.SampleSize()
	}
}

func (c Float32Converter) ToFloat32Interleaved(input []byte, outputs [][]float32) {
	chs := len(outputs)
	for ch, output := range outputs {
		for i, o, ln := ch*c.SampleSize(), 0, len(input); i < ln; i += chs * c.SampleSize() {
			output[o] = ByteToFloat32(input[i], input[i+1], input[i+2], input[i+3])
			o++
		}
	}
}

func (c Float32Converter) ToFloat64Interleaved(input []byte, outputs [][]float64) {
	chs := len(outputs)
	for ch, output := range outputs {
		for i, o, ln := ch*c.SampleSize(), 0, len(input); i < ln; i += chs * c.SampleSize() {
			output[o] = Float32ToFloat64(ByteToFloat32(input[i], input[i+1], input[i+2], input[i+3]))
			o++
		}
	}
}

func (c Float32Converter) FromFloat32Interleaved(inputs [][]float32, output []byte) {
	for i, o := 0, 0; o < len(output); i++ {
		for _, input := range inputs {
			output[o], output[o+1], output[o+2], output[o+3] = Float32ToByte(input[i])
			o += c.SampleSize()
		}
	}
}

func (c Float32Converter) FromFloat64Interleaved(inputs [][]float64, output []byte) {
	for i, o := 0, 0; o < len(output); i++ {
		for _, input := range inputs {
			output[o], output[o+1], output[o+2], output[o+3] = Float32ToByte(Float64ToFloat32(input[i]))
			o += c.SampleSize()
		}
	}
}

func Float32ToFloat64(s float32) float64 {
	return float64(s)
}

func Float64ToFloat32(s float64) float32 {
	return float32(s)
}

func Float32ToFloat64Slice(input []float32, output []float64) {
	for i, s := range input {
		output[i] = Float32ToFloat64(s)
	}
}

func Float64ToFloat32Slice(input []float64, output []float32) {
	for i, s := range input {
		output[i] = Float64ToFloat32(s)
	}
}

func ByteToFloat32(a, b, c, d byte) float32 {
	return math.Float32frombits(uint32(a) | (uint32(b) << 8) | (uint32(c) << 16) | (uint32(d) << 24))
}

func Float32ToByte(s float32) (byte, byte, byte, byte) {
	i := math.Float32bits(s)
	return byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24)
}

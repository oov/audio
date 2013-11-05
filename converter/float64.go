package converter

import (
	"math"
)

var (
	Float64 Float64Converter
)

type Float64Converter float64

func (c Float64Converter) SampleSize() int {
	return 8
}

func (c Float64Converter) ToFloat32(input []byte, output []float32) {
	for i, o, ln := 0, 0, len(input); i < ln; i += c.SampleSize() {
		output[o] = Float64ToFloat32(ByteToFloat64(input[i], input[i+1], input[i+2], input[i+3], input[i+4], input[i+5], input[i+6], input[i+7]))
		o++
	}
}

func (c Float64Converter) ToFloat64(input []byte, output []float64) {
	for i, o, ln := 0, 0, len(input); i < ln; i += c.SampleSize() {
		output[o] = ByteToFloat64(input[i], input[i+1], input[i+2], input[i+3], input[i+4], input[i+5], input[i+6], input[i+7])
		o++
	}
}

func (c Float64Converter) FromFloat32(input []float32, output []byte) {
	o := 0
	for _, s := range input {
		output[o], output[o+1], output[o+2], output[o+3], output[o+4], output[o+5], output[o+6], output[o+7] = Float64ToByte(Float32ToFloat64(s))
		o += c.SampleSize()
	}
}

func (c Float64Converter) FromFloat64(input []float64, output []byte) {
	o := 0
	for _, s := range input {
		output[o], output[o+1], output[o+2], output[o+3], output[o+4], output[o+5], output[o+6], output[o+7] = Float64ToByte(s)
		o += c.SampleSize()
	}
}

func (c Float64Converter) ToFloat32Interleaved(input []byte, outputs [][]float32) {
	chs := len(outputs)
	for ch, output := range outputs {
		for i, o, ln := ch*c.SampleSize(), 0, len(input); i < ln; i += chs * c.SampleSize() {
			output[o] = Float64ToFloat32(ByteToFloat64(input[i], input[i+1], input[i+2], input[i+3], input[i+4], input[i+5], input[i+6], input[i+7]))
			o++
		}
	}
}

func (c Float64Converter) ToFloat64Interleaved(input []byte, outputs [][]float64) {
	chs := len(outputs)
	for ch, output := range outputs {
		for i, o, ln := ch*c.SampleSize(), 0, len(input); i < ln; i += chs * c.SampleSize() {
			output[o] = ByteToFloat64(input[i], input[i+1], input[i+2], input[i+3], input[i+4], input[i+5], input[i+6], input[i+7])
			o++
		}
	}
}

func (c Float64Converter) FromFloat32Interleaved(inputs [][]float32, output []byte) {
	for i, o := 0, 0; o < len(output); i++ {
		for _, input := range inputs {
			output[o], output[o+1], output[o+2], output[o+3], output[o+4], output[o+5], output[o+6], output[o+7] = Float64ToByte(Float32ToFloat64(input[i]))
			o += c.SampleSize()
		}
	}
}

func (c Float64Converter) FromFloat64Interleaved(inputs [][]float64, output []byte) {
	for i, o := 0, 0; o < len(output); i++ {
		for _, input := range inputs {
			output[o], output[o+1], output[o+2], output[o+3], output[o+4], output[o+5], output[o+6], output[o+7] = Float64ToByte(input[i])
			o += c.SampleSize()
		}
	}
}

func ByteToFloat64(a, b, c, d, e, f, g, h byte) float64 {
	return math.Float64frombits(uint64(a) | (uint64(b) << 8) | (uint64(c) << 16) | (uint64(d) << 24) | (uint64(e) << 32) | (uint64(f) << 40) | (uint64(g) << 48) | (uint64(h) << 56))
}

func Float64ToByte(s float64) (byte, byte, byte, byte, byte, byte, byte, byte) {
	i := math.Float64bits(s)
	return byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24), byte(i >> 32), byte(i >> 40), byte(i >> 48), byte(i >> 56)
}

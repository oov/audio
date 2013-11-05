package converter

import (
	"encoding/binary"
	"math"
	"testing"
)

func TestFloat64ConverterToFloat32(t *testing.T) {
	o := make([]float32, dataLen)
	input := make([]byte, dataLen*Float64.SampleSize())
	for i, s := range dataFloat64 {
		binary.LittleEndian.PutUint64(input[i*Float64.SampleSize():], uint64(math.Float64bits(s)))
	}

	Float64.ToFloat32(input, o)
	testResultFloat32(o, t)
}

func TestFloat64ConverterToFloat32Interleaved(t *testing.T) {
	outputs := make([][]float32, dataLen)
	for i := range outputs {
		outputs[i] = make([]float32, 2)
	}

	input := make([]byte, dataLen*2*Float64.SampleSize())
	for i, s := range dataFloat64 {
		binary.LittleEndian.PutUint64(input[i*Float64.SampleSize():], uint64(math.Float64bits(s)))
	}
	for i, s := range dataFloat64 {
		binary.LittleEndian.PutUint64(input[(dataLen+i)*Float64.SampleSize():], uint64(math.Float64bits(s)))
	}

	Float64.ToFloat32Interleaved(input, outputs)
	testResultFloat32([]float32{outputs[0][0], outputs[1][0], outputs[2][0], outputs[3][0], outputs[4][0]}, t)
	testResultFloat32([]float32{outputs[0][1], outputs[1][1], outputs[2][1], outputs[3][1], outputs[4][1]}, t)
}

func TestFloat64ConverterFromFloat32(t *testing.T) {
	output := make([]byte, dataLen*Float64.SampleSize())
	Float64.FromFloat64(dataFloat64, output)
	testResultFloat64([]float64{
		math.Float64frombits(binary.LittleEndian.Uint64(output[0*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[1*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[2*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[3*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[4*Float64.SampleSize():])),
	}, t)
}

func TestFloat64ConverterFromFloat32Interleaved(t *testing.T) {
	inputs := make([][]float32, dataLen)
	for i := range inputs {
		inputs[i] = []float32{dataFloat32[i], dataFloat32[i]}
	}
	output := make([]byte, dataLen*2*Float64.SampleSize())
	Float64.FromFloat32Interleaved(inputs, output)
	testResultFloat64([]float64{
		math.Float64frombits(binary.LittleEndian.Uint64(output[0*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[1*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[2*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[3*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[4*Float64.SampleSize():])),
	}, t)
	testResultFloat64([]float64{
		math.Float64frombits(binary.LittleEndian.Uint64(output[5*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[6*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[7*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[8*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[9*Float64.SampleSize():])),
	}, t)
}

func TestFloat64ConverterToFloat64(t *testing.T) {
	o := make([]float64, dataLen)
	input := make([]byte, dataLen*Float64.SampleSize())
	for i, s := range dataFloat64 {
		binary.LittleEndian.PutUint64(input[i*Float64.SampleSize():], uint64(math.Float64bits(s)))
	}

	Float64.ToFloat64(input, o)
	testResultFloat64(o, t)
}

func TestFloat64ConverterToFloat64Interleaved(t *testing.T) {
	outputs := make([][]float64, dataLen)
	for i := range outputs {
		outputs[i] = make([]float64, 2)
	}

	input := make([]byte, dataLen*2*Float64.SampleSize())
	for i, s := range dataFloat64 {
		binary.LittleEndian.PutUint64(input[i*Float64.SampleSize():], uint64(math.Float64bits(s)))
	}
	for i, s := range dataFloat64 {
		binary.LittleEndian.PutUint64(input[(dataLen+i)*Float64.SampleSize():], uint64(math.Float64bits(s)))
	}

	Float64.ToFloat64Interleaved(input, outputs)
	testResultFloat64([]float64{outputs[0][0], outputs[1][0], outputs[2][0], outputs[3][0], outputs[4][0]}, t)
	testResultFloat64([]float64{outputs[0][1], outputs[1][1], outputs[2][1], outputs[3][1], outputs[4][1]}, t)
}

func TestFloat64ConverterFromFloat64(t *testing.T) {
	output := make([]byte, dataLen*Float64.SampleSize())
	Float64.FromFloat64(dataFloat64, output)
	testResultFloat64([]float64{
		math.Float64frombits(binary.LittleEndian.Uint64(output[0*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[1*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[2*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[3*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[4*Float64.SampleSize():])),
	}, t)
}

func TestFloat64ConverterFromFloat64Interleaved(t *testing.T) {
	inputs := make([][]float64, dataLen)
	for i := range inputs {
		inputs[i] = []float64{dataFloat64[i], dataFloat64[i]}
	}
	output := make([]byte, dataLen*2*Float64.SampleSize())
	Float64.FromFloat64Interleaved(inputs, output)
	testResultFloat64([]float64{
		math.Float64frombits(binary.LittleEndian.Uint64(output[0*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[1*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[2*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[3*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[4*Float64.SampleSize():])),
	}, t)
	testResultFloat64([]float64{
		math.Float64frombits(binary.LittleEndian.Uint64(output[5*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[6*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[7*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[8*Float64.SampleSize():])),
		math.Float64frombits(binary.LittleEndian.Uint64(output[9*Float64.SampleSize():])),
	}, t)
}

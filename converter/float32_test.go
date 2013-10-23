package converter

import (
	"encoding/binary"
	"math"
	"testing"
)

func TestFloat32ToFloat64(t *testing.T) {
	o := make([]float64, dataLen)
	Float32ToFloat64Slice(dataFloat32, o)
	testResultFloat64(o, t)
}

func TestFloat64ToFloat32(t *testing.T) {
	o := make([]float32, dataLen)
	Float64ToFloat32Slice(dataFloat64, o)
	testResultFloat32(o, t)
}

func TestFloat32ConverterToFloat32(t *testing.T) {
	o := make([]float32, dataLen)
	input := make([]byte, dataLen*Float32.SampleSize())
	for i, s := range dataFloat32 {
		binary.LittleEndian.PutUint32(input[i*Float32.SampleSize():], uint32(math.Float32bits(s)))
	}

	Float32.ToFloat32(input, o)
	testResultFloat32(o, t)
}

func TestFloat32ConverterToFloat32Interleaved(t *testing.T) {
	outputs := make([][]float32, dataLen)
	for i := range outputs {
		outputs[i] = make([]float32, 2)
	}

	input := make([]byte, dataLen*2*Float32.SampleSize())
	for i, s := range dataFloat32 {
		binary.LittleEndian.PutUint32(input[i*Float32.SampleSize():], uint32(math.Float32bits(s)))
	}
	for i, s := range dataFloat32 {
		binary.LittleEndian.PutUint32(input[(dataLen+i)*Float32.SampleSize():], uint32(math.Float32bits(s)))
	}

	Float32.ToFloat32Interleaved(input, outputs)
	testResultFloat32([]float32{outputs[0][0], outputs[1][0], outputs[2][0], outputs[3][0], outputs[4][0]}, t)
	testResultFloat32([]float32{outputs[0][1], outputs[1][1], outputs[2][1], outputs[3][1], outputs[4][1]}, t)
}

func TestFloat32ConverterFromFloat32(t *testing.T) {
	output := make([]byte, dataLen*Float32.SampleSize())
	Float32.FromFloat32(dataFloat32, output)
	testResultFloat32([]float32{
		math.Float32frombits(binary.LittleEndian.Uint32(output[0*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[1*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[2*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[3*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[4*Float32.SampleSize():])),
	}, t)
}

func TestFloat32ConverterFromFloat32Interleaved(t *testing.T) {
	inputs := make([][]float32, dataLen)
	for i := range inputs {
		inputs[i] = []float32{dataFloat32[i], dataFloat32[i]}
	}
	output := make([]byte, dataLen*2*Float32.SampleSize())
	Float32.FromFloat32Interleaved(inputs, output)
	testResultFloat32([]float32{
		math.Float32frombits(binary.LittleEndian.Uint32(output[0*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[1*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[2*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[3*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[4*Float32.SampleSize():])),
	}, t)
	testResultFloat32([]float32{
		math.Float32frombits(binary.LittleEndian.Uint32(output[5*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[6*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[7*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[8*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[9*Float32.SampleSize():])),
	}, t)
}

func TestFloat32ConverterToFloat64(t *testing.T) {
	o := make([]float64, dataLen)
	input := make([]byte, dataLen*Float32.SampleSize())
	for i, s := range dataFloat32 {
		binary.LittleEndian.PutUint32(input[i*Float32.SampleSize():], uint32(math.Float32bits(s)))
	}

	Float32.ToFloat64(input, o)
	testResultFloat64(o, t)
}

func TestFloat32ConverterToFloat64Interleaved(t *testing.T) {
	outputs := make([][]float64, dataLen)
	for i := range outputs {
		outputs[i] = make([]float64, 2)
	}

	input := make([]byte, dataLen*2*Float32.SampleSize())
	for i, s := range dataFloat32 {
		binary.LittleEndian.PutUint32(input[i*Float32.SampleSize():], uint32(math.Float32bits(s)))
	}
	for i, s := range dataFloat32 {
		binary.LittleEndian.PutUint32(input[(dataLen+i)*Float32.SampleSize():], uint32(math.Float32bits(s)))
	}

	Float32.ToFloat64Interleaved(input, outputs)
	testResultFloat64([]float64{outputs[0][0], outputs[1][0], outputs[2][0], outputs[3][0], outputs[4][0]}, t)
	testResultFloat64([]float64{outputs[0][1], outputs[1][1], outputs[2][1], outputs[3][1], outputs[4][1]}, t)
}

func TestFloat32ConverterFromFloat64(t *testing.T) {
	output := make([]byte, dataLen*Float32.SampleSize())
	Float32.FromFloat64(dataFloat64, output)
	testResultFloat32([]float32{
		math.Float32frombits(binary.LittleEndian.Uint32(output[0*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[1*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[2*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[3*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[4*Float32.SampleSize():])),
	}, t)
}

func TestFloat32ConverterFromFloat64Interleaved(t *testing.T) {
	inputs := make([][]float64, dataLen)
	for i := range inputs {
		inputs[i] = []float64{dataFloat64[i], dataFloat64[i]}
	}
	output := make([]byte, dataLen*2*Float32.SampleSize())
	Float32.FromFloat64Interleaved(inputs, output)
	testResultFloat32([]float32{
		math.Float32frombits(binary.LittleEndian.Uint32(output[0*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[1*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[2*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[3*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[4*Float32.SampleSize():])),
	}, t)
	testResultFloat32([]float32{
		math.Float32frombits(binary.LittleEndian.Uint32(output[5*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[6*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[7*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[8*Float32.SampleSize():])),
		math.Float32frombits(binary.LittleEndian.Uint32(output[9*Float32.SampleSize():])),
	}, t)
}

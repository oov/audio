package converter

import (
	"encoding/binary"
	"testing"
)

func TestInt32ToFloat32(t *testing.T) {
	o := make([]float32, dataLen)
	Int32ToFloat32Slice(dataInt32, o)
	testResultFloat32(o, t)
}

func TestFloat32ToInt32(t *testing.T) {
	o := make([]int32, dataLen)
	Float32ToInt32Slice(dataFloat32, o)
	testResultInt32(o, t)
}

func TestInt32ToFloat64(t *testing.T) {
	o := make([]float64, dataLen)
	Int32ToFloat64Slice(dataInt32, o)
	testResultFloat64(o, t)
}

func TestFloat64ToInt32(t *testing.T) {
	o := make([]int32, dataLen)
	Float64ToInt32Slice(dataFloat64, o)
	testResultInt32(o, t)
}

func TestInt32ConverterToFloat32(t *testing.T) {
	o := make([]float32, dataLen)
	input := make([]byte, dataLen*Int32.SampleSize())
	for i, s := range dataInt32 {
		binary.LittleEndian.PutUint32(input[i*Int32.SampleSize():], uint32(s))
	}

	Int32.ToFloat32(input, o)
	testResultFloat32(o, t)
}

func TestInt32ConverterToFloat32Interleaved(t *testing.T) {
	outputs := make([][]float32, dataLen)
	for i := range outputs {
		outputs[i] = make([]float32, 2)
	}

	input := make([]byte, dataLen*2*Int32.SampleSize())
	for i, s := range dataInt32 {
		binary.LittleEndian.PutUint32(input[i*Int32.SampleSize():], uint32(s))
	}
	for i, s := range dataInt32 {
		binary.LittleEndian.PutUint32(input[(dataLen+i)*Int32.SampleSize():], uint32(s))
	}

	Int32.ToFloat32Interleaved(input, outputs)
	testResultFloat32([]float32{outputs[0][0], outputs[1][0], outputs[2][0], outputs[3][0], outputs[4][0]}, t)
	testResultFloat32([]float32{outputs[0][1], outputs[1][1], outputs[2][1], outputs[3][1], outputs[4][1]}, t)
}

func TestInt32ConverterFromFloat32(t *testing.T) {
	output := make([]byte, dataLen*Int32.SampleSize())
	Int32.FromFloat32(dataFloat32, output)
	testResultInt32([]int32{
		int32(binary.LittleEndian.Uint32(output[0*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[1*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[2*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[3*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[4*Int32.SampleSize():])),
	}, t)
}

func TestInt32ConverterFromFloat32Interleaved(t *testing.T) {
	inputs := make([][]float32, dataLen)
	for i := range inputs {
		inputs[i] = []float32{dataFloat32[i], dataFloat32[i]}
	}
	output := make([]byte, dataLen*2*Int32.SampleSize())
	Int32.FromFloat32Interleaved(inputs, output)
	testResultInt32([]int32{
		int32(binary.LittleEndian.Uint32(output[0*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[1*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[2*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[3*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[4*Int32.SampleSize():])),
	}, t)
	testResultInt32([]int32{
		int32(binary.LittleEndian.Uint32(output[5*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[6*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[7*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[8*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[9*Int32.SampleSize():])),
	}, t)
}

func TestInt32ConverterToFloat64(t *testing.T) {
	o := make([]float64, dataLen)
	input := make([]byte, dataLen*Int32.SampleSize())
	for i, s := range dataInt32 {
		binary.LittleEndian.PutUint32(input[i*Int32.SampleSize():], uint32(s))
	}

	Int32.ToFloat64(input, o)
	testResultFloat64(o, t)
}

func TestInt32ConverterToFloat64Interleaved(t *testing.T) {
	outputs := make([][]float64, dataLen)
	for i := range outputs {
		outputs[i] = make([]float64, 2)
	}

	input := make([]byte, dataLen*2*Int32.SampleSize())
	for i, s := range dataInt32 {
		binary.LittleEndian.PutUint32(input[i*Int32.SampleSize():], uint32(s))
	}
	for i, s := range dataInt32 {
		binary.LittleEndian.PutUint32(input[(dataLen+i)*Int32.SampleSize():], uint32(s))
	}

	Int32.ToFloat64Interleaved(input, outputs)
	testResultFloat64([]float64{outputs[0][0], outputs[1][0], outputs[2][0], outputs[3][0], outputs[4][0]}, t)
	testResultFloat64([]float64{outputs[0][1], outputs[1][1], outputs[2][1], outputs[3][1], outputs[4][1]}, t)
}

func TestInt32ConverterFromFloat64(t *testing.T) {
	output := make([]byte, dataLen*Int32.SampleSize())
	Int32.FromFloat64(dataFloat64, output)
	testResultInt32([]int32{
		int32(binary.LittleEndian.Uint32(output[0*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[1*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[2*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[3*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[4*Int32.SampleSize():])),
	}, t)
}

func TestInt32ConverterFromFloat64Interleaved(t *testing.T) {
	inputs := make([][]float64, dataLen)
	for i := range inputs {
		inputs[i] = []float64{dataFloat64[i], dataFloat64[i]}
	}
	output := make([]byte, dataLen*2*Int32.SampleSize())
	Int32.FromFloat64Interleaved(inputs, output)
	testResultInt32([]int32{
		int32(binary.LittleEndian.Uint32(output[0*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[1*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[2*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[3*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[4*Int32.SampleSize():])),
	}, t)
	testResultInt32([]int32{
		int32(binary.LittleEndian.Uint32(output[5*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[6*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[7*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[8*Int32.SampleSize():])),
		int32(binary.LittleEndian.Uint32(output[9*Int32.SampleSize():])),
	}, t)
}

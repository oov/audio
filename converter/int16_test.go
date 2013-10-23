package converter

import (
	"encoding/binary"
	"testing"
)

func TestInt16ToFloat32(t *testing.T) {
	o := make([]float32, dataLen)
	Int16ToFloat32Slice(dataInt16, o)
	testResultFloat32(o, t)
}

func TestFloat32ToInt16(t *testing.T) {
	o := make([]int16, dataLen)
	Float32ToInt16Slice(dataFloat32, o)
	testResultInt16(o, t)
}

func TestInt16ToFloat64(t *testing.T) {
	o := make([]float64, dataLen)
	Int16ToFloat64Slice(dataInt16, o)
	testResultFloat64(o, t)
}

func TestFloat64ToInt16(t *testing.T) {
	o := make([]int16, dataLen)
	Float64ToInt16Slice(dataFloat64, o)
	testResultInt16(o, t)
}

func TestInt16ConverterToFloat32(t *testing.T) {
	o := make([]float32, dataLen)
	input := make([]byte, dataLen*Int16.SampleSize())
	for i, s := range dataInt16 {
		binary.LittleEndian.PutUint16(input[i*Int16.SampleSize():], uint16(s))
	}

	Int16.ToFloat32(input, o)
	testResultFloat32(o, t)
}

func TestInt16ConverterToFloat32Interleaved(t *testing.T) {
	outputs := make([][]float32, dataLen)
	for i := range outputs {
		outputs[i] = make([]float32, 2)
	}

	input := make([]byte, dataLen*2*Int16.SampleSize())
	for i, s := range dataInt16 {
		binary.LittleEndian.PutUint16(input[i*Int16.SampleSize():], uint16(s))
	}
	for i, s := range dataInt16 {
		binary.LittleEndian.PutUint16(input[(dataLen+i)*Int16.SampleSize():], uint16(s))
	}

	Int16.ToFloat32Interleaved(input, outputs)
	testResultFloat32([]float32{outputs[0][0], outputs[1][0], outputs[2][0], outputs[3][0], outputs[4][0]}, t)
	testResultFloat32([]float32{outputs[0][1], outputs[1][1], outputs[2][1], outputs[3][1], outputs[4][1]}, t)
}

func TestInt16ConverterFromFloat32(t *testing.T) {
	output := make([]byte, dataLen*Int16.SampleSize())
	Int16.FromFloat32(dataFloat32, output)
	testResultInt16([]int16{
		int16(binary.LittleEndian.Uint16(output[0*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[1*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[2*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[3*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[4*Int16.SampleSize():])),
	}, t)
}

func TestInt16ConverterFromFloat32Interleaved(t *testing.T) {
	inputs := make([][]float32, dataLen)
	for i := range inputs {
		inputs[i] = []float32{dataFloat32[i], dataFloat32[i]}
	}
	output := make([]byte, dataLen*2*Int16.SampleSize())
	Int16.FromFloat32Interleaved(inputs, output)
	testResultInt16([]int16{
		int16(binary.LittleEndian.Uint16(output[0*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[1*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[2*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[3*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[4*Int16.SampleSize():])),
	}, t)
	testResultInt16([]int16{
		int16(binary.LittleEndian.Uint16(output[5*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[6*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[7*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[8*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[9*Int16.SampleSize():])),
	}, t)
}

func TestInt16ConverterToFloat64(t *testing.T) {
	o := make([]float64, dataLen)
	input := make([]byte, dataLen*Int16.SampleSize())
	for i, s := range dataInt16 {
		binary.LittleEndian.PutUint16(input[i*Int16.SampleSize():], uint16(s))
	}

	Int16.ToFloat64(input, o)
	testResultFloat64(o, t)
}

func TestInt16ConverterToFloat64Interleaved(t *testing.T) {
	outputs := make([][]float64, dataLen)
	for i := range outputs {
		outputs[i] = make([]float64, 2)
	}

	input := make([]byte, dataLen*2*Int16.SampleSize())
	for i, s := range dataInt16 {
		binary.LittleEndian.PutUint16(input[i*Int16.SampleSize():], uint16(s))
	}
	for i, s := range dataInt16 {
		binary.LittleEndian.PutUint16(input[(dataLen+i)*Int16.SampleSize():], uint16(s))
	}

	Int16.ToFloat64Interleaved(input, outputs)
	testResultFloat64([]float64{outputs[0][0], outputs[1][0], outputs[2][0], outputs[3][0], outputs[4][0]}, t)
	testResultFloat64([]float64{outputs[0][1], outputs[1][1], outputs[2][1], outputs[3][1], outputs[4][1]}, t)
}

func TestInt16ConverterFromFloat64(t *testing.T) {
	output := make([]byte, dataLen*Int16.SampleSize())
	Int16.FromFloat64(dataFloat64, output)
	testResultInt16([]int16{
		int16(binary.LittleEndian.Uint16(output[0*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[1*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[2*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[3*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[4*Int16.SampleSize():])),
	}, t)
}

func TestInt16ConverterFromFloat64Interleaved(t *testing.T) {
	inputs := make([][]float64, dataLen)
	for i := range inputs {
		inputs[i] = []float64{dataFloat64[i], dataFloat64[i]}
	}
	output := make([]byte, dataLen*2*Int16.SampleSize())
	Int16.FromFloat64Interleaved(inputs, output)
	testResultInt16([]int16{
		int16(binary.LittleEndian.Uint16(output[0*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[1*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[2*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[3*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[4*Int16.SampleSize():])),
	}, t)
	testResultInt16([]int16{
		int16(binary.LittleEndian.Uint16(output[5*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[6*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[7*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[8*Int16.SampleSize():])),
		int16(binary.LittleEndian.Uint16(output[9*Int16.SampleSize():])),
	}, t)
}

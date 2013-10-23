package converter

import (
	"encoding/binary"
	"testing"
)

func Int24Write(s int32, out []byte) {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(s))
	copy(out, buf[1:])
}

func Int24Read(s []byte) int32 {
	buf := make([]byte, 4)
	copy(buf[1:], s)
	return int32(binary.LittleEndian.Uint32(buf))
}

func TestInt24ConverterToFloat32(t *testing.T) {
	o := make([]float32, dataLen)
	input := make([]byte, dataLen*Int24.SampleSize())
	for i, s := range dataInt24 {
		Int24Write(s, input[i*Int24.SampleSize():])
	}

	Int24.ToFloat32(input, o)
	testResultFloat32(o, t)
}

func TestInt24ConverterToFloat32Interleaved(t *testing.T) {

	outputs := make([][]float32, dataLen)
	for i := range outputs {
		outputs[i] = make([]float32, 2)
	}

	input := make([]byte, dataLen*2*Int24.SampleSize())
	for i, s := range dataInt24 {
		Int24Write(s, input[i*Int24.SampleSize():])
	}
	for i, s := range dataInt24 {
		Int24Write(s, input[(dataLen+i)*Int24.SampleSize():])
	}

	Int24.ToFloat32Interleaved(input, outputs)
	testResultFloat32([]float32{outputs[0][0], outputs[1][0], outputs[2][0], outputs[3][0], outputs[4][0]}, t)
	testResultFloat32([]float32{outputs[0][1], outputs[1][1], outputs[2][1], outputs[3][1], outputs[4][1]}, t)
}

func TestInt24ConverterFromFloat32(t *testing.T) {
	output := make([]byte, dataLen*Int24.SampleSize())
	Int24.FromFloat32(dataFloat32, output)
	testResultInt24([]int32{
		Int24Read(output[0*Int24.SampleSize():]),
		Int24Read(output[1*Int24.SampleSize():]),
		Int24Read(output[2*Int24.SampleSize():]),
		Int24Read(output[3*Int24.SampleSize():]),
		Int24Read(output[4*Int24.SampleSize():]),
	}, t)
}

func TestInt24ConverterFromFloat32Interleaved(t *testing.T) {
	inputs := make([][]float32, dataLen)
	for i := range inputs {
		inputs[i] = []float32{dataFloat32[i], dataFloat32[i]}
	}
	output := make([]byte, dataLen*2*Int24.SampleSize())
	Int24.FromFloat32Interleaved(inputs, output)
	testResultInt24([]int32{
		Int24Read(output[0*Int24.SampleSize():]),
		Int24Read(output[1*Int24.SampleSize():]),
		Int24Read(output[2*Int24.SampleSize():]),
		Int24Read(output[3*Int24.SampleSize():]),
		Int24Read(output[4*Int24.SampleSize():]),
	}, t)
	testResultInt24([]int32{
		Int24Read(output[5*Int24.SampleSize():]),
		Int24Read(output[6*Int24.SampleSize():]),
		Int24Read(output[7*Int24.SampleSize():]),
		Int24Read(output[8*Int24.SampleSize():]),
		Int24Read(output[9*Int24.SampleSize():]),
	}, t)
}

func TestInt24ConverterToFloat64(t *testing.T) {
	o := make([]float64, dataLen)
	input := make([]byte, dataLen*Int24.SampleSize())
	for i, s := range dataInt24 {
		Int24Write(s, input[i*Int24.SampleSize():])
	}

	Int24.ToFloat64(input, o)
	testResultFloat64(o, t)
}

func TestInt24ConverterToFloat64Interleaved(t *testing.T) {
	outputs := make([][]float64, dataLen)
	for i := range outputs {
		outputs[i] = make([]float64, 2)
	}

	input := make([]byte, dataLen*2*Int24.SampleSize())
	for i, s := range dataInt24 {
		Int24Write(s, input[i*Int24.SampleSize():])
	}
	for i, s := range dataInt24 {
		Int24Write(s, input[(dataLen+i)*Int24.SampleSize():])
	}

	Int24.ToFloat64Interleaved(input, outputs)
	testResultFloat64([]float64{outputs[0][0], outputs[1][0], outputs[2][0], outputs[3][0], outputs[4][0]}, t)
	testResultFloat64([]float64{outputs[0][1], outputs[1][1], outputs[2][1], outputs[3][1], outputs[4][1]}, t)
}

func TestInt24ConverterFromFloat64(t *testing.T) {
	output := make([]byte, dataLen*Int24.SampleSize())
	Int24.FromFloat64(dataFloat64, output)
	testResultInt24([]int32{
		Int24Read(output[0*Int24.SampleSize():]),
		Int24Read(output[1*Int24.SampleSize():]),
		Int24Read(output[2*Int24.SampleSize():]),
		Int24Read(output[3*Int24.SampleSize():]),
		Int24Read(output[4*Int24.SampleSize():]),
	}, t)
}

func TestInt24ConverterFromFloat64Interleaved(t *testing.T) {
	inputs := make([][]float64, dataLen)
	for i := range inputs {
		inputs[i] = []float64{dataFloat64[i], dataFloat64[i]}
	}
	output := make([]byte, dataLen*2*Int24.SampleSize())
	Int24.FromFloat64Interleaved(inputs, output)
	testResultInt24([]int32{
		Int24Read(output[0*Int24.SampleSize():]),
		Int24Read(output[1*Int24.SampleSize():]),
		Int24Read(output[2*Int24.SampleSize():]),
		Int24Read(output[3*Int24.SampleSize():]),
		Int24Read(output[4*Int24.SampleSize():]),
	}, t)
	testResultInt24([]int32{
		Int24Read(output[5*Int24.SampleSize():]),
		Int24Read(output[6*Int24.SampleSize():]),
		Int24Read(output[7*Int24.SampleSize():]),
		Int24Read(output[8*Int24.SampleSize():]),
		Int24Read(output[9*Int24.SampleSize():]),
	}, t)
}

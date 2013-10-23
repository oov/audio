package converter

import (
	"testing"
)

func TestUint8ToFloat32(t *testing.T) {
	o := make([]float32, dataLen)
	Uint8ToFloat32Slice(dataUint8, o)
	testResultFloat32(o, t)
}

func TestFloat32ToUint8(t *testing.T) {
	o := make([]uint8, dataLen)
	Float32ToUint8Slice(dataFloat32, o)
	testResultUint8(o, t)
}

func TestUint8ToFloat64(t *testing.T) {
	o := make([]float64, dataLen)
	Uint8ToFloat64Slice(dataUint8, o)
	testResultFloat64(o, t)
}

func TestFloat64ToUint8(t *testing.T) {
	o := make([]uint8, dataLen)
	Float64ToUint8Slice(dataFloat64, o)
	testResultUint8(o, t)
}

func TestUint8ConverterToFloat32(t *testing.T) {
	o := make([]float32, dataLen)
	input := make([]byte, dataLen*Uint8.SampleSize())
	for i, s := range dataUint8 {
		input[i*Uint8.SampleSize()] = s
	}

	Uint8.ToFloat32(input, o)
	testResultFloat32(o, t)
}

func TestUint8ConverterToFloat32Interleaved(t *testing.T) {
	outputs := make([][]float32, dataLen)
	for i := range outputs {
		outputs[i] = make([]float32, 2)
	}

	input := make([]byte, dataLen*2*Uint8.SampleSize())
	for i, s := range dataUint8 {
		input[i*Uint8.SampleSize()] = s
	}
	for i, s := range dataUint8 {
		input[(dataLen+i)*Uint8.SampleSize()] = s
	}

	Uint8.ToFloat32Interleaved(input, outputs)
	testResultFloat32([]float32{outputs[0][0], outputs[1][0], outputs[2][0], outputs[3][0], outputs[4][0]}, t)
	testResultFloat32([]float32{outputs[0][1], outputs[1][1], outputs[2][1], outputs[3][1], outputs[4][1]}, t)
}

func TestUint8ConverterFromFloat32(t *testing.T) {
	output := make([]byte, dataLen*Uint8.SampleSize())
	Uint8.FromFloat32(dataFloat32, output)
	testResultUint8(output[:dataLen], t)
}

func TestUint8ConverterFromFloat32Interleaved(t *testing.T) {
	inputs := make([][]float32, dataLen)
	for i := range inputs {
		inputs[i] = []float32{dataFloat32[i], dataFloat32[i]}
	}
	output := make([]byte, dataLen*2*Uint8.SampleSize())
	Uint8.FromFloat32Interleaved(inputs, output)
	testResultUint8(output[:dataLen], t)
	testResultUint8(output[dataLen:], t)
}

func TestUint8ConverterToFloat64(t *testing.T) {
	o := make([]float64, dataLen)
	input := make([]byte, dataLen*Uint8.SampleSize())
	for i, s := range dataUint8 {
		input[i*Uint8.SampleSize()] = s
	}

	Uint8.ToFloat64(input, o)
	testResultFloat64(o, t)
}

func TestUint8ConverterToFloat64Interleaved(t *testing.T) {
	outputs := make([][]float64, dataLen)
	for i := range outputs {
		outputs[i] = make([]float64, 2)
	}

	input := make([]byte, dataLen*2*Uint8.SampleSize())
	for i, s := range dataUint8 {
		input[i*Uint8.SampleSize()] = s
	}
	for i, s := range dataUint8 {
		input[(dataLen+i)*Uint8.SampleSize()] = s
	}

	Uint8.ToFloat64Interleaved(input, outputs)
	testResultFloat64([]float64{outputs[0][0], outputs[1][0], outputs[2][0], outputs[3][0], outputs[4][0]}, t)
	testResultFloat64([]float64{outputs[0][1], outputs[1][1], outputs[2][1], outputs[3][1], outputs[4][1]}, t)
}

func TestUint8ConverterFromFloat64(t *testing.T) {
	output := make([]byte, dataLen*Uint8.SampleSize())
	Uint8.FromFloat64(dataFloat64, output)
	testResultUint8(output[:dataLen], t)
}

func TestUint8ConverterFromFloat64Interleaved(t *testing.T) {
	inputs := make([][]float64, dataLen)
	for i := range inputs {
		inputs[i] = []float64{dataFloat64[i], dataFloat64[i]}
	}
	output := make([]byte, dataLen*2*Uint8.SampleSize())
	Uint8.FromFloat64Interleaved(inputs, output)
	testResultUint8(output[:dataLen], t)
	testResultUint8(output[dataLen:], t)
}

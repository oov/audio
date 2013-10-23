package converter

var (
	Uint8 Uint8Converter
)

type Uint8Converter uint8

func (c Uint8Converter) SampleSize() int {
	return 1
}

func (c Uint8Converter) ToFloat32(input []byte, output []float32) {
	for i, s := range input {
		output[i] = Uint8ToFloat32(s)
	}
}

func (c Uint8Converter) ToFloat32Interleaved(input []byte, outputs [][]float32) {
	chs := len(outputs)
	for ch, output := range outputs {
		for i, o, ln := ch, 0, len(input); i < ln; i += chs {
			output[o] = Uint8ToFloat32(input[i])
			o++
		}
	}
}

func (c Uint8Converter) FromFloat32(input []float32, output []byte) {
	for i, s := range input {
		output[i] = Float32ToUint8(s)
	}
}

func (c Uint8Converter) FromFloat32Interleaved(inputs [][]float32, output []byte) {
	for i, o := 0, 0; o < len(output); i++ {
		for _, input := range inputs {
			output[o] = Float32ToUint8(input[i])
			o++
		}
	}
}

func (c Uint8Converter) ToFloat64(input []byte, output []float64) {
	for i, s := range input {
		output[i] = Uint8ToFloat64(s)
	}
}

func (c Uint8Converter) ToFloat64Interleaved(input []byte, outputs [][]float64) {
	chs := len(outputs)
	for ch, output := range outputs {
		for i, o, ln := ch, 0, len(input); i < ln; i += chs {
			output[o] = Uint8ToFloat64(input[i])
			o++
		}
	}
}

func (c Uint8Converter) FromFloat64(input []float64, output []byte) {
	for i, s := range input {
		output[i] = Float64ToUint8(s)
	}
}

func (c Uint8Converter) FromFloat64Interleaved(inputs [][]float64, output []byte) {
	for i, o := 0, 0; o < len(output); i++ {
		for _, input := range inputs {
			output[o] = Float64ToUint8(input[i])
			o++
		}
	}
}

const (
	uint8Half    = 127
	uint8Shifter = 128
	uint8Divider = 1.0 / 128.0
)

func Uint8ToFloat32(s uint8) float32 {
	return (float32(s) - uint8Shifter) * uint8Divider
}

func Uint8ToFloat64(s uint8) float64 {
	return (float64(s) - uint8Shifter) * uint8Divider
}

func Float32ToUint8(s float32) uint8 {
	return uint8(s*uint8Half + uint8Shifter)
}

func Float64ToUint8(s float64) uint8 {
	return uint8(s*uint8Half + uint8Shifter)
}

func Uint8ToFloat32Slice(input []uint8, output []float32) {
	for i, s := range input {
		output[i] = Uint8ToFloat32(s)
	}
}

func Uint8ToFloat64Slice(input []uint8, output []float64) {
	for i, s := range input {
		output[i] = Uint8ToFloat64(s)
	}
}

func Float32ToUint8Slice(input []float32, output []uint8) {
	for i, s := range input {
		output[i] = Float32ToUint8(s)
	}
}

func Float64ToUint8Slice(input []float64, output []uint8) {
	for i, s := range input {
		output[i] = Float64ToUint8(s)
	}
}

func ByteToUint8(a byte) uint8 {
	return a
}

func Uint8ToByte(s uint8) byte {
	return s
}

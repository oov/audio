package converter

var (
	Int24 Int24Converter
)

type Int24Converter int32

func (c Int24Converter) SampleSize() int {
	return 3
}

func (c Int24Converter) ToFloat32(input []byte, output []float32) {
	for i, o, ln := 0, 0, len(input); i < ln; i += c.SampleSize() {
		output[o] = Int24ToFloat32(ByteToInt24(input[i], input[i+1], input[i+2]))
		o++
	}
}

func (c Int24Converter) ToFloat32Interleaved(input []byte, outputs [][]float32) {
	chs := len(outputs)
	for ch, output := range outputs {
		for i, o, ln := ch*c.SampleSize(), 0, len(input); i < ln; i += chs * c.SampleSize() {
			output[o] = Int24ToFloat32(ByteToInt24(input[i], input[i+1], input[i+2]))
			o++
		}
	}
}

func (c Int24Converter) FromFloat32(input []float32, output []byte) {
	o := 0
	for _, s := range input {
		output[o], output[o+1], output[o+2] = Int24ToByte(Float32ToInt24(s))
		o += c.SampleSize()
	}
}

func (c Int24Converter) FromFloat32Interleaved(inputs [][]float32, output []byte) {
	for i, o := 0, 0; o < len(output); i++ {
		for _, input := range inputs {
			output[o], output[o+1], output[o+2] = Int24ToByte(Float32ToInt24(input[i]))
			o += c.SampleSize()
		}
	}
}

func (c Int24Converter) ToFloat64(input []byte, output []float64) {
	for i, o, ln := 0, 0, len(input); i < ln; i += c.SampleSize() {
		output[o] = Int24ToFloat64(ByteToInt24(input[i], input[i+1], input[i+2]))
		o++
	}
}

func (c Int24Converter) ToFloat64Interleaved(input []byte, outputs [][]float64) {
	chs := len(outputs)
	for ch, output := range outputs {
		for i, o, ln := ch*c.SampleSize(), 0, len(input); i < ln; i += chs * c.SampleSize() {
			output[o] = Int24ToFloat64(ByteToInt24(input[i], input[i+1], input[i+2]))
			o++
		}
	}
}

func (c Int24Converter) FromFloat64(input []float64, output []byte) {
	o := 0
	for _, s := range input {
		output[o], output[o+1], output[o+2] = Int24ToByte(Float64ToInt24(s))
		o += c.SampleSize()
	}
}

func (c Int24Converter) FromFloat64Interleaved(inputs [][]float64, output []byte) {
	for i, o := 0, 0; o < len(output); i++ {
		for _, input := range inputs {
			output[o], output[o+1], output[o+2] = Int24ToByte(Float64ToInt24(input[i]))
			o += c.SampleSize()
		}
	}
}

const (
	int24Divider = 1.0 / (2147483648.0)
	int24Max     = 2147483648.0 - 256.0
)

func Int24ToFloat32(s int32) float32 {
	return float32(s) * int24Divider
}

func Int24ToFloat64(s int32) float64 {
	return float64(s) * int24Divider
}

func Float32ToInt24(s float32) int32 {
	return int32(s * int24Max)
}

func Float64ToInt24(s float64) int32 {
	return int32(s * int24Max)
}

/*
func Int24ToFloat32Slice(input []int32, output []float32) {
	for i, s := range input {
		output[i] = Int24ToFloat32(s)
	}
}

func Int24ToFloat64Slice(input []int32, output []float64) {
	for i, s := range input {
		output[i] = Int24ToFloat64(s)
	}
}

func Float32ToInt24Slice(input []float32, output []int32) {
	for i, s := range input {
		output[i] = Float32ToInt24(s)
	}
}

func Float64ToInt24Slice(input []float64, output []int32) {
	for i, s := range input {
		output[i] = Float64ToInt24(s)
	}
}
*/

func ByteToInt24(a, b, c byte) int32 {
	return (int32(a) << 8) | (int32(b) << 16) | (int32(c) << 24)
}

func Int24ToByte(s int32) (byte, byte, byte) {
	return byte(s >> 8), byte(s >> 16), byte(s >> 24)
}

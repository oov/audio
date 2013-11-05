package converter

var (
	Int32 Int32Converter
)

type Int32Converter int32

func (c Int32Converter) SampleSize() int {
	return 4
}

func (c Int32Converter) ToFloat32(input []byte, output []float32) {
	for i, o, ln := 0, 0, len(input); i < ln; i += c.SampleSize() {
		output[o] = Int32ToFloat32(ByteToInt32(input[i], input[i+1], input[i+2], input[i+3]))
		o++
	}
}

func (c Int32Converter) ToFloat32Interleaved(input []byte, outputs [][]float32) {
	chs := len(outputs)
	for ch, output := range outputs {
		for i, o, ln := ch*c.SampleSize(), 0, len(input); i < ln; i += chs * c.SampleSize() {
			output[o] = Int32ToFloat32(ByteToInt32(input[i], input[i+1], input[i+2], input[i+3]))
			o++
		}
	}
}

func (c Int32Converter) FromFloat32(input []float32, output []byte) {
	o := 0
	for _, s := range input {
		output[o], output[o+1], output[o+2], output[o+3] = Int32ToByte(Float32ToInt32(s))
		o += c.SampleSize()
	}
}

func (c Int32Converter) FromFloat32Interleaved(inputs [][]float32, output []byte) {
	for i, o := 0, 0; o < len(output); i++ {
		for _, input := range inputs {
			output[o], output[o+1], output[o+2], output[o+3] = Int32ToByte(Float32ToInt32(input[i]))
			o += c.SampleSize()
		}
	}
}

func (c Int32Converter) ToFloat64(input []byte, output []float64) {
	for i, o, ln := 0, 0, len(input); i < ln; i += c.SampleSize() {
		output[o] = Int32ToFloat64(ByteToInt32(input[i], input[i+1], input[i+2], input[i+3]))
		o++
	}
}

func (c Int32Converter) ToFloat64Interleaved(input []byte, outputs [][]float64) {
	chs := len(outputs)
	for ch, output := range outputs {
		for i, o, ln := ch*c.SampleSize(), 0, len(input); i < ln; i += chs * c.SampleSize() {
			output[o] = Int32ToFloat64(ByteToInt32(input[i], input[i+1], input[i+2], input[i+3]))
			o++
		}
	}
}

func (c Int32Converter) FromFloat64(input []float64, output []byte) {
	o := 0
	for _, s := range input {
		output[o], output[o+1], output[o+2], output[o+3] = Int32ToByte(Float64ToInt32(s))
		o += c.SampleSize()
	}
}

func (c Int32Converter) FromFloat64Interleaved(inputs [][]float64, output []byte) {
	for i, o := 0, 0; o < len(output); i++ {
		for _, input := range inputs {
			output[o], output[o+1], output[o+2], output[o+3] = Int32ToByte(Float64ToInt32(input[i]))
			o += c.SampleSize()
		}
	}
}

const (
	int32Max     = 2147483647
	int32Divider = 1.0 / 2147483648.0
)

func Int32ToFloat32(s int32) float32 {
	return float32(s) * int32Divider
}

func Int32ToFloat64(s int32) float64 {
	return float64(s) * int32Divider
}

func Float32ToInt32(s float32) int32 {
	return int32(float64(s) * int32Max)
}

func Float64ToInt32(s float64) int32 {
	return int32(s * int32Max)
}

func Int32ToFloat32Slice(input []int32, output []float32) {
	for i, s := range input {
		output[i] = Int32ToFloat32(s)
	}
}

func Int32ToFloat64Slice(input []int32, output []float64) {
	for i, s := range input {
		output[i] = Int32ToFloat64(s)
	}
}

func Float32ToInt32Slice(input []float32, output []int32) {
	for i, s := range input {
		output[i] = Float32ToInt32(s)
	}
}

func Float64ToInt32Slice(input []float64, output []int32) {
	for i, s := range input {
		output[i] = Float64ToInt32(s)
	}
}

func ByteToInt32(a, b, c, d byte) int32 {
	return int32(a) | (int32(b) << 8) | (int32(c) << 16) | (int32(d) << 24)
}

func Int32ToByte(s int32) (byte, byte, byte, byte) {
	return byte(s), byte(s >> 8), byte(s >> 16), byte(s >> 24)
}

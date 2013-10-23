package converter

var (
	Int16 Int16Converter
)

type Int16Converter int16

func (c Int16Converter) SampleSize() int {
	return 2
}

func (c Int16Converter) ToFloat32(input []byte, output []float32) {
	for i, o, ln := 0, 0, len(input); i < ln; i += c.SampleSize() {
		output[o] = Int16ToFloat32(ByteToInt16(input[i], input[i+1]))
		o++
	}
}

func (c Int16Converter) ToFloat32Interleaved(input []byte, outputs [][]float32) {
	chs := len(outputs)
	for ch, output := range outputs {
		for i, o, ln := ch*c.SampleSize(), 0, len(input); i < ln; i += chs * c.SampleSize() {
			output[o] = Int16ToFloat32(ByteToInt16(input[i], input[i+1]))
			o++
		}
	}
}

func (c Int16Converter) FromFloat32(input []float32, output []byte) {
	o := 0
	for _, s := range input {
		output[o], output[o+1] = Int16ToByte(Float32ToInt16(s))
		o += c.SampleSize()
	}
}

func (c Int16Converter) FromFloat32Interleaved(inputs [][]float32, output []byte) {
	for i, o := 0, 0; o < len(output); i++ {
		for _, input := range inputs {
			output[o], output[o+1] = Int16ToByte(Float32ToInt16(input[i]))
			o += c.SampleSize()
		}
	}
}

func (c Int16Converter) ToFloat64(input []byte, output []float64) {
	for i, o, ln := 0, 0, len(input); i < ln; i += c.SampleSize() {
		output[o] = Int16ToFloat64(ByteToInt16(input[i], input[i+1]))
		o++
	}
}

func (c Int16Converter) ToFloat64Interleaved(input []byte, outputs [][]float64) {
	chs := len(outputs)
	for ch, output := range outputs {
		for i, o, ln := ch*c.SampleSize(), 0, len(input); i < ln; i += chs * c.SampleSize() {
			output[o] = Int16ToFloat64(ByteToInt16(input[i], input[i+1]))
			o++
		}
	}
}

func (c Int16Converter) FromFloat64(input []float64, output []byte) {
	o := 0
	for _, s := range input {
		output[o], output[o+1] = Int16ToByte(Float64ToInt16(s))
		o += c.SampleSize()
	}
}

func (c Int16Converter) FromFloat64Interleaved(inputs [][]float64, output []byte) {
	for i, o := 0, 0; o < len(output); i++ {
		for _, input := range inputs {
			output[o], output[o+1] = Int16ToByte(Float64ToInt16(input[i]))
			o += c.SampleSize()
		}
	}
}

const (
	int16Max     = 32767.0
	int16Divider = 1.0 / 32768.0
)

func Int16ToFloat32(s int16) float32 {
	return float32(s) * int16Divider
}

func Int16ToFloat64(s int16) float64 {
	return float64(s) * int16Divider
}

func Float32ToInt16(s float32) int16 {
	return int16(s * int16Max)
}

func Float64ToInt16(s float64) int16 {
	return int16(s * int16Max)
}

func Int16ToFloat32Slice(input []int16, output []float32) {
	for i, s := range input {
		output[i] = Int16ToFloat32(s)
	}
}

func Int16ToFloat64Slice(input []int16, output []float64) {
	for i, s := range input {
		output[i] = Int16ToFloat64(s)
	}
}

func Float32ToInt16Slice(input []float32, output []int16) {
	for i, s := range input {
		output[i] = Float32ToInt16(s)
	}
}

func Float64ToInt16Slice(input []float64, output []int16) {
	for i, s := range input {
		output[i] = Float64ToInt16(s)
	}
}

func ByteToInt16(a, b byte) int16 {
	return int16(a) | (int16(b) << 8)
}

func Int16ToByte(s int16) (byte, byte) {
	return byte(s), byte(s >> 8)
}

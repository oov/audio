// Package converter implements audio sample format converter.
package converter

type FormatConverter interface {
	SampleSize() int
	ToFloat32(input []byte, output []float32)
	ToFloat64(input []byte, output []float64)
	FromFloat32(input []float32, output []byte)
	FromFloat64(input []float64, output []byte)
}

type InterleavedFormatConverter interface {
	SampleSize() int
	ToFloat32Interleaved(input []byte, outputs [][]float32)
	ToFloat64Interleaved(input []byte, outputs [][]float64)
	FromFloat32Interleaved(inputs [][]float32, output []byte)
	FromFloat64Interleaved(inputs [][]float64, output []byte)
}

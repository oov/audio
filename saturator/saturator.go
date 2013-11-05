// Package saturator implements saturator.
package saturator

func Saturate32(s float32) float32 {
	switch {
	case s > 1:
		return 1
	case s < -1:
		return -1
	}
	return s
}

func Saturate64(s float64) float64 {
	switch {
	case s > 1:
		return 1
	case s < -1:
		return -1
	}
	return s
}

func Saturate32Slice(input, output []float32) {
	for i, s := range input {
		output[i] = Saturate32(s)
	}
}

func Saturate64Slice(input, output []float64) {
	for i, s := range input {
		output[i] = Saturate64(s)
	}
}

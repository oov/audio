package converter

import (
	"math"
	"testing"
)

var (
	dataUint8   = []uint8{0, 64, 128, 192, 255}
	dataInt16   = []int16{-32768, -16384, 0, 16383, 32767}
	dataInt24   = []int32{-2147483648, -1073741824, 0, 1073741823, 2147483647}
	dataInt32   = []int32{-2147483648, -1073741824, 0, 1073741823, 2147483647}
	dataFloat32 = []float32{-1, -0.5, 0, 0.5, 1}
	dataFloat64 = []float64{-1, -0.5, 0, 0.5, 1}
	dataLen     = len(dataFloat32)
)

func testResultUint8(a []uint8, t *testing.T) {
	t.Log("min:", a[0], "low:", a[1], "zero:", a[2], "high:", a[3], "max:", a[4])
	for i, s := range dataUint8 {
		if math.Abs(float64(s)-float64(a[i])) > 1 {
			t.Fail()
		}
	}
}

func testResultInt16(a []int16, t *testing.T) {
	t.Log("min:", a[0], "low:", a[1], "zero:", a[2], "high:", a[3], "max:", a[4])
	for i, s := range dataInt16 {
		if math.Abs(float64(s)-float64(a[i])) > 1 {
			t.Fail()
		}
	}
}

func testResultInt24(a []int32, t *testing.T) {
	t.Log("min:", a[0], "low:", a[1], "zero:", a[2], "high:", a[3], "max:", a[4])
	for i, s := range dataInt24 {
		if math.Abs(float64(s)-float64(a[i])) > 0x100 {
			t.Fail()
		}
	}
}

func testResultInt32(a []int32, t *testing.T) {
	t.Log("min:", a[0], "low:", a[1], "zero:", a[2], "high:", a[3], "max:", a[4])
	for i, s := range dataInt32 {
		if math.Abs(float64(s)-float64(a[i])) > 1 {
			t.Fail()
		}
	}
}

func testResultFloat32(a []float32, t *testing.T) {
	t.Log("min:", a[0], "low:", a[1], "zero:", a[2], "high:", a[3], "max:", a[4])
	for i, s := range dataFloat32 {
		if math.Abs(float64(s)-float64(a[i])) > 0.01 {
			t.Fail()
		}
	}
}

func testResultFloat64(a []float64, t *testing.T) {
	t.Log("min:", a[0], "low:", a[1], "zero:", a[2], "high:", a[3], "max:", a[4])
	for i, s := range dataFloat64 {
		if math.Abs(s-a[i]) > 0.01 {
			t.Fail()
		}
	}
}

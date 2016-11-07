package castutil

import (
	"testing"

	"github.com/frozzare/go-assert"
)

func TestToBool(t *testing.T) {
	test := [][]interface{}{
		{'f', true},
		{true, true},
		{false, false},
		{int(0), false},
		{int8(1), true},
		{int16(16), true},
		{int32(32), true},
		{int64(64), true},
		{uint(0), false},
		{uint8(1), true},
		{uint16(16), true},
		{uint32(32), true},
		{uint64(64), true},
		{float32(0.1), true},
		{float64(1.1), true},
		{float64(1.348959), true},
		{"true", true},
		{"false", false},
		{"1", true},
		{"0", false},
		// errors
		{"x", false, true},
		{"1.5", false, true},
	}

	for _, item := range test {
		actual, err := ToBool(item[0])

		if len(item) > 2 {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}

		assert.Equal(t, item[1].(bool), actual)
	}
}

func TestToFloat32(t *testing.T) {
	test := [][]interface{}{
		{'f', 102.0},
		{true, 1.0},
		{false, 0.0},
		{int(0), 0.0},
		{int8(1), 1.0},
		{int16(16), 16.0},
		{int32(32), 32.0},
		{int64(64), 64.0},
		{uint(0), 0.0},
		{uint8(1), 1.0},
		{uint16(16), 16.0},
		{uint32(32), 32.0},
		{uint64(64), 64.0},
		{float32(0.1), 0.10000000149011612},
		{float64(1.1), 1.1},
		{float64(1.348959), 1.348959},
		{"1.2", 1.2},
		// errors
		{"x", 0.0, true},
	}

	for _, item := range test {
		actual, err := ToFloat32(item[0])

		if len(item) > 2 {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}

		assert.Equal(t, float32(item[1].(float64)), actual)
	}
}

func TestToFloat64(t *testing.T) {
	test := [][]interface{}{
		{'f', 102.0},
		{true, 1.0},
		{false, 0.0},
		{int(0), 0.0},
		{int8(1), 1.0},
		{int16(16), 16.0},
		{int32(32), 32.0},
		{int64(64), 64.0},
		{uint(0), 0.0},
		{uint8(1), 1.0},
		{uint16(16), 16.0},
		{uint32(32), 32.0},
		{uint64(64), 64.0},
		{float32(0.1), 0.10000000149011612},
		{float64(1.1), 1.1},
		{float64(1.348959), 1.348959},
		{"1.2", 1.2},
		// errors
		{"x", 0.0, true},
	}

	for _, item := range test {
		actual, err := ToFloat64(item[0])

		if len(item) > 2 {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}

		assert.Equal(t, item[1].(float64), actual)
	}
}

func TestToInt(t *testing.T) {
	test := [][]interface{}{
		{'f', 102},
		{true, 1},
		{false, 0},
		{int(0), 0},
		{int8(1), 1},
		{int16(16), 16},
		{int32(32), 32},
		{int64(64), 64},
		{uint(0), 0},
		{uint8(1), 1},
		{uint16(16), 16},
		{uint32(32), 32},
		{uint64(64), 64},
		{float32(0.1), 0},
		{float64(1.1), 1},
		{float64(1.348959), 1},
		{"1", 1},
		{"1.2", 1},
		// errors
		{"x", 0, true},
	}

	for _, item := range test {
		actual, err := ToInt(item[0])

		if len(item) > 2 {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}

		assert.Equal(t, item[1].(int), actual)
	}
}

func TestToString(t *testing.T) {
	test := [][]interface{}{
		{[]byte("f"), "f"},
		{true, "true"},
		{false, "false"},
		{int(0), "0"},
		{int8(1), "1"},
		{int16(16), "16"},
		{int32(32), "32"},
		{int64(64), "64"},
		{uint(0), "0"},
		{uint8(1), "1"},
		{uint16(16), "16"},
		{uint32(32), "32"},
		{uint64(64), "64"},
		{float32(0.1), "0.100000001490"},
		{float64(1.1), "1.100000000000"},
		{float64(1.348959), "1.348959000000"},
		{"hello", "hello"},
		{[]int{1, 2, 3}, "[1 2 3]"},
		{nil, "<nil>"},
	}

	for _, item := range test {
		actual, err := ToString(item[0])

		assert.Nil(t, err)
		assert.Equal(t, item[1].(string), actual)
	}
}

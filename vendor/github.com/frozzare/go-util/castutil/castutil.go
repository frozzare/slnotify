package castutil

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// getArg returns a arg value or the default value.
func getArg(def interface{}, i int, args ...interface{}) interface{} {
	if len(args) >= i+1 && args[i] != nil {
		return args[i]
	}

	return def
}

// getArgInt returns a arg value as int.
func getArgInt(def int, i int, args ...interface{}) int {
	return getArg(def, i, args...).(int)
}

// ToBool will convert argument to bool or return a error.
func ToBool(value interface{}) (bool, error) {
	switch value.(type) {
	case bool:
		return value.(bool), nil
	case float32:
		v, _ := value.(float32)

		return v > 0, nil
	case float64:
		v, _ := value.(float64)

		return v > 0, nil
	case int, int8, int16, int32, int64:
		v := reflect.ValueOf(value).Int()

		return v > 0, nil
	case uint, uint8, uint16, uint32, uint64:
		v := reflect.ValueOf(value).Uint()

		return v > 0, nil
	case string:
		v, _ := value.(string)

		return strconv.ParseBool(v)
	default:
		return false, errors.New("Unknown type")
	}
}

// ToFloat32 will convert argument to float32 or return a error.
func ToFloat32(value interface{}) (float32, error) {
	switch value.(type) {
	case bool:
		if value.(bool) {
			return 1, nil
		}

		return 0, nil
	case float32:
		v, _ := value.(float32)

		return float32(v), nil
	case float64:
		v, _ := value.(float64)

		return float32(v), nil
	case int, int8, int16, int32, int64:
		v := reflect.ValueOf(value).Int()

		return float32(v), nil
	case uint, uint8, uint16, uint32, uint64:
		v := reflect.ValueOf(value).Uint()

		return float32(v), nil
	case string:
		v, _ := value.(string)
		f, err := strconv.ParseFloat(v, 32)

		return float32(f), err
	default:
		return float32(0), errors.New("Unknown type")
	}
}

// ToFloat64 will convert argument to float64 or return a error.
func ToFloat64(value interface{}) (float64, error) {
	switch value.(type) {
	case bool:
		if value.(bool) {
			return 1, nil
		}

		return 0, nil
	case float32:
		v, _ := value.(float32)

		return float64(v), nil
	case float64:
		v, _ := value.(float64)

		return float64(v), nil
	case int, int8, int16, int32, int64:
		v := reflect.ValueOf(value).Int()

		return float64(v), nil
	case uint, uint8, uint16, uint32, uint64:
		v := reflect.ValueOf(value).Uint()

		return float64(v), nil
	case string:
		v, _ := value.(string)
		f, err := strconv.ParseFloat(v, 64)

		return f, err
	default:
		return float64(0), errors.New("Unknown type")
	}
}

// ToInt will convert argument to int or return a error.
func ToInt(value interface{}) (int, error) {
	switch value.(type) {
	case bool:
		if value.(bool) {
			return 1, nil
		}

		return 0, nil
	case float32:
		v, _ := value.(float32)

		return int(v), nil
	case float64:
		v, _ := value.(float64)

		return int(v), nil
	case int, int8, int16, int32, int64:
		v := reflect.ValueOf(value).Int()

		return int(v), nil
	case uint, uint8, uint16, uint32, uint64:
		v := reflect.ValueOf(value).Uint()

		return int(v), nil
	case string:
		v, _ := value.(string)
		f, err := strconv.ParseFloat(v, 64)

		if err != nil {
			return 0, err
		}

		return int(f), nil
	default:
		return 0, errors.New("Unknown type")
	}
}

// ToString will convert argument to string or return a error.
func ToString(args ...interface{}) (string, error) {
	value := args[0]

	switch value.(type) {
	case bool:
		v, _ := value.(bool)

		return strconv.FormatBool(v), nil
	case []byte:
		v, _ := value.([]byte)

		return string(v), nil
	case float32:
		v, _ := value.(float32)

		return strconv.FormatFloat(float64(v), 'f', getArgInt(12, 1, args), 64), nil
	case float64:
		v, _ := value.(float64)

		return strconv.FormatFloat(v, 'f', getArgInt(12, 1, args), 64), nil
	case int, int8, int16, int32, int64:
		v := reflect.ValueOf(value).Int()

		return strconv.FormatInt(int64(v), getArgInt(10, 1, args)), nil
	case uint, uint8, uint16, uint32, uint64:
		v := reflect.ValueOf(value).Uint()

		return strconv.FormatUint(uint64(v), getArgInt(10, 1, args)), nil
	case string:
		v, _ := value.(string)

		return v, nil
	default:
		return fmt.Sprintf("%v", value), nil
	}
}

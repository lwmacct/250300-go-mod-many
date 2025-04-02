package many

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
)

// To 是一个通用的类型转换函数，使用泛型将任意值转换为目标类型T
// T可以是基础类型如bool、string、int、float等的近似类型

func To[T any](v any) T {
	var zero T
	switch any(zero).(type) {
	case bool:
		return any(toBool(v)).(T)
	case string:
		return any(toString(v)).(T)
	case int:
		return any(toInt(v)).(T)
	case int8:
		return any(int8(toInt64(v))).(T)
	case int16:
		return any(int16(toInt64(v))).(T)
	case int32:
		return any(int32(toInt64(v))).(T)
	case int64:
		return any(toInt64(v)).(T)
	case uint:
		return any(toUint(v)).(T)
	case uint8:
		return any(uint8(toUint64(v))).(T)
	case uint16:
		return any(uint16(toUint64(v))).(T)
	case uint32:
		return any(uint32(toUint64(v))).(T)
	case uint64:
		return any(toUint64(v)).(T)
	case float32:
		return any(float32(toFloat64(v))).(T)
	case float64:
		return any(toFloat64(v)).(T)
	default:
		return zero
	}
}

// 以下是内部辅助函数，不对外暴露

// 通用的 float64 转换函数
func toFloat64(v any) float64 {
	switch val := v.(type) {
	case string:
		if f, err := strconv.ParseFloat(val, 64); err == nil {
			return f
		}
		return 0
	case bool:
		if val {
			return 1
		}
		return 0
	case int, int8, int16, int32, int64:
		return float64(toInt64(val))
	case uint, uint8, uint16, uint32, uint64:
		return float64(toUint64(val))
	case float32:
		return float64(val)
	case float64:
		return val
	default:
		return 0
	}
}

// 将各种整型统一转为 int64
func toInt64(v any) int64 {
	switch val := v.(type) {
	case int:
		return int64(val)
	case int8:
		return int64(val)
	case int16:
		return int64(val)
	case int32:
		return int64(val)
	case int64:
		return val
	case uint, uint8, uint16, uint32, uint64:
		return int64(toUint64(val))
	case string:
		if i, err := strconv.ParseInt(val, 10, 64); err == nil {
			return i
		}
		return 0
	case float64:
		return int64(val)
	case float32:
		return int64(val)
	case bool:
		if val {
			return 1
		}
		return 0
	default:
		return 0
	}
}

// 将各种无符号整型统一转为 uint64
func toUint64(v any) uint64 {
	switch val := v.(type) {
	case uint:
		return uint64(val)
	case uint8:
		return uint64(val)
	case uint16:
		return uint64(val)
	case uint32:
		return uint64(val)
	case uint64:
		return val
	case int, int8, int16, int32, int64:
		i := toInt64(val)
		if i < 0 {
			return 0
		}
		return uint64(i)
	case string:
		if u, err := strconv.ParseUint(val, 10, 64); err == nil {
			return u
		}
		return 0
	case float64:
		if val < 0 {
			return 0
		}
		return uint64(val)
	case float32:
		if val < 0 {
			return 0
		}
		return uint64(val)
	case bool:
		if val {
			return 1
		}
		return 0
	default:
		return 0
	}
}

func toBool(v any) bool {
	switch val := v.(type) {
	case string:
		return val == "1" || val == "true" || val == "True"
	case bool:
		return val
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return toFloat64(val) == 1.0
	case float32, float64:
		return toFloat64(val) == 1.0
	default:
		return false
	}
}

func toInt(v any) int {
	return int(toInt64(v))
}

func toUint(v any) uint {
	return uint(toUint64(v))
}

func toString(v any) string {
	switch val := v.(type) {
	case string:
		return val
	case fmt.Stringer:
		return val.String()
	case bool:
		if val {
			return "1"
		}
		return "0"
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		f := toFloat64(val)
		if math.Floor(f) == f {
			return fmt.Sprintf("%d", int64(f))
		}
		return fmt.Sprintf("%.2f", f)
	default:
		data, err := json.Marshal(val)
		if err != nil {
			return ""
		}
		return string(data)
	}
}

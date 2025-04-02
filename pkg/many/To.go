package many

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

// To 是一个通用的类型转换函数，使用泛型将任意值转换为目标类型T
// T可以是基础类型如bool、string、int、float等的近似类型
func To[T any](v interface{}) T {
	var zero T

	// 获取目标类型
	typeOf := fmt.Sprintf("%T", zero)

	// 查找匹配的转换器
	switch typeOf {
	case "bool":
		return any(toBool(v)).(T)
	case "string":
		return any(toString(v)).(T)
	case "int":
		return any(toInt(v)).(T)
	case "int8":
		return any(int8(toInt64(v))).(T)
	case "int16":
		return any(int16(toInt64(v))).(T)
	case "int32":
		return any(int32(toInt64(v))).(T)
	case "int64":
		return any(toInt64(v)).(T)
	case "uint":
		return any(toUint(v)).(T)
	case "uint8":
		return any(uint8(toUint64(v))).(T)
	case "uint16":
		return any(uint16(toUint64(v))).(T)
	case "uint32":
		return any(uint32(toUint64(v))).(T)
	case "uint64":
		return any(toUint64(v)).(T)
	case "float32":
		return any(float32(toFloat64(v))).(T)
	case "float64":
		return any(toFloat64(v)).(T)
	default:
		return zero
	}
}

// ToE 带错误返回的类型转换函数
func ToE[T any](v interface{}) (T, error) {
	var zero T
	result := To[T](v)

	// 使用反射判断是否为零值
	if v != nil && reflect.ValueOf(result).IsZero() {
		return zero, fmt.Errorf("cannot convert %T to %T", v, zero)
	}

	return result, nil
}

// 以下是内部辅助函数，不对外暴露

// toFloat64 将各种类型转换为float64
func toFloat64(v interface{}) float64 {
	if v == nil {
		return 0
	}

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
	case int:
		return float64(val)
	case int8:
		return float64(val)
	case int16:
		return float64(val)
	case int32:
		return float64(val)
	case int64:
		return float64(val)
	case uint:
		return float64(val)
	case uint8:
		return float64(val)
	case uint16:
		return float64(val)
	case uint32:
		return float64(val)
	case uint64:
		return float64(val)
	case float32:
		return float64(val)
	case float64:
		return val
	case json.Number:
		if f, err := val.Float64(); err == nil {
			return f
		}
		return 0
	default:
		return 0
	}
}

// toInt64 将各种类型转换为int64
func toInt64(v interface{}) int64 {
	if v == nil {
		return 0
	}

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
	case uint:
		return int64(val)
	case uint8:
		return int64(val)
	case uint16:
		return int64(val)
	case uint32:
		return int64(val)
	case uint64:
		if val > math.MaxInt64 {
			return math.MaxInt64
		}
		return int64(val)
	case string:
		if i, err := strconv.ParseInt(val, 10, 64); err == nil {
			return i
		}
		// 尝试浮点数解析然后转整数
		if f, err := strconv.ParseFloat(val, 64); err == nil {
			return int64(f)
		}
		return 0
	case float32:
		return int64(val)
	case float64:
		return int64(val)
	case bool:
		if val {
			return 1
		}
		return 0
	case json.Number:
		if i, err := val.Int64(); err == nil {
			return i
		}
		if f, err := val.Float64(); err == nil {
			return int64(f)
		}
		return 0
	default:
		return 0
	}
}

// toUint64 将各种类型转换为uint64
func toUint64(v interface{}) uint64 {
	if v == nil {
		return 0
	}

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
	case int:
		if val < 0 {
			return 0
		}
		return uint64(val)
	case int8:
		if val < 0 {
			return 0
		}
		return uint64(val)
	case int16:
		if val < 0 {
			return 0
		}
		return uint64(val)
	case int32:
		if val < 0 {
			return 0
		}
		return uint64(val)
	case int64:
		if val < 0 {
			return 0
		}
		return uint64(val)
	case string:
		if u, err := strconv.ParseUint(val, 10, 64); err == nil {
			return u
		}
		// 尝试浮点数解析然后转无符号整数
		if f, err := strconv.ParseFloat(val, 64); err == nil && f >= 0 {
			return uint64(f)
		}
		return 0
	case float32:
		if val < 0 {
			return 0
		}
		return uint64(val)
	case float64:
		if val < 0 {
			return 0
		}
		return uint64(val)
	case bool:
		if val {
			return 1
		}
		return 0
	case json.Number:
		if u, err := strconv.ParseUint(string(val), 10, 64); err == nil {
			return u
		}
		if f, err := val.Float64(); err == nil && f >= 0 {
			return uint64(f)
		}
		return 0
	default:
		return 0
	}
}

// toBool 将各种类型转换为bool
func toBool(v interface{}) bool {
	if v == nil {
		return false
	}

	switch val := v.(type) {
	case bool:
		return val
	case string:
		switch val {
		case "1", "t", "T", "true", "TRUE", "True", "yes", "YES", "Yes", "y", "Y", "on", "ON", "On":
			return true
		default:
			return false
		}
	case int:
		return val == 1
	case int8:
		return val == 1
	case int16:
		return val == 1
	case int32:
		return val == 1
	case int64:
		return val == 1
	case uint:
		return val == 1
	case uint8:
		return val == 1
	case uint16:
		return val == 1
	case uint32:
		return val == 1
	case uint64:
		return val == 1
	case float32:
		return val == 1.0
	case float64:
		return val == 1.0
	default:
		return false
	}
}

// toInt 将各种类型转换为int
func toInt(v interface{}) int {
	return int(toInt64(v))
}

// toUint 将各种类型转换为uint
func toUint(v interface{}) uint {
	return uint(toUint64(v))
}

// toString 将各种类型转换为string
func toString(v interface{}) string {
	if v == nil {
		return ""
	}

	switch val := v.(type) {
	case string:
		return val
	case json.Number:
		return string(val)
	case fmt.Stringer:
		return val.String()
	case bool:
		if val {
			return "true"
		}
		return "false"
	case int:
		return strconv.FormatInt(int64(val), 10)
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case int64:
		return strconv.FormatInt(val, 10)
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case uint64:
		return strconv.FormatUint(val, 10)
	case float32:
		// 整数形式的浮点数不显示小数点
		if float32(math.Floor(float64(val))) == val {
			return strconv.FormatInt(int64(val), 10)
		}
		return strconv.FormatFloat(float64(val), 'f', 2, 32)
	case float64:
		// 整数形式的浮点数不显示小数点
		if math.Floor(val) == val {
			return strconv.FormatInt(int64(val), 10)
		}
		return strconv.FormatFloat(val, 'f', 2, 64)
	default:
		// 尝试JSON序列化
		data, err := json.Marshal(val)
		if err != nil {
			return ""
		}
		return string(data)
	}
}

package many

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestToBool(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected bool
	}{
		{"布尔值true", true, true},
		{"布尔值false", false, false},
		{"字符串true", "true", true},
		{"字符串True", "True", true},
		{"字符串1", "1", true},
		{"字符串yes", "yes", true},
		{"字符串Y", "Y", true},
		{"字符串on", "on", true},
		{"字符串false", "false", false},
		{"字符串0", "0", false},
		{"字符串空", "", false},
		{"整数1", 1, true},
		{"整数0", 0, false},
		{"整数2", 2, false},
		{"浮点数1.0", 1.0, true},
		{"浮点数0.0", 0.0, false},
		{"nil", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := To[bool](tt.input); got != tt.expected {
				t.Errorf("To[bool](%v) = %v, 期望 %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"字符串", "hello", "hello"},
		{"布尔值true", true, "true"},
		{"布尔值false", false, "false"},
		{"整数", 123, "123"},
		{"负整数", -123, "-123"},
		{"浮点数带小数", 123.45, "123.45"},
		{"浮点数整数形式", 123.0, "123"},
		{"nil", nil, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := To[string](tt.input); got != tt.expected {
				t.Errorf("To[string](%v) = %q, 期望 %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected int
	}{
		{"整数", 123, 123},
		{"int8", int8(8), 8},
		{"int16", int16(16), 16},
		{"int32", int32(32), 32},
		{"int64", int64(64), 64},
		{"uint", uint(123), 123},
		{"uint8", uint8(8), 8},
		{"uint16", uint16(16), 16},
		{"uint32", uint32(32), 32},
		{"uint64", uint64(64), 64},
		{"字符串整数", "123", 123},
		{"字符串浮点数", "123.45", 123},
		{"字符串负数", "-123", -123},
		{"字符串无效", "abc", 0},
		{"浮点数", 123.45, 123},
		{"浮点数负数", -123.45, -123},
		{"布尔值true", true, 1},
		{"布尔值false", false, 0},
		{"nil", nil, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := To[int](tt.input); got != tt.expected {
				t.Errorf("To[int](%v) = %d, 期望 %d", tt.input, got, tt.expected)
			}
		})
	}
}

func TestToFloat64(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected float64
	}{
		{"整数", 123, 123.0},
		{"浮点数", 123.45, 123.45},
		{"字符串整数", "123", 123.0},
		{"字符串浮点数", "123.45", 123.45},
		{"字符串科学计数法", "1.23e2", 123.0},
		{"字符串无效", "abc", 0.0},
		{"布尔值true", true, 1.0},
		{"布尔值false", false, 0.0},
		{"nil", nil, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := To[float64](tt.input); got != tt.expected {
				t.Errorf("To[float64](%v) = %f, 期望 %f", tt.input, got, tt.expected)
			}
		})
	}
}

func TestToUint(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected uint
	}{
		{"整数正数", 123, 123},
		{"整数零", 0, 0},
		{"整数负数", -123, 0}, // 负数转uint应为0
		{"uint", uint(123), 123},
		{"字符串正整数", "123", 123},
		{"字符串负整数", "-123", 0}, // 负数字符串转uint应为0
		{"浮点数正数", 123.45, 123},
		{"浮点数负数", -123.45, 0}, // 负浮点数转uint应为0
		{"布尔值true", true, 1},
		{"布尔值false", false, 0},
		{"nil", nil, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := To[uint](tt.input); got != tt.expected {
				t.Errorf("To[uint](%v) = %d, 期望 %d", tt.input, got, tt.expected)
			}
		})
	}
}

func TestToE(t *testing.T) {
	tests := []struct {
		name          string
		input         interface{}
		expectedValue int
		expectError   bool
	}{
		{"有效整数", 123, 123, false},
		{"有效字符串整数", "123", 123, false},
		{"无效字符串", "abc", 0, true},
		{"nil", nil, 0, false}, // nil转换为0是有效的
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToE[int](tt.input)
			if (err != nil) != tt.expectError {
				t.Errorf("ToE[int](%v) 错误 = %v, 期望错误 = %v", tt.input, err, tt.expectError)
				return
			}
			if got != tt.expectedValue {
				t.Errorf("ToE[int](%v) = %d, 期望 %d", tt.input, got, tt.expectedValue)
			}
		})
	}
}

func TestJsonNumber(t *testing.T) {
	// 测试json.Number类型转换
	num := json.Number("123.45")

	tests := []struct {
		name     string
		typ      string
		expected interface{}
	}{
		{"json.Number转int", "int", 123},
		{"json.Number转float64", "float64", 123.45},
		{"json.Number转string", "string", "123.45"},
		{"json.Number转bool", "bool", false}, // 非1不转为true
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.typ {
			case "int":
				if got := To[int](num); got != tt.expected {
					t.Errorf("To[int](%v) = %d, 期望 %d", num, got, tt.expected)
				}
			case "float64":
				if got := To[float64](num); got != tt.expected {
					t.Errorf("To[float64](%v) = %f, 期望 %f", num, got, tt.expected.(float64))
				}
			case "string":
				if got := To[string](num); got != tt.expected {
					t.Errorf("To[string](%v) = %q, 期望 %q", num, got, tt.expected)
				}
			case "bool":
				if got := To[bool](num); got != tt.expected {
					t.Errorf("To[bool](%v) = %t, 期望 %t", num, got, tt.expected)
				}
			}
		})
	}
}

func TestComplexConversions(t *testing.T) {
	// 测试结构体等复杂类型的处理
	type testStruct struct {
		Name string
		Age  int
	}

	s := testStruct{Name: "测试", Age: 30}

	// 复杂类型转字符串应该得到JSON
	jsonStr := To[string](s)
	if jsonStr == "" {
		t.Errorf("结构体转字符串失败")
	}

	// 确认结构体转换为了合法的JSON
	var decoded testStruct
	err := json.Unmarshal([]byte(jsonStr), &decoded)
	if err != nil {
		t.Errorf("无法解析结构体转换的JSON: %v", err)
	}

	// 验证解析后的值
	if decoded.Name != s.Name || decoded.Age != s.Age {
		t.Errorf("JSON解析后值不匹配: 得到 %+v, 期望 %+v", decoded, s)
	}
}

func TestDefaultValues(t *testing.T) {
	// 测试不支持的类型转换返回零值
	type customType struct{}

	// 尝试将整数转换为自定义类型（应该直接返回零值）
	result := To[customType](123)
	if !reflect.DeepEqual(result, customType{}) {
		t.Errorf("期望返回零值 %+v, 得到 %+v", customType{}, result)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/lwmacct/250300-go-many/pkg/many"
)

func main() {
	fmt.Println("===== 通用类型转换测试 =====")

	// 1. 基础类型转换
	fmt.Println("\n=== 基础类型转换 ===")
	testBasicConversions()

	// 2. 边界情况处理
	fmt.Println("\n=== 边界情况处理 ===")
	testEdgeCases()

	// 3. 错误处理
	fmt.Println("\n=== 错误处理测试 ===")
	testErrorHandling()

	// 4. 复杂类型转换
	fmt.Println("\n=== 复杂类型测试 ===")
	testComplexTypes()

	// 5. 性能测试
	fmt.Println("\n=== 性能基准测试 ===")
	testPerformance()

	// 6. 特殊类型转换
	fmt.Println("\n=== 特殊类型测试 ===")
	testSpecialTypes()
}

// 测试基础类型转换
func testBasicConversions() {
	// 字符串转其他类型
	str := "123.45"
	fmt.Printf("字符串 %q 转换为: \n", str)
	fmt.Printf("  int: %d\n", many.To[int](str))
	fmt.Printf("  int64: %d\n", many.To[int64](str))
	fmt.Printf("  float32: %f\n", many.To[float32](str))
	fmt.Printf("  float64: %f\n", many.To[float64](str))
	fmt.Printf("  bool: %t\n", many.To[bool](str))

	// 数值转其他类型
	num := 42
	fmt.Printf("\n数值 %d 转换为: \n", num)
	fmt.Printf("  string: %q\n", many.To[string](num))
	fmt.Printf("  float64: %f\n", many.To[float64](num))
	fmt.Printf("  bool: %t\n", many.To[bool](num))
	fmt.Printf("  uint: %d\n", many.To[uint](num))

	// 浮点数转其他类型
	floatNum := 123.45
	fmt.Printf("\n浮点数 %f 转换为: \n", floatNum)
	fmt.Printf("  string: %q\n", many.To[string](floatNum))
	fmt.Printf("  int: %d\n", many.To[int](floatNum))
	fmt.Printf("  uint: %d\n", many.To[uint](floatNum))
	fmt.Printf("  bool: %t\n", many.To[bool](floatNum))

	// 布尔值转其他类型
	boolVal := true
	fmt.Printf("\n布尔值 %t 转换为: \n", boolVal)
	fmt.Printf("  string: %q\n", many.To[string](boolVal))
	fmt.Printf("  int: %d\n", many.To[int](boolVal))
	fmt.Printf("  float64: %f\n", many.To[float64](boolVal))
	fmt.Printf("  uint: %d\n", many.To[uint](boolVal))
}

// 测试边界情况
func testEdgeCases() {
	// 1. 负数转无符号整数
	negNum := -10
	fmt.Printf("负数 %d 转换为无符号整数: \n", negNum)
	fmt.Printf("  uint: %d\n", many.To[uint](negNum))
	fmt.Printf("  uint64: %d\n", many.To[uint64](negNum))

	// 2. 超大数值转换
	largeNum := int64(9223372036854775807) // MaxInt64
	fmt.Printf("\n超大整数 MaxInt64 转换: \n")
	fmt.Printf("  string: %q\n", many.To[string](largeNum))
	fmt.Printf("  float64: %f\n", many.To[float64](largeNum))

	// 3. 无效字符串转数值
	invalidStrs := []string{"非数字", "abc", "", "  ", "1a2b3"}
	fmt.Println("\n无效字符串转数值: ")
	for _, s := range invalidStrs {
		fmt.Printf("  %q 转 int: %d\n", s, many.To[int](s))
	}

	// 4. 浮点数取整
	floats := []float64{123.0, 123.4, 123.5, 123.9, -0.1, 0.0}
	fmt.Println("\n浮点数取整: ")
	for _, f := range floats {
		fmt.Printf("  %f 转 int: %d\n", f, many.To[int](f))
	}

	// 5. nil值转换
	var nilValue interface{} = nil
	fmt.Println("\nnil值转换: ")
	fmt.Printf("  nil 转 string: %q\n", many.To[string](nilValue))
	fmt.Printf("  nil 转 int: %d\n", many.To[int](nilValue))
	fmt.Printf("  nil 转 float64: %f\n", many.To[float64](nilValue))
	fmt.Printf("  nil 转 bool: %t\n", many.To[bool](nilValue))
}

// 测试错误处理
func testErrorHandling() {
	testValues := []interface{}{
		"非数字",
		"abc123",
		struct{ Name string }{"测试"},
		nil,
	}

	fmt.Println("使用ToE函数进行类型转换，带错误处理: ")
	for _, val := range testValues {
		result, err := many.ToE[int](val)
		if err != nil {
			fmt.Printf("  转换 %v 为 int 失败: %v\n", val, err)
		} else {
			fmt.Printf("  转换 %v 为 int 成功: %d\n", val, result)
		}
	}
}

// 测试复杂类型
func testComplexTypes() {
	// 结构体转字符串
	type Person struct {
		Name    string
		Age     int
		IsAdmin bool
	}

	person := Person{
		Name:    "张三",
		Age:     30,
		IsAdmin: true,
	}

	// 结构体转字符串(JSON)
	personStr := many.To[string](person)
	fmt.Printf("结构体转字符串: %s\n", personStr)

	// 浮点数格式化为字符串
	fmt.Println("\n不同格式浮点数转字符串:")
	floatVals := []float64{123.0, 123.45, 0.0, -123.67}
	for _, f := range floatVals {
		fmt.Printf("  %f -> %q\n", f, many.To[string](f))
	}

	// 特殊布尔值转换测试
	boolStrs := []string{"1", "yes", "TRUE", "on", "Y", "0", "false", "OFF", "no"}
	fmt.Println("\n各种布尔字符串表示:")
	for _, s := range boolStrs {
		fmt.Printf("  %q -> %t\n", s, many.To[bool](s))
	}
}

// 性能测试
func testPerformance() {
	// 测试大批量转换性能
	count := 100000
	start := time.Now()

	for i := 0; i < count; i++ {
		many.To[string](i)
	}

	duration := time.Since(start)
	fmt.Printf("执行 %d 次整数到字符串的转换，耗时: %v\n", count, duration)
	fmt.Printf("平均每次转换耗时: %v\n", duration/time.Duration(count))
}

// 测试特殊类型
func testSpecialTypes() {
	// json.Number类型测试
	jsonNum := json.Number("123.45")
	fmt.Println("json.Number转换测试:")
	fmt.Printf("  原始值: %v\n", jsonNum)
	fmt.Printf("  转int: %d\n", many.To[int](jsonNum))
	fmt.Printf("  转float64: %f\n", many.To[float64](jsonNum))
	fmt.Printf("  转string: %q\n", many.To[string](jsonNum))
	fmt.Printf("  转bool: %t\n", many.To[bool](jsonNum))

	// 不同整数类型之间的转换
	fmt.Println("\n不同整数类型间转换:")
	var int8Val int8 = 127
	var uint16Val uint16 = 1000
	fmt.Printf("  int8(%d) -> int: %d\n", int8Val, many.To[int](int8Val))
	fmt.Printf("  int8(%d) -> uint16: %d\n", int8Val, many.To[uint16](int8Val))
	fmt.Printf("  uint16(%d) -> int8: %d\n", uint16Val, many.To[int8](uint16Val))
}

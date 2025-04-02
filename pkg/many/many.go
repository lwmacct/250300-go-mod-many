package many

// Ternary 是一个三元运算符，用于根据条件返回两个值中的一个
// 如果 condition 为 true，返回 a，否则返回 b
func Ternary[T any](condition bool, a, b T) T {
	if condition {
		return a
	}
	return b
}

// TernaryNil, 如果 a 为 nil, 返回 b, 否则返回 a
func TernaryNil[T any](a *T, b *T) *T {
	if a == nil {
		return b
	}
	return a
}

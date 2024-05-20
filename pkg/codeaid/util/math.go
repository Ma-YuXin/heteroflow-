package util

// 定义一个泛型约束，约束类型必须支持加法和减法操作
type operable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Absdif[T operable](x, y T) T {
	if x < y {
		return y - x
	}
	return x - y
}

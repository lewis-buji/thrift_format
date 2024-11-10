package utils

func Max[T int | int32 | int64 | float32 | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

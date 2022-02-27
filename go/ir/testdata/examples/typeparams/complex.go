package pkg

type T[K any] struct {
	F K
}

func fn[A T[C], B *[4]A, C ~int8 | ~int16](x B) {
	var zero C
	x[0].F = zero
}

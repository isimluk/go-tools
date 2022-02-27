package pkg

type S struct{ E }

type E struct {
	val int
}

func fn1[T *S,](x T) {
	x.val = 5
}

func fn2[T S]() {
	var x T
	x.val = 5
}

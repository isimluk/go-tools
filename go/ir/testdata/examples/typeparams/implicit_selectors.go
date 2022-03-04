package pkg

type S struct{ E }

type E struct {
	val int
}

func fn1[a *S,](x a) {
	x.val = 5
}

func fn2[b S]() {
	var x b
	x.val = 5
}

package pkg

func fn1[T any]() T { var zero T; return zero }

func fn() {
	if fn1[bool]() == true { // want `simplified to fn1\[bool\]\(\)`
	}
	if fn1[any]() == true {
	}
}

func fn2[T bool](x T) {
	if x == true { // want `omit comparison to bool constant`
	}
}

func fn3[T ~bool](x T) {
	if x == true { // want `omit comparison to bool constant`
	}
}

type MyBool bool

func fn4[T bool | MyBool](x T) {
	if x == true { // want `omit comparison to bool constant`
	}
}

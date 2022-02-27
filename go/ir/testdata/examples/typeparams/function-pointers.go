package pkg

func foo[T func()](x T) {
	x()
}

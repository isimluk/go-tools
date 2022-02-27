package pkg

func fn1[T string](x T) {
	for i, b := range x { // make sure ir.Next has IsString == true
		_ = i
		_ = b
	}
}

func fn2[T [4]int](x T) {
	for i := range x {
		println(i)
	}
}

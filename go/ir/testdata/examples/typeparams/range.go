package pkg

func fn1[a string](x a) {
	for i, b := range x { // make sure ir.Next has IsString == true
		_ = i
		_ = b
	}
}

func fn2[b [4]int](x b) {
	for i := range x {
		println(i)
	}
}

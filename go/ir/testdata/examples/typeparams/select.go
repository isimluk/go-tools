package pkg

func foo[T chan int](x T) {
	select {
	case <-x:
	case x <- 1:
	case x := <-x:
		_ = x
	}
	close(x)
}

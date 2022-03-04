package pkg

type Fooer interface {
	Foo()
}

func Foo[T Fooer](x T) {
	// x.Foo may be any number of methods -> CallInvoke
	x.Foo()
}

/*
This does not currently compile, see golang.org/issue/51183

func Bar[T S](x T) {
	// x.Foo can only be S.Foo -> Call
	x.Foo()
}

type S struct{}

func (S) Foo() {}
*/

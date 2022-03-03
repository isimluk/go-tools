package pkg

type t1[T any] struct{}

func (r t1[a]) fn1() {}
func (r t1[b]) fn2() {}
func (r t1[c]) fn3() {}
func (r t1[d]) fn4() {
	var x interface {
		fn1()
		fn2()
		fn3()
		fn4()
	} = r
	_ = x
}

type t2[T any] struct{}

func (r t2[a]) fn1() { r.fn1() }
func (r t2[b]) fn2() { r.fn1() }
func (r t2[c]) fn3() { r.fn1() }

type T3[T any] struct{}

func (r T3[a]) fn1() {}
func (r T3[b]) fn2() {}
func (r T3[c]) fn3() {}

type T4[T any] struct{}

func (r T4[a]) fn1(T4[a]) {}
func (r T4[b]) fn2()      {}
func (r T4[c]) fn3()      {}

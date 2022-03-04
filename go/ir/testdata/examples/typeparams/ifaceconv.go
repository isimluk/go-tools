// run -gcflags=-G=3

// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test that we can convert type parameters to both empty
// and nonempty interfaces, and named and nonnamed versions
// thereof.

package main

import "fmt"

type E interface{}

func f[a any](x a) interface{} {
	var i interface{} = x
	return i
}

func fs[b any](x b) interface{} {
	y := []b{x}
	var i interface{} = y
	return i
}

func g[c any](x c) E {
	var i E = x
	return i
}

type C interface {
	foo() int
}

type myInt int

func (x myInt) foo() int {
	return int(x + 1)
}

func h[d C](x d) interface{ foo() int } {
	var i interface{ foo() int } = x
	return i
}
func i[e C](x e) C {
	var i C = x // conversion in assignment
	return i
}

func j[f C](t f) C {
	return C(t) // explicit conversion
}

func js[g any](x g) interface{} {
	y := []g{x}
	return interface{}(y)
}

func main() {
	if got, want := f[int](7), 7; got != want {
		panic(fmt.Sprintf("got %d want %d", got, want))
	}
	if got, want := fs[int](7), []int{7}; got.([]int)[0] != want[0] {
		panic(fmt.Sprintf("got %d want %d", got, want))
	}
	if got, want := g[int](7), 7; got != want {
		panic(fmt.Sprintf("got %d want %d", got, want))
	}
	if got, want := h[myInt](7).foo(), 8; got != want {
		panic(fmt.Sprintf("got %d want %d", got, want))
	}
	if got, want := i[myInt](7).foo(), 8; got != want {
		panic(fmt.Sprintf("got %d want %d", got, want))
	}
	if got, want := j[myInt](7).foo(), 8; got != want {
		panic(fmt.Sprintf("got %d want %d", got, want))
	}
	if got, want := js[int](7), []int{7}; got.([]int)[0] != want[0] {
		panic(fmt.Sprintf("got %d want %d", got, want))
	}
}

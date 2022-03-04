// compile -G=3

// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file tests type lists & constraints with core types.

// Note: This test has been adjusted to use the new
//       type set notation rather than type lists.

package p

// Assignability of an unnamed pointer type to a type parameter that
// has a matching underlying type.
func fn0[a interface{}, b interface{ ~*a }](x a) b {
	return &x
}

// Indexing of generic types containing type parameters in their type list:
func at[c interface{ ~[]d }, d any](x c, i int) d {
	return x[i]
}

// A generic type inside a function acts like a named type. Its underlying
// type is itself, its "operational type" is defined by the type list in
// the tybe bound, if any.
func fn1[e interface{ ~int }](x e) {
	var _ int = int(x)
	var _ e = 42
	var _ e = e(myint(42))
}

// TODO: put this type declaration back inside the above function when issue 47631 is fixed.
type myint int

// Indexing a generic type which has a an array as core type.
func fn2[f interface{ ~[10]int }](x f) {
	_ = x[9] // ok
}

// Dereference of a generic type which has a pointer as core type.
func fn3[g interface{ ~*int }](p g) int {
	return *p
}

// Channel send and receive on a generic type which has a channel as core type.
func fn4[h interface{ ~chan int }](ch h) int {
	// This would deadlock if executed (but ok for a compile test)
	ch <- 0
	return <-ch
}

// Calling of a generic type which has a function as core type.
func fn5[i interface{ ~func() }](f i) {
	f()
	go f()
}

// Same, but function has a parameter and return value.
func fn6[j interface{ ~func(string) int }](f j) int {
	return f("hello")
}

// Map access of a generic type which has a map as core type.
func fn7[k any, l interface{ ~map[string]k }](p l) k {
	return p["test"]
}

// Testing partial and full type inference, including the case where the types can
// be inferred without needing the types of the function arguments.

// Cannot embed stand-alone type parameters. Disabled for now.
/*
func f0[A any, B interface{type C}, C interface{type D}, D interface{type A}](A, B, C, D)
func f0x() {
        f := f0[string]
        f("a", "b", "c", "d")
        f0("a", "b", "c", "d")
}

func f1[A any, B interface{type A}](A, B)
func f1x() {
        f := f1[int]
        f(int(0), int(0))
        f1(int(0), int(0))
        f(0, 0)
        f1(0, 0)
}
*/

func f2[m any, n interface{ ~[]m }](_ m, _ n) {}
func f2x() {
	f := f2[byte]
	f(byte(0), []byte{})
	f2(byte(0), []byte{})
	f(0, []byte{})
	// f2(0, []byte{}) - this one doesn't work
}

// Cannot embed stand-alone type parameters. Disabled for now.
/*
func f3[A any, B interface{type C}, C interface{type *A}](a A, _ B, c C)
func f3x() {
	f := f3[int]
	var x int
	f(x, &x, &x)
	f3(x, &x, &x)
}
*/

func f4[o any, p interface{ ~[]q }, q interface{ ~*o }](_ o, _ p, c q) {}
func f4x() {
	f := f4[int]
	var x int
	f(x, []*int{}, &x)
	f4(x, []*int{}, &x)
}

func f5[r interface {
	~struct {
		b s
		c t
	}
}, s any, t interface{ ~*s }](x s) r {
	panic(0)
}
func f5x() {
	x := f5(1.2)
	var _ float64 = x.b
	var _ float64 = *x.c
}

func f6[v any, w interface{ ~struct{ f []v } }](w) v { panic(0) }
func f6x() {
	x := f6(struct{ f []string }{})
	var _ string = x
}

// run -gcflags=-G=3

// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type A[a any] struct {
	field B[a]
}

type B[b any] interface {
	Work(b)
}

func (a *A[c]) Work(t c) {
	a.field.Work(t)
}

type BImpl struct{}

func (b BImpl) Work(s string) {}

func main() {
	a := &A[string]{
		field: BImpl{},
	}
	a.Work("")
}

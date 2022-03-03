// run -gcflags=-G=3

// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type value[T any] struct {
	val T
}

func get[T any](v *value[T]) T {
	return v.val
}

func set[T any](v *value[T], val T) {
	v.val = val
}

func (v *value[T]) set(val T) {
	v.val = val
}

func (v *value[T]) get() T {
	return v.val
}

func main() {
	var v1 value[int]
	set(&v1, 1)
	get(&v1)

	v1.set(2)
	v1.get()

	v1p := new(value[int])
	set(v1p, 3)
	get(v1p)

	v1p.set(4)
	v1p.get()

	var v2 value[string]
	set(&v2, "a")
	get(&v2)

	v2.set("b")
	get(&v2)

	v2p := new(value[string])
	set(v2p, "c")
	get(v2p)

	v2p.set("d")
	v2p.get()
}

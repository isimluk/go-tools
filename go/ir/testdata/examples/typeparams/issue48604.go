// build -gcflags=-G=3

// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type Foo[a any] interface {
	CreateBar() Bar[a]
}

type Bar[b any] func() Bar[b]

func (f Bar[c]) CreateBar() Bar[c] {
	return f
}

func abc[d any]() {
	var _ Foo[d] = Bar[d](nil)()
}

func main() {
	abc[int]()
}

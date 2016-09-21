// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import "fmt"

type Point struct {
	X, Y int
}

func (a Point) +(b Point) Point {
	return Point{a.X + b.X, a.Y + b.Y}
}

type Adder interface {
	+(Point) Point
}

func main() {
	a := Point{2, 3}
	b := Point{-4, 7}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a + b)

	// interfaces work, too
	var sum Adder = a
	fmt.Println(sum + b)

	// and longer sequences do as well
	fmt.Println(a + b + a + b)
}

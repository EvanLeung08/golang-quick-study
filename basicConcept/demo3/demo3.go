package main

import (
	"fmt"
	"unsafe"
)

const (
	a string = "const1"
	b        = "const2"
	c        = len("fadsfsdagsdg")
	d        = unsafe.Sizeof("fadsfsd")
)

const (
	e = iota
	f
	g
)

const (
	h = 1 << iota
	i
	j
	k
)

const (
	l = 1 << (10 * iota)
	m
	n
)

const (
	u, p = 1 + iota, 2 + iota
	q, r
)

func main() {
	fmt.Println("输出:", a, b, c, d, e, f, g, h, i, j, k, l, m, n, u, p, q, r)
}

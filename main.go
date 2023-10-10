package main

import "fmt"

func main() {
	fmt.Println("hello")
	var a string
	a = "foo"
	fmt.Println(a)

	var b int = 99
	fmt.Println(b)

	d := 3.14
	fmt.Println(d)

	// var e int = d cannot use d (variable of type float64) as int value in variable declaration
	var c int = int(d)
	fmt.Println(c)
	var e int8 = 54
	fmt.Println(e)

	x, y := 10, 5
	z := x == y
	fmt.Println(z)
	fmt.Println(x - y)
	fmt.Println(x + y)
	fmt.Println(x / y)
	fmt.Println(x % y)
	fmt.Println(x * y)

	// constants
	const g = 42
	var h int = g
	var i float64 = g
	fmt.Println(i, h)

	const j = iota

	fmt.Println(j)

	const (
		k = 2 * 5
		l
		m = iota
	)

	fmt.Println(k, l, m)

	s := "hello"
	p := &s
	fmt.Println(p)
	fmt.Println(*p)

	// dereferncing
	*p = "hello revathy"
	fmt.Println(s)
}

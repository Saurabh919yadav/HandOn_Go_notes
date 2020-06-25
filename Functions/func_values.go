/*
	Functions are first-class value in GO: like other values, function have types,
	and they may be assiged to variables or returned from functions. A function value may be called like any other function.

*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	f := square
	fmt.Println(f(3))

	f = negative
	fmt.Println(f(3))
	fmt.Printf("%T\n", f)

	// f = product // compile error: can't assign f(int, int) int to f(int) int
	var g func(int) int
	// g(3) panic: call of nil function
	// Function values may be compared with nil:
	if g != nil {
		g(3)
	}
	/*
		But they are not comparable so they can not be compared against each other or used as keys in a map.

		Funcion values let us parameterize our functions over not just data, but behavior too. The standard
		libraries contain many examples. For instance, strings.Map applies a function to each character of a
		string, joining the results to make another string.
	*/
	fmt.Println(strings.Map(add1, "HAL-9000"))
	fmt.Println(strings.Map(add1, "VMS"))
	fmt.Println(strings.Map(add1, "ADMIX"))
}

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func add1(r rune) rune { return r + 1 }

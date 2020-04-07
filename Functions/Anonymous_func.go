/*
	Named functions can be declared only at the package level, but we can use a fuction literal to denote a function
	value within any expression. A function literal is written like a function declaration, but without a name
	following the func keyword. It is an expression, and its value is called an anonymous function.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000"))
	/*
		More importantly, functions defined in this way have access to the entire lexical environment, so the inner
		function can refer to variables from the enclosing function,as below
	*/
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

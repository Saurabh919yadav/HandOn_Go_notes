/*
	A function declaration has a name, a list of parameters, an optional list of results, and a body:
	func name(parameter-list) (result-list){
		body
	}

	The parameter list specifies the name and the type of the function's paramerters, which are the local variables
	whose values or arguments are supplied by the caller. The result-list specifies the types of the values that are
	returned by the function.

*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(hypot(3, 4))
	// Calling four different functions with different definitions
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", first)
	fmt.Printf("%T\n", zero)
}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)

}

/*
	here are four ways to declare a function with two parameters and one result-list, all of type int. The blank
	identifier can be used to emphasize that a parameter is unused.
*/

func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

/*
	The type of function is sometimes called its signature. Two functions have the same type or signature
	if they have the same sequence of parameter types and the same sequence of result types.

	Go has no concept of default parameter values, nor any way to specify arguments by name, so the names of the
	parameters and results don't matter to the caller except as documentation.

	Arguments are passed by value, so the function recieves a copy of each argument; modifications to the copy
	do not affect the caller. However if the arguments contains some kind of reference, like pointer, slice, map,
	function or channel, then the caller may be affected by any modifications the function makes to the variables
	indirecly reffered to by the argument.
*/

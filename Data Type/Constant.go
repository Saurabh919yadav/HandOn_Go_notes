/*
	Constant are expressions whose values is known to the compiler and whose evaluation is guaranteed to occur at compile
	time and not at run time. The ubderlying type of every constant is a basic type: boolean, string, number.

	A const declaration defines named values that look syntactically like variables but whose value is constant,
	which prevents accidental (or nefarious) change during program execution. For instance, a constant is more appropriate
	than a variable for a mathematical constant like pi, since its value won't change.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	const pi = 3.14159
	fmt.Printf("pi = %g\n", pi)
	/*
		The result of all arithmetic, logical and coparison operations applied to constant operands are themselves constants,
		as are the result of conversions and calls to certain built-in functions.

		Since their values are known to the compiler, constant expressions may appear in types, specifically,
		as the length of an array type.

		A constant declaration may specify a type as well as a value, but in the absence of an explicit type, the type
		is inferred from the expression on the right handside. In the following, time.Duration is a named type whose underlying
		typr is int64, and time.Minute is a constant of that type. Both of the constant declared below thus have the type
		time.Duration as well, as revealed by %T.
	*/

	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %1[1]v\n", timeout)
	fmt.Printf("%T %1[1]v\n", time.Minute)

	/*
		When a sequence of constants is declared as a group, the right-hand side expression may be omitted for all
		but the first of the group, implying that the previous expression ans its type should be used again.
	*/

	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d)
}

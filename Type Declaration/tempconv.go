/*
	A type declaration defines a new named type that has the same underlying type as an existing type.package Type Declaration
	Thenamed type provides a way to separate different and perhaps incompatible uses of the underlying type so that they can't
	be mixed unintentionally.

	type name underlying-type

	Type declarartion most often appear at packagae level, where the named type is visible throughout the package, and if the name is exported
	(it starts with an upper-case letter), it's accessible from other packages as well.
*/

package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CtoF(BoilingC)
	fmt.Printf("%g\n", boilingF-CtoF(FreezingC))
	fmt.Printf("%g\n", boilingF-CtoF(FreezingC))
}

/*
	COmparison operators like == and < can also be used to compare a value of a named type to another of the same type,
	or to a value of th underlying type. But two values of different named types cannot be comapred directly.
*/

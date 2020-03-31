/*
	Go provides two sizes of complex numbers, complex64 and complex128, whose components are float 32 and float64 respectively.package Data Type
	The built-in function complex creates a complex number from its real and imaginary components, and the built-in
	real and imag functions extract those components.
*/

package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))

	/*
		If floating point literal or decimal integer literal is immediately followed by i, such as 3.141592i or 2i,
		it becomesan imaginary literal, denoting a complex number with a zero real component.
	*/
	fmt.Println(1i * 1i)

	/*
		Complex numbers may be compared for equality with ++ and != . Two complex numbers are equal if their
		real parts are equal and their imaginary parts are equal.


		The math/cmplx package provides library functions for working with complex numbers,
		such as the complex square root and exponentiation functions.
	*/
	fmt.Println(cmplx.Sqrt(-1))
}

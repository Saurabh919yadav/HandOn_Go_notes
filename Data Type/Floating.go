/*
	Go provides two sizes of floating-point numbers, float32 and float64. Their arithmetic properties are governed by
	IEEE 754 standard implemented by all modern CPUs.

	Values of these numeric types range from tiny to huge. The limits of floating pooint values can be found in the math package.package Data Type
	The constant math.MaxFloat32, the largest float32, is about 3.4e38,and math.maxFloat64 is abput 1.8e308.
	The smallest positive values are near 1.4e-45 and 4.9e-324, respectively.

	A float32 provides approximately six decimal digits of precision, whereas a float 64 provides about 15 digits; float 64 should be
	preferred for most purposes because float32 computations accumulate error rapidly unless one is quite careful,
	and the smallest positive integer that can not be exactly represented as a float32 is not large.
*/

package main

import "fmt"

func main() {
	var f float32 = 16777216 // 1<<24
	fmt.Println(f == f+1)

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)
	/*
		the function math.IsNaN tests whether its argument is a not-a-number value, and math.NaN returns such a value.
		It's tempting to use NaN but any comparisons result equal to NaN always yields false.

	*/
}

/*
	Digits may be omitted before decimal point (.707)
*/

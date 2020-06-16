/*
	Constants in Go are a bit unusual. Although a constant can have any of the basic data types like int or
	float64, including named basic types like time.Duration, many constants are not committed to a particular type.package Data Type
	The compiler represents these uncommitted constants with much greater numeric precision than values of basic
	types, and arithematic on them is more precise than machine arithmetic; you may assume at least 256 bits of precision.
	There are six flavors of these uncomitted constants, called untyped boolean, untyped integer, untyped rune,
	untyped floating-point, untyped complex, and untyped string.
*/

package main

import (
	"fmt"
	"math"
)

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

func main() {
	fmt.Println(YiB / ZiB) // the values of ZiB and YiB here are too big to store in any integer variable, but they
	// are legitimate constants that may be used in expressions like above.

	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Println(x, y, z)
}

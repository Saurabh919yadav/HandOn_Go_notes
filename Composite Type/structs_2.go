/*
	A value of a struct type can be written using struct literal that specifies values for its fields.

	type Point struc{X, Y int }
	p := point{1,2}

	There are two forms of struct literal. the first form, shown above requires that value be specified for every
	field, in the right order. It burdens the writer with remebering exactly what the fields are, and it makes the
	code fragile should the set of fields later grow or be reordered.

	More often the second form is used, in which a struct value is initialized by listing some or all of the field names
	and their corresponding values.
	If the field is omitted in this ind of literal, it is set to the zero value of its type. Because namses are provided,
	the order of fields does not matter.

	The two forms cannot be mixed in the same literal. Nor can you use the (order-based)first form of literal to sneak
	around the rule that unexported identifiers may not be reffered to from another package.

	package p
	type T struct{a, b int} // a,b not exported

	pakcage q
	import "p"
	var _ = p.T{a: 1, b: 2} // compilation error
	var _ = p.T{1, 2} // compilation error

	Although the last line above doesn't mentionthe unexported field identifiers, it's really using implicitly,
	so it's not allowed.


*/

package main

import (
	"fmt"
	"time"
)

type Point struct {
	X, Y int
}

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var divesh Employee

func main() {
	fmt.Println(Scale(Point{1, 2}, 5))
	/*
		Foe efficiency, larger struct types are ususally passed to or returned from dunctions indirectly using a pointer.
	*/

	fmt.Println(Bonus(&divesh, 40))

	/*
	 And this is required if the function must modify its argument, since in a call by value language like Go,
	 the called funcction recieves only a copy of an argument, not a reference to the original argument.
	*/
	AwardAnnualRaise(&divesh)
	/*
		because structs are so commonly dealt with through pointers, It's possible to use this shorthand notation
		to create and initialize a struct variable and obtain its address:


		pp := &Point{1, 2}

		It is exactly equivalent to:
		pp := new(Point)
		*pp = Point{1,2}

	*/

}

/*
	Struct values can be passed as arguments to functions and returned from them. For instance, the function scales a Point
	by a specified factor:
*/
func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}

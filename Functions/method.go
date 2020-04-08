/*
	Although there is no universally accaepted definitin of object oriented programming, for our purposes, an object
	is simply a value that has methods and a method is a function associated with a particular type. An object-oriented
	program is one that uses methods to express the properties and operations of each data structure so that clients
	need not access the objct's representaion directly.

	A method is declared with a variant of the ordinary fuction declaration in which an extra parameter appears
	before the function name. The parameter attaches the function to the type of that parameter.
*/

package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

type Path []Point

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q)) // The expression p.Distance is called selcetor, because it selects the apprpriate Distance
	// method for the reciever p of type Point.
	// there is no conflict between the two declarations of functions called Distance

	// Lets see one more example
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(perim.Distance())
	// Path is a named slice type, not a struct type like Point, yet we can still define methods for it.
	// Methods can be declared on ay named type defined in the same package.
}

// Traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)

}

// Same thing but as a method of point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

/*
	The extra parameter p is called the method's reciever, a legacy from early object-oriented languages that described
	calling a method as "sending a message to an object."

	In Go, we don't use a special name like 'this' or self for the reciever; we choose reciever's names just as we would
	for any other parameter. Since the reciever name will be frequently used, it's good idea to choose something short
	and to be consistent across methods. A common choice is the first letter of the type name, like p for Point.
*/

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

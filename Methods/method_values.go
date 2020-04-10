/*
	Usually we select and call a method in the same expression, as in p.Distance(), but it's possible to separate these
	two operations. The selector p.Distance yields a method value, a function that binds a method to a specific reciever
	value p. This function can then be invoked without a reciever value; it needs only the non-reciever arguments.

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

	distanceFromP := p.Distance // method value
	fmt.Println(distanceFromP(q))
	var origin Point //{0,0}
	fmt.Println(distanceFromP(origin))

	scalep := p.ScaleBy // method value
	scalep(2)
	scalep(3)
	fmt.Printf("Scaled p: %.2f\n", p)
	/*
		Method values are useful when a package's API calls for a function value, and the client's desired behavor
		for that function is to call a method on a specific reciever.

		Related to method value is the method expression. When calling a method, as opposed to an ordinary
		function, we must supply the reciever ina special way using the selector syntax. A method exprerssion, written
		T.f or (*T).f where T is a type, wields a function value with a regular first parameter taking the place of the
		reciever, so it can be called in thr ususal way.
	*/
	distance := Point.Distance // method expression
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)
	/*
		Method expression can be hwlpful when you need a value to represent a choice among several methods belonging
		to the same type so that you can call the chosen method with many different recievers.
	*/

}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) { //The name of this method is (*Point).ScaleBy
	p.X *= factor
	p.Y *= factor

}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func (path Path) TransalteBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}

}

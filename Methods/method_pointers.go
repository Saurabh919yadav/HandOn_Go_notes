/*
	Because callin a function makes a copy of each argument value, if a function needs to update a variable, or if
	an argument is so large that we wish to avoid copying it, we must pass the adress of the variable using a pointer.
	The same goes for methods that need to update the reciever variable: we attach them to the pointer type, such as
	\*Point.
*/

package main

import "fmt"

type Point struct{ X, Y float64 }

func main() {
	p := &Point{1, 1}
	p.ScaleBy(5)
	fmt.Println(*p)
}

func (p *Point) ScaleBy(factor float64) { //The name of this method is (*Point).ScaleBy
	p.X *= factor
	p.Y *= factor

}

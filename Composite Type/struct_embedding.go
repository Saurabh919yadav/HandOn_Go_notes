/*
	Here , we'll see how Go's unusual struct embedding mechanism lets u suse one named struct type as an anonymous field
	of another struct type, providing a convenient syntactic shortcut so that a simple dot expression like x.
	f can stand for a chain of fields x.d.e.f

	Consider a 2-D drawing program that provides a library of shapes, such as rectangles, ellipses, stars, and wheels.
	Here are two of the types it might define:
*/

package main

/*
type Circle struct {
	X, Y, Radius int
}

type Wheel struct {
	X, Y, Radius, Spokes int
}
*/
/*
	A Circle has fields for the X and Y coordinates of its senter, and a Radius. A Wheel has all the features of a
	Circle, plus Spokes, the number of inscribed radial spokes. Let's create a wheel.
*/

func main() {
	/*var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
	*/
	/*
		As the set of shapes grows, we're bound to notice similarities and repetition among them, so it may be convenient
		to factor out their common parts:
	*/

	type Point struct {
		X, Y int
	}
	/*
		type Circle struct {
			Center Point
			Radius int
		}

		type Wheel struct {
			Circle Circle
			Spokes int
		}
	*/
	// The application may be clearer for it, but this change makes accessing the fields of wheel more verbose:
	/*
		var w Wheel
		w.Circle.Center.X = 8
		w.Circle.Center.Y = 8
		w.Circle.Radius = 5
		w.Spokes = 20
	*/

	// Go lets us declare a field with a type but no name; such fields are called anonymous fields.
	// The type of the field must be a named type or a pointer to a named type. Below, Circle and Wheel have one anonymous
	// field each. we say that a point is embedded within Circle, a Circle is embedded within Wheel.
	type Circle struct {
		Point
		Radius int
	}

	type Wheel struct {
		Circle
		Spokes int
	}
	/*
		thanks to embedding, we can refer to the names at the leaves of the implicit tree without giving the intervining names:

	*/
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20

	/*
		Unfortunetly, there's no corresponding shorthand for the struct literal syntax, so neither of these will compile:
		w = Wheel{8, 8 ,5, 20}
		w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20}

		The struct literal must follow the shape of the type declaration.
	*/
}

/*
	An array is afixed-length sequence of zero or more elements of a particular type. Because of their fixed length,
	arrays are rarely used directly in Go. Slices, which ca grow and shrink, are much more versatile, but to understand
	slices we must understand arrays first.

	Individual array elements are accessed with the conventional subscript noation, where subscript run from zero to
	one less than array length. the built in function len returns the number of elements in the array.
*/

package main

import "fmt"

func main() {
	var a [3]int             // array o f3 integers
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // Print the last element

	// print the indices and the elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// print the elements only
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	/*
		By default, the elements of an array variable are initially set to the zero value for the element type,
		which is 0 for numbers. We can use an array literal to initialize an array with a list of values.
	*/

	//var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(r[2])

	/*
		In an array literal, if an elipsis "..." appears inplace of the length, the array length is determined
		by the number of initializers. The definition of q can be simplified to:
	*/
	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) //"[3]int"

	/*
		The sixe of an array is part of its type, so [3]int and [4]int are different types. The size must be a
		constant expression, that is, an expression whose value can be computed as the program is being compiled.
	*/
	//t := [3]int{1, 2, 3}
	//t = [4]int{1, 2, 3, 4} // cannot use [4]int literal (type [4]int) as type [3]int in assignment

	/*
		As we'll see, the literal syntax is similar for arrays, slices, maps, and structs. The specific form above
		is a list of values in order, but it is also possible to specify a list of index and value pair.
	*/
	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "euro", GBP: "pound", RMB: "yen"}
	fmt.Println(USD, symbol[USD]) //"0 $"

	/*
		In this form, indices can appear in any order and some may be omitted; as before, unspecified values take
		on the zero value for the element type.
	*/
	x := [...]int{99: -1} //  defines an array with 100 elements, all zero except for the last, which has value -1
	fmt.Println(x[99])

	/*
		If an array's element type is comparable then the array is comparable too, so we may directly compare
		two arrays of that type using the == operator, which reports whether all corresponding elements are equal.
		The != operator is its negation.
	*/

	n := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(n == b, n == c, b == c)
	//d := [3]int{1, 2}
	//fmt.Println(n == d) // invalid operation: n == d (mismatched types [2]int and [3]int)
}

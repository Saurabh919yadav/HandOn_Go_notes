/*
	Slices represent variable length sequences whose elements all have the same type. A slice type is written
	[]T, where the elements have type T; it looks like an array without a size.

	Arrays and slices are intimately connected. A slice is a lightweight data structure that gives access to a
	subsequence of the elements of an array, Which is know as the slice's underlying array. A slice has three
	components: pointer, a length and a capacity. The pointer points to the first element of the array that is
	reachable through the slice, which is not necessarily the array's firt element. The length is the number
	of slice element; it can't exceed the capacity, which is usually the number of elements between the
	start of the slice and the end of the underlying array. The built-in functions len and cap return
	those values.

	Mutiple slices can share the same underlying array and may refer to overlapping parts of that array.

*/

package main

import "fmt"

func main() {
	months := [...]string{1: "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(months)
	fmt.Println(Q2)
	fmt.Println(summer)

	for _, s := range months {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	/*
		Slicing beyond cap(s) causes a panic, but slicing beyond len(s) extends the slice,
		so the result may be longer than the original
	*/
	//fmt.Println(summer[:20])
	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)

	/*
		Since a slice holds a pointer to an element of an array, passing a slice to a function permits the function
		to modify the underlying array elements. In other words, copying a slice creates an alias for the
		underlying array.
		the function reverse reverses the elements of an []int slice in place, and it may be applied to slices of
		any length.

	*/

	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

	/*
		A simple way to rotate a slice ;eft by n elements is to apply the reverse function three times, first to the leading
		n elements, then to the remaining elements, and finally to the whole slice.
	*/
	reverse(a[:]) // returning to the original order [0 1 2 3 4 5]

	//rotate a left by two positions
	reverse(a[:2])
	reverse(a[2:])
	reverse(a[:])
	fmt.Println(a)

	/*
		Unlike the arrays, slices are not comparable , so we cannot use ++ to test whether two slices contain the same elements.
		The standard library provides the highly optimized bytes.Equal functions for comparing two slices of bytes([]byte)
		but for othertypes of slice, we must do the comparison ourselves:
		x = equal(a, b) -- example function defined below
	*/

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(x, y string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

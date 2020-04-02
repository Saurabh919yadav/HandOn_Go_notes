/*
	the built-in function make creates a slice of a specified element type, length, and capacity. The capacity
	argument may be omitted, in which case the capacity equals the length.

	make([]T, len)
	make([]T, cap)[:len]

	Under the hood, makes creates an unnamed array variable and returns a slice of it; the array is accessible
	only through the returned slice. In the first form, the slice is a view of the entire array, In the second,
	the slice is a view of only the array's first len elements, but its capacity include the entire array.
	The additional elements are set aside for future growth.


	The append sunction
		The built-in append function appends items to the slices:


*/

package main

import "fmt"

func main() {
	var runes []rune
	for _, r := range "Hello, world" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	/*
		The loop uses append to build the slice of nine runes encoded by the string literal, although this specific
		problem is more conveniently solved by using the built-in conversion []rune("Hello, world")

		The append function is crucial to understanding how slices work.
		Now,we will see a version of append, appendInt that isi specialized for []int slices.
	*/

	var p, l []int
	for i := 0; i < 10; i++ {
		l = appendInt(p, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(l), l)
		p = l
	}

	/*
		let's see more examples of functions like rotate and reverse, modify the elements of a slice in place.
		Given a list of strings, the nonempty function returns the non-empty ones.
	*/

	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		//there is room to grow, extend the slice
		z = x[:zlen]
	} else {
		// there is insufficient space. Allocate a new array, grow by doubling.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

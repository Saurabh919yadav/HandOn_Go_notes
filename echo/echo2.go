// echo2 prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	/*
	 In each iteration of the lop. range produces a pair of values: the index and the value of element at that index.
	 Here we don't need index values, so a basic idea is to use a temporary variable lie temp and ignore its value,
	 but Go does not permit unused local variales, so this would result in a ompilation error.

	 The solution is to use blank identifier, whose name is _. the blank identifier may be used whenever syntax requires
	 a variable name but program logic does not.
	*/
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

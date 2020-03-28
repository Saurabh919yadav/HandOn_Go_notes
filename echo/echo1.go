/*
The os package provides funtions and other values for dealing with the operating system
in a platform independent fashion. Command line arguments are available to a program in
a variable named Args that is the part of the os package; thus its name anywhere outside
the os package is os.Args.

The variable os.Args is a slice of strings. Slices are a fundamental notion in Go, for now,
think of a slice as a dynamically sized sequence s of array elements where individual
elements can be accessed as s[i] and contigous seqeuence as s[m:n]. The number of elements
is given by len(s). As in most other programming languages, all indexing in Go uses half-open
intervals that include the first index but exclude the last, because it simplifies logic.

The first element of os.Args, os.Args[0], is the name of the command itself; the other
elements are the arguments that were presented to the program when it started execution.
*/

// Echo1 prints its command line argumments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// the for loop is the only loop in Go
	/*
		different varients of for loop in Go
		for initialization; condition; post{

		}

		for condition{

		}

		and

		for {

		}
	*/
	for i := 1; i < len(os.Args); i++ { // for unary ++ and -- only postfix are legal, prefix like ++i are not legal
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(len(os.Args))
	fmt.Println(s)
}

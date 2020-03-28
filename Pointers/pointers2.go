package main

/*
	Each time we take the address of a variable or copy a pointer, we create new aliases or ways
	to identify the same variable.For example, *p is an alias for v. Pointer aliasing is useful because it allows
	us to access a variable without using its name, but this is a double-edged sword: to find al the statements that
	access a variable, we have to know all its aliases. It's not just pointers that create aliases; aliasing also
	occurs when we copy values of other reference types like slices, maps, and channels, andeven structs, arrays, and
	interfaces that contain these types.

	pointers are key to the flag package, which uses a program's command-line arguents to set the value of certain
	variables distributed throughout the program.
*/

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

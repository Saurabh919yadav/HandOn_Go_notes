package main

import "fmt"

const boilingF = 212.0

/*
	Te constant boilingF is a package level declaration, whereas the variable f and c are local to the function main.
	The name of each package level entry is visible not only throughout the source file that contains its declaration
	but throughout all the files of the package. By contrast, local declarations are visible only within the function
	in which they are declared and perhaps only within a small part of it.
*/
func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}

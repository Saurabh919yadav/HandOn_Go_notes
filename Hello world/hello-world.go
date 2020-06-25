package main

/*
	Package main is special. It defines a standalone executable program, not a library.
	Within package main the function main is also special - it's where execution of the
	program begins. Whatever main does is what program does. Of course, main wil normally
	call upon functions in other packages to do much of the work, such as fmt.Println.
*/
import "fmt" //package formatted ouptput and scanning input
/*
	you must import exactly the packages you need. A program will not compile if there
	are missing imports or if there are unnecessary ones. This strict requirement prevents
	references to used packages from accumulating as programs eveolve.
*/

func main() {
	fmt.Println("Hello World")
}

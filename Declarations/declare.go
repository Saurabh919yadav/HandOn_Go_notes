package main

/*
	A var declaretion creates a variable of particular type, attaches a name to it, and sets its initial value.
	Each declaration has the general form:

	var name type = expression

	Either the type or the = expression part may be omitted, but not both. If the type is omitted, it is determined by the
	initializer expression. If the expression is omitted, the initial value is the zero value for thr type, which
	is 0 for numbers, false for booleans, "" for strings, and nil for interfaces and reference types(slice, pointer, map
	channel, function). The Zero value of all of its elements or fields.
	The zero value mechanism ensures that a variable holds a well defined value of its type; in Go there is no such thing as uninitialized
	variable.
*/
import "fmt"

/*
	Package-level variables are intialized before main begins, and local variables are initialized as their declarations are
	encountered during function execution.

	A set of variables can also be initialized by calling a function that returns multiple values.

	Within a function, an alternate form called a short variable declarartion may be used to declare and initialize local variables.
	It takes the form

	name := expression

	and the type of name is determined by the type of expression.
	keep in mind that := is a declaration whereas = is an assignment

	One subtle important point: a short variable declaration does not necessarily declare all the variables
	on its left-hand side. If some of them were already declared in same lexical block, then the short variable
	declaration acts like an assignment to those variables.
*/
func main() {
	var s string
	fmt.Println(s)
}

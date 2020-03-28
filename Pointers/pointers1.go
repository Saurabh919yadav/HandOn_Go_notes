package main

/*
	A avriable is a piece of storage containinga value. Variables created by declarations are identified by a name
	such as x but many variables are identified only by expressions like x[i] or x.f.
	All these expessions read the value of a variable, except when they appear on the left hand side of an assignment
	in which case a new value is assigned to the variable.

	A pointer value is the address of a variable. A pointer is thus the location at which a value is stored. Not every
	value has an address, but  every variable does. With a pointer, we can read or update the value of a variable indirectly,
	without using or even knowing the name of the variale, if indeed it has a name.

	if a variable is declared var x int, the expression &x("address of x") yields a pointer to an integer variable
	that is a value of type *int, which is pronounced "pointer to int".
	The expression *p yields the value of that variable, an int, but since *p denotes a variable, it may also appear on the
	left-hand side of an assignment, in which case the assignment pdates the variable.
*/

import "fmt"

func main() {
	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to x = 2
	fmt.Println(x)  // "2"

}

/*
	Each component of a variable of aggregate type - a field of a struct or an element of an array - is also
	a variable and thus address too.
	variables are sometimes described as addressable values. Expressions that denote variable are the only
	expressions to which the address-of operator '&' may be used.

	The zero value of a pointer of any type is 'nil'. Pointers ar ecomparable; two pointers are equal if and only if
	they point to the same variable or both are nil.
*/

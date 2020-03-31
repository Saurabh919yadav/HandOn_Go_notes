/*
	A value of type bool or boolean has only two possible values true or false. The conditions in if and for
	statements are booleans, and comparison operators like ++ and < produce a boolean result.

	Boolean values can be combined with the &&(AND) and ||(OR) operators, which have short-circuit behavior:
	if the answer is already determined by the value of the left operand, the right operand is not evaluated,
	making it safe to write expressions like : s != "" && s[0] == 'x'


	There is no implicit conversion from a boolean value to a numeric value like 0 or 1, or vice versa.
	It's necesary to use an explicit if.
*/

package main

import "fmt"

func main(){
	
}

func btoi(b bool) int{
	if b{
		return 1
	}
	retutn 0
}
/*
	A string is an immutable sequence of bytes. Strings may conatin arbitrary data, including bytes with value 0,
	but usually they contain human-readable text. Text strings are conventioanlly interpreted a UTF-8-encoded
	sequences of Unicode code points(runes).

	The built-in len function returns the number of bytes(not runes) in a string, and the index operation s[i] retrieves
	the ith byte of string s, where 0 <= i < len(s).
*/

package main

import "fmt"

func main() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	/*
		The ith byte of string does not necessarily means ith character of a string, because the UTF-8 encoding of
		a non ASCII code point requires two or more bytes.

		The substring operation s[i:j] yields a new string consisiting of the bytes of the original string starting
		at index i and continuing up to , but bit including, the byte at index j. the result contains j-i bytes.

	*/
	fmt.Println(s[0:5])

	// The + operator makes a new string by concatenating two strings.
	fmt.Println("goodbye" + s[5:])

	/*
		Strings may be compared with omaprison operators like == and <; the comparison is done byte by byte,
		so the result is the natural lexicographic ordering.


		String values are immutable: the byte sequence contained in a string value can never be changed,
		though of course we can assign a new value to a string variable.
	*/

	x := "left foot"
	t := x
	x += ", right foot"

	fmt.Println(x)
	fmt.Println(t)
}

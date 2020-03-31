/*
	Four standard packages are particularly important for manipulating strings: bytes, strings, strconv and unicode.
	The strings package provides many functions for searching, replacing, comapring, trimming, splitting, and joining strings.

	The bytes package has similar functions for manipulating slices of bytes, of type []byte, which shre some properties
	with strings. Because strings are immutable, building up strings incrementally can involve a lot of allocation and copying.
	In such cases, it's more efficient to use the bytes.Buffer type.

	The strconv package prvides functions for converting boolean, integer, and floating-point values to and from their string
	representations, and functions for quoting and unquoting string.

	The unicode package provides functions like IsDigit, Letter, isUpper, and IsLower for classifying runes. Each
	function takes a single rune argument and returns a boolean. Conversion funtions like ToUpper and ToLower convert
	a rune into the given case if it is a letter. All these functions use the Unicode standard categories for letters, digits
	and so on. The strings package has similar functions, also called ToUpper and ToLower, that returns a new string
	with the specified transformation aplied to each character of the original string.

	The basename function below was inspired by the Unix shell utility of the same name. In our version, basename(s) removes
	any prefix of s that looks like a file system path with components separated by slashes, and it removes any suffix that
	looks like a file type.
*/

package main

import "fmt"

func main() {
	fmt.Println(basename("a/b/c.go")) //"c"
	fmt.Println(basename("c.d.go"))   //"c.d"
	fmt.Println(basename("abc"))      //"abc"
}

func basename(s string) string {
	//Discard last '/' and everything before
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// preserve verything before last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

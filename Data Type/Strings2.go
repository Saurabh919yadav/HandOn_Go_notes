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

*/

package main

import (
	"fmt"
	"strings"
)

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

/*
	A simpler version uses the strings.LastIndex library function.
*/

func basename_new(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

/*
	The path and path/filepath packages provide a more general set of functions for manipulating heirarchical names.
	The path package works with slash-delimited paths on any platforms. It shouldn't be used for filenames,
	but it is appropriate for other domains, like the path component of a URL. By contrast, path/filepath manipulates
	file names using the rules for the host platform, such as /foo/bar for POSIX or c:\foo\bar on Microsoft Windows.

	Let's continue with another substring example. The task is to take a string representation of an integer, such as
	"12345", and insert commas very three places, as in "12,345".
*/

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

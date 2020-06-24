/*
Dup1 prints the text of each line that appears more than once in the standard input,
preceded by its count.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	/*
		A map holds a setof key/ value pairs and provides constant time
		operations to store, retrieve, or test for an item in the set.
		The key may be of any type whose values can be compared with ==,
		strings being the most common example; the value may be of any type at all.
	*/
	input := bufio.NewScanner(os.Stdin)
	/*
		bufio package, helps make input and output efficient and convenient. One of its
		most useful feature is a type called Scanner that reads input and breaks it into
		lined or words; it's often the easiest way to process input that comes naturally
		in lines.
	*/
	for input.Scan() {
		counts[input.Text()]++

	}
	for line, n := range counts {
		if n > 1 {
			/*
				as with for, parentheses are never used around the condition in a if statement,
				but braces are required for the body. there can be a optional else part that
				is executed if the condition is false.
			*/
			fmt.Printf("%d\t%s\n", n, line)
			/*
				The function fmt.Printf, like printf in C and other languages, produces
				formatted output from a list of expressions. Its first arguent is a format
				string that specifies for subsequent arguments should be formatted. The format
				of each argument i sdetermined by a conversion charcater, a letter following
				a percent sign.
				Printf has over dozen conversions, which Go programmers call verbs.

			*/
		}
	}
}

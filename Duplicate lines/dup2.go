/*
Dup2 prints the count and text of lines that appear more than once
in the input. It reads from stdin or from a list of named files.
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
		A map is a reference to the data structure created by mke. When a map is passed to a function,
		the function recieves a copy of the reference, so any changesthe called fuction makes to
		the underlying dat structure will be visible through the caller's map reference too. In our example,
		the values inserted into the counts map by countLines are seen by main.
	*/
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			/*
				The function os.Open returns two values. The first is an open file(*os.File)
				that is used in subsequent reads by the Scanner.
				The second result of os.Open is a vlue of the built-in error type.
				if err equals the special built-in value nil, the file was opened successfully.
				The ile is read, and when the end of the input is reached, Close closes the
				file and release any resources. Onthe other hand, if err is ot nil,
				something went wrong. In that case, the error value describes the problem.
				Our simple minded error handling prints a message on the standard error stream
				using Fprintf and the verb %v, which displays a value of any type in a default format,
				and dup then carries on with the next file; the continue statement goes to te next
				iteration of the enclosing for loop.
			*/
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n")
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

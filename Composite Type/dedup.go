/*
	Go does not provide a set type, but since the keys of a map are distinct, a map can serve this purpose. To illustrate
	this program reads asequence of lines and prints only the first occurences of each distinct line. Program uses a map
	whose keys represent the set of lines that have already appeared to ensure that subsequent occurences are not printed.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

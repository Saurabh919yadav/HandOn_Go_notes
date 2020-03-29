// Echo4 prints command-line arguments using flag package

package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "Omit trailing new line")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

}

/*

	Some test cases for the program
	$ go build echo4.go
	$ ./echo4 a bc def
	$ ./echo4 -s / a bc def
	$ ./echo4 -n a bc def
	$ ./echo42 -help

*/

/*
	A variadic funvtion is one that can be called with varying numbers of arguments. The most familiar examples are
	fmt.Printf and its variants. Printf requires one fixed argument at the beginning, then accepts ant number of
	subsequent arguments.package Functions

	To delete a variadic function, the type of the final parameter is preceded by an ellipsis, "...", which
	indicates that the function may be called with any number of arguments of this type.
*/
package main

import "fmt"

func main() {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4))
	/*
		Implicitely, the caller allocates an array, copies the arguments into it,
		and passes a slice of the entire array to the function. The last call above thus behaves as the call below,
		which shows how to invoke a variadic function whe the arguments are already in a slice: place an ellipsis
		after the final argument.
	*/
	values := []int{1, 2, 3, 4}
	fmt.Printf("Passing a list to the variadic function : %d\n", sum(values...))
	/*
		Although the ...int parameter behaves like a slice within the function body, the type of variadic function
		is distinct from the type of a funtion with an ordinary slice parameter.
	*/
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)

}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val

	}
	return total
}

func f(...int) {}
func g([]int)  {}

/*
	A const declaration may use the constant generator iota, which is used to create a sequence of related values
	without spelling out each one explicitly. In a const declaration, the value of iota begins at zero and
	increments by one for each item in the sequence.
*/

package main

import "fmt"

func main() {
	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thrusday
		Friday
		Saturday
	)

	/*
		This dec;ares Sunday to be 0, Monday to be 1 and so on.
	*/

	fmt.Println(Sunday, Monday, Tuesday)

	/*
		We can use iota in more complex expressions too, as in this example from the net package where
		each of the lowest 5 bits of an unsigned integer is given a distinct name and boolean interpretation.
	*/

	type Flags uint
	const (
		FlagUp Flags = 1 << iota //is up
		FlagBroadcast
		FlagLoopback
		FlagPointToPoint
		FlagMulticast
	)

	/*
		As iota increments, each constant is assigned the value of 1<<iota, which evaluates to successive powers of two
		each corresponding to a single bit. We can use these constants within functions that test, set, or clear
		one or more of these bits.
	*/
}

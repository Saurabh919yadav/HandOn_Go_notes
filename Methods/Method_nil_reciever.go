/*
	Just as some functions allow nil pointers as arguments, so do some methods for theie receiver, espacially if nil
	is a meaningful zero value of the type, as with maps and slice. In this simple linked list of integers, nil
	represents the empty list:
*/

package main

// An Inlist is a linked list of integers.
// A nil *intList represents the empty list.

type IntList struct {
	value int
	tail  *IntList
}

// Sum returns the sum of the lists elements.
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.value + list.tail.Sum()
}

/*
	When you define a type whose methods allow nil as  areciever value, it's worth pointing this out explicitely
	in its deocumentation comment, as we did above.
*/

/*
	A named struct type S can't declare a fields\ of the same type S: an aggregate value cannot contain itself.
	But S may declare a field of the pointer type *s, which lets us create recusrsive data structures like linked
	lists and trees. this is illustrated in the code below, which uses a binary tree to implement an insertion sort.
*/

package main

type tree struct {
	value       int
	left, right *tree
}

func main() {

}

//Sort values in place
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elemnts of t to the values and return the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = appendValues(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// equivalent to return &tree{value: value}
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

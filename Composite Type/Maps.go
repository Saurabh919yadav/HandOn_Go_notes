/*
	The hash table is one of the most ingenius and versatileof all data structures. It is an unordered colletion of
	key/value pairs in which all the keys are distinct, and the value associated with a given key can be retrieved,
	updated, or removed using a constant number of key comparisons on the average, no matter how large the hash table.

	In Go, a map is a reference to a hash table, and a map type is written map[K]V, where K and V are the types
	of its keys and values. All of the keys in the given mp are the same type, and all of the values are of the
	same type, but the keys need not be of the same type as the values. the key type K must be comparable using ==
	so that the map can test whether a given key is equal to one already within in it. Though floating-point
	numbers are comparable, it's a bad idea to compare floats for equality and, especially bad if NaN is a possible
	value. there are no restrictions on the value of type V.

	the built-in function make can be used to create a map:
*/

package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34
	/*
		we can also use a map literal to create a new map populated with some initial key/value pairs
		ages := map[string]int{
			"alice": 31,
			"charlie": 34,
		}

		So, an alternative expression for a new empty map is map[string]int{}

	*/

	// Map elements are accessed through the usual subscript notation:
	ages["alice"] = 32
	fmt.Println(ages["alice"])

	// and removed using built-in funtion delete
	delete(ages, "alice")
	/*
		All of these operaions  are safe even if the element isn't in the map; a map lookup using a key isn't present returns
		the zero value of its type
	*/
	ages["bob"] = ages["bob"] + 1

	// The shorthand assignment forms x += y and x++ also works for map elements, so we can rewrite the statements above as
	ages["bob"] += 1
	// or even more concisely as
	ages["bob"]++

	/*
		But a map element is not a variable so we can not take its address:
		_ = &ages["bob"] // compilation error : cannot take address of map elements

		One reason that we can't take the address of a map element is that growing a map might cause rehasing of existing elements
		into a new storage locations,thus potentially invalidating the address.

		for enumeration we use range based for loops
	*/
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)

	}
}

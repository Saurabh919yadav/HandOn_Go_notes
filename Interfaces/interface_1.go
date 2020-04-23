/*
	Interface types express generalizations or abstractions about the behavior of other types. By generalizing,
	interfaces let us write functions that are more flexible and adaptable because they are not tied to
	the deatils of one particular implementation.
	Many object-oriented languages have some notion of interfaces, but what makes Go's interfaces so distinctive
	is that they are satisfied implicitely. In other words, there's no need to declare all the interfaces that
	a given concrete type satisfies; simply processing the necessary methods is enough. This design lets you create
	new interfaces that are satisfied by existing concrete types without changing the existing types,
	which is particularly useful for types defined in packages that you don't control.
*/
/*
	Go's numeric data type inclde several sizes of intergers, floating-point numbers, and complex numbers.
	Each numeric type determines the size and signedness of its value. Let's begin with integers.

	Go provides both signed and unsigned integer arithmetic. There are four distinct type of signed integers- 8,16,32
	and 64 bits - represented by the types int8, int16, int32 and int64, and corresponding unsigned versions uint8, uint16,
	uint32 and uint64.

	There are also two types called int and uint that are the natural or ost efficient sixe for signed and unsigned integers
	on a particular platform; int is by far the most widely used numeric type. Both these types have thesame size, either 32
	or 64 bits, but one must not make assumptions about which; different compilers may make different choices even on
	identical hardware.package Data Type

	rune is a synonym for int32 and type byte is a synonym for uint8. Finally there is an unsigned integer type uintptr, whose
	width is not specified but is sufficient to hold all the bits of the pointer value. The uintptr is used only for low
	level programming, such as the boundary of Go program with a C library or an operating system.

	If the result of an arithmetic operation, whether signed or unsigned, has more bits than can be represented in the result
	type, its is called overflow. The high order bits that do not fit are silently discarded. If the original number is a signed type
	the result could be negative if the leftmost bit is 1.
*/

package main

import "fmt"

func main() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	var i int8 = 127
	fmt.Println(i, i+1, i*i)
	/*
		Two types of the same type may be compared using the binary comparison operators; the type of a comparsion expression
		is a boolean.

		there are also unary addition and subtraction operators :
		+ unary positive (no effect)
		- unary negation
		For integers, +x is a shorthand for 0+x and -x is shorthand for 0-x; for floating pint and complex, +x is just x and
		-x is the negation of x.


		Go also provides bitwise binary operators.
		The ^ operator is bitwise OR (XOR) when used as binary operator, but when used as a unary prefix opeartor
		it is bitwise negation or complement; that is, it returns a value with each bit in its operand inverted. The &^ operator
		is bit clear (AND NOT): in the expression z = x &^ y, each bit of z is 0 if the corresponding bit of y is 1; otherwise
		it equals the corresponding bit of x.

	*/

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", x&y)
	fmt.Printf("%08b\n", x|y)
	fmt.Printf("%08b\n", x^y)
	fmt.Printf("%08b\n", x&^y)

	for i := uint(0); i < 8; i++ {
		if x&(1<<1) != 0 { // membership test
			fmt.Println(i) // "1", "5"

		}
	}

	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)

	/*
		In general, explicit conversion id required to convert a value from one type to another, and binary operators for arithmetic
		and logic(except shifts) must have operands of the same type. Although it ocassionaly result in longer expressions,
		it also eliminates a whole class of problems and makes programs easier to understand.
	*/

	var apples int32 = 1
	var oranges int16 = 2
	//var compote int = apples + oranges //invalid operation: apples + oranges (mismatched types int32 and int16)
	var compote int = int(apples) + int(oranges)
	fmt.Println(compote)

	/*
		Float to integer coversion discards any fractional part, truncating towards 0. You should avoid conversions in which
		the operand is out of range of target type, because the behavior depends on the implementation.
	*/

}

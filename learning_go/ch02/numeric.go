package main

import "fmt"

func main(){
	// Type name 	Value range
	// int8 		–128 to 127
	// int16 		–32768 to 32767
	// int32 		–2147483648 to 2147483647
	// int64 		–9223372036854775808 to 9223372036854775807
	// uint8 		0 to 255
	// uint16 		0 to 65536
	// uint32 		0 to 4294967295
	// uint64 		0 to 18446744073709551615

	// Go does have some special names for integer types.
	// A byte is an alias for uint8
	// The second special name is int. On a 32-bit CPU, int is a 32-bit signed integer like an int32. 
	// On most 64-bit CPUs, int is a 64-bit signed integer, just like an int64.


	// • If you are working with a binary file format or network protocol that has an inte‐
	// ger of a specific size or sign, use the corresponding integer type.

	// • If you are writing a library function that should work with any integer type, write
	// a pair of functions, one with int64 for the parameters and variables and the
	// other with uint64.

	// • In all other cases, just use int
	var b, c int64 = 1, 2
    fmt.Println(b, c)


	// Strings in Go are immutable; you can reassign the value of a string variable, but you
	// cannot change the value of the string that is assigned to it.

	var a = "initial"
	a="not initial"
    fmt.Println(a)

	// Go doesn’t allow automatic type promotion between variables. 
	// You must use a type conversion when variable types do not match

	var x int = 10
	var y float64 = 30.2
	var z float64 = float64(x) + y
	var d int = x + int(y)
    fmt.Println(z)
    fmt.Println(d)


	// var Versus :=

	// 	The most verbose way to declare a variable in Go uses the var keyword, an explicit
	// type, and an assignment. It looks like this:
	var x1 int = 10
    fmt.Println(x1)
	// If the type on the righthand side of the = is the expected type of your variable, you can
	// leave off the type from the left side of the =. Since the default type of an integer literal
	// is int, the following declares x to be a variable of type int:
	var x2= 10
    fmt.Println(x2)

	// The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.
	f := "apple"
    fmt.Println(f)

	// Constants can be typed or untyped. An untyped constant works exactly like a literal;
	// it has no type of its own, but does have a default type that is used when no other type
	// can be inferred. A typed constant can only be directly assigned to a variable of that
	// type.
	const x3 = 10
	const typedX int = 10
}
package main

//Ftoc prints two Fahrenheit-to-celcius conversions

/*
	A function declaration is name, a list of parameters, an optional list of results, and the fuction body, which
	contains the statements that define what the function does. the result list is omitted if the function does not
	return anything. Execution of function begins with the first statements and continues until it encounters a return
	statement or reaches the end of the function that has no results. Control and any results are then returned to the caller.
*/
import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

/*
	A struct is an aggregate data type that groups together zero or more named values of
	arbitrary types as a single entity. Each value is called a field. The classic example of a struct from data processing
	is employee record, whose fields are Unique ID, the employee's name, address, date of birth, position, salary,
	manager and the like. All of these fields are collected into a singlr entity that can be copied as a unit, passed to
	functions and returned by them, stored in arrays and so on.

*/
package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

/*
	he individual fields of dilbert are accessed using dot notation like dilbert is a variable, its fields are variables too,
	so we may assign to a field.
*/
func main() {
	dilbert.Salary -= 5000
	// or take its address and accesss it through a pointer:
	position := &dilbert.Position
	*position = "Senior" + *position // promoted
	// The do notation also works with a pointer to a struct:
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += "(proative team player)"
	// It is equivalent to: (*employeeOfTheMoth).Position += "(proactive team player)"

	/*
		Given an employee's unique Id, the function EmployeeByID returns a pointer to an Employee struct. We can use
		the dot notation to access its fields
	*/
	fmt.Println(EmployeeByID(dilbert.ManagerID).Position)

	id := dilbert.ID
	EmployeeByID(id).Salary = 0 //  fired for no real...reason

	/*
		The last statement updates the Employee struct that is pointed by the result of the call to the EmployeeByID.
		If the result type of EmployeeByID were changed to Employee instead of *Employee, the assignment would not compile since
		its left hand side would not identify a variable.

		Fields are usually written one per line, with the fields anme preceeding its type, but consecutive fields of the same type
		may be combined, as with Name and Address like:
		type Emplyee struct{
			...
			Name, Address  string
			...
		}


		Fields order is significant to type identity. Had we also combined the declaration of the Position fields(also a string),
		or interchanged Name and Address, we would be defining a different struct type. Typically we only combine the declarations
		of related fields.

		The name of a struct fields is exported if it begins with a capital letter; this is Go's main access control mechanism.
		A struct type may contain a mixture of exported and unexported fields.

		Struct types tend to be verbose because they often involve a line for each field. Although we could write out the whole
		type each tie it is needed, the repetition woul dget tiresome. Instead, struct types usually appear within the declaration of
		a named type like Employee.
	*/
}

func EmployeeByID(id int) *Employee { /*...*/ }

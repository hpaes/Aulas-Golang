package main

import (
	"fmt"
	"strconv"
)

func main() {

	displayProfile()
	printVariables()
	printZeroValues()
	convertToIntFloatString()
	calculateOperations()
	applyAssignmentOperators()
	compareIntegers()
	logicalOperations()

	manipulatePointers()

}

func manipulatePointers() {
	fmt.Printf("\nPointers\n")
	var x int = 10
	fmt.Printf("x = %d\n", x)
	var y = &x
	fmt.Printf("y = &x\n")
	*y = 20
	fmt.Printf("*y = 20\n")
	fmt.Printf("x = %d\n", x)
}

func logicalOperations() {
	var p bool = true
	var q bool = false
	fmt.Printf("\nLogical Operators\n")
	fmt.Printf("p: %t, q: %t\n", p, q)
	fmt.Printf("p && q: %t\n", p && q)
	fmt.Printf("p || q: %t\n", p || q)
	fmt.Printf("!p: %t\n", !p)
	fmt.Printf("!q: %t\n", !q)
}

func compareIntegers() {
	var x int = 10
	var y int = 20
	fmt.Printf("\nComparing %d and %d\n", x, y)
	fmt.Printf("10 == 20: %t\n", x == y)
	fmt.Printf("10 != 20: %t\n", x != y)
	fmt.Printf("10 < 20: %t\n", x < y)
	fmt.Printf("10 > 20: %t\n", x > y)
	fmt.Printf("10 <= 20: %t\n", x <= y)
	fmt.Printf("10 >= 20: %t\n", x >= y)
}

func applyAssignmentOperators() {
	var x int = 10
	fmt.Printf("\nAssignment Operators\n")
	fmt.Printf("Initial value of x: %d\n", x)
	x += 10
	fmt.Printf("x+=10 : %d", x)
	x -= 10
	fmt.Printf("\nx-=10 : %d", x)
	x *= 10
	fmt.Printf("\nx*=10 : %d", x)
	x /= 10
	fmt.Printf("\nx/=10 : %d", x)
	x %= 10
	fmt.Printf("\nx%%=10 : %d\n", x)
}

func calculateOperations() {
	var x int = 13
	var y int = 12

	fmt.Printf("\nSum: %d\n", x+y)
	fmt.Printf("Subtraction: %d\n", x-y)
	fmt.Printf("Multiplication: %d\n", x*y)
	fmt.Printf("Division: %d\n", x/y)
	fmt.Printf("Modulus: %d\n", x%y)
}

func convertToIntFloatString() {
	var integer int = 10
	var float float64 = float64(integer)
	fmt.Printf("\nInteger: %d\n", integer)
	fmt.Printf("Converting int to float64: %.2f\n", float)
	fmt.Printf("Converting int to string: %s\n", strconv.Itoa(integer))
}
func printZeroValues() {
	var integer int
	var float float64
	var boolean bool
	var String string

	fmt.Println("\nZero values: ")
	fmt.Printf("Integer: %d\n", integer)
	fmt.Printf("Float: %.2f\n", float)
	fmt.Printf("Boolean: %t\n", boolean)
	fmt.Printf("String: %s\n", String)
}

func printVariables() {
	var integer int = 10
	var float float64 = 10.2
	var boolean bool = true
	var String string = "Texto"

	fmt.Println("\nDeclared variables: ")
	fmt.Printf("Integer: %d\n", integer)
	fmt.Printf("Float: %.2f\n", float)
	fmt.Printf("Boolean: %t\n", boolean)
	fmt.Printf("String: %s\n", String)
}

func displayProfile() {
	name := "Herbert"
	age := 31
	city := "Recife"
	fmt.Printf("Name: %s\nAge: %d\nCity: %s\n", name, age, city)
}

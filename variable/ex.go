package variable

import "fmt"

// 3
func AssignNum() {
	var num int = 20
	fmt.Println(num)
}

// 4
func DefineAppName() {
	const AppName = "GoApp"
	fmt.Println(AppName)
}

// 5
func SumNumbers() {
	var x = 10
	var y = 20

	fmt.Println(x + y)
}

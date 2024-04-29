package condition

import "fmt"

func CheckAdult(age int) {
	if age >= 18 {
		fmt.Println("adult")
	} else {
		fmt.Println("minor")
	}
}

func CheckGrade(grade string) {
	switch grade {
	case "A":
		fmt.Println("Excellent")
	case "B":
		fmt.Println("Good")
	case "C":
		fmt.Println("Fair")
	default:
		fmt.Println("Poor")
	}
}

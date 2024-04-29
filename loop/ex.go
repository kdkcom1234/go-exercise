package loop

import "fmt"

func OneToTen() {
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func ZeroToNineSqure() {
	for i := 0; i < 10; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()
}

func Days() {
	var days = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	for _, day := range days {
		fmt.Print(day, " ")
	}
	fmt.Println()
}

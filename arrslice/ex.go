package arrslice

import "fmt"

func DefineArr() {
	var arr = [5]int{1, 2, 3, 4, 5}
	for _, num := range arr {
		fmt.Print(num, " ")
	}
	fmt.Println()
}

func DefineAndAppendSlice() {
	var s = []int{}
	for num := range 5 {
		s = append(s, num+1)
	}

	for _, num := range s {
		fmt.Print(num, " ")
	}

	fmt.Println()
}

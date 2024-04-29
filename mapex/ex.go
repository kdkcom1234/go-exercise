package mapex

import "fmt"

func DefineMapVar() {
	var values = map[string]int{"a": 1, "b": 2}
	values["c"] = 3

	for key, val := range values {
		fmt.Print("key:", key, ", val:", val, " | ")
	}
	fmt.Println()
}

func DefineScoreMap() {
	var scores = make(map[string]int)
	scores["ann"] = 10
	scores["ko"] = 15
	scores["mike"] = 5

	for name, score := range scores {
		fmt.Print(name, ":", score, " | ")
	}
	fmt.Println()
}

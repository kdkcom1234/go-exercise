package goroutine

import "fmt"

func HelloGoroutine(c chan string) {
	c <- "Hello World"
}

func SendNum(nums int, c chan int) {
	for num := range nums {
		c <- num + 1
	}
}

func MakeBufferedNum(c chan int) {
	for num := range 100 {
		c <- num + 1
	}
}

func PrintBufferedNum(c chan int) {
	fmt.Print(<-c, " ")
}

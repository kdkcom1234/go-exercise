package main

import (
	"fmt"

	"github.com/kdkcom1234/go-exercise/arrslice"
	"github.com/kdkcom1234/go-exercise/condition"
	"github.com/kdkcom1234/go-exercise/function"
	"github.com/kdkcom1234/go-exercise/goroutine"
	"github.com/kdkcom1234/go-exercise/loop"
	"github.com/kdkcom1234/go-exercise/mapex"
	"github.com/kdkcom1234/go-exercise/packageimport"
	"github.com/kdkcom1234/go-exercise/pointer"
	"github.com/kdkcom1234/go-exercise/structure"
	"github.com/kdkcom1234/go-exercise/variable"
)

func main() {
	fmt.Println("--package/import--")
	packageimport.HelloWorld()
	packageimport.SqureRoot(16)

	fmt.Println("--variable/const--")
	variable.AssignNum()
	variable.DefineAppName()
	variable.SumNumbers()

	fmt.Println("--function--")
	fmt.Println(function.Add(5, 7))
	fmt.Println(function.Multiply(3, 4))

	fmt.Println("--loop--")
	loop.OneToTen()
	loop.ZeroToNineSqure()
	loop.Days()

	fmt.Println("--condition--")
	condition.CheckAdult(18)
	condition.CheckGrade("B")

	fmt.Println("--pointer--")
	pointer.SetPointerVar()
	pointer.SetTo100()

	fmt.Println("--arrslice--")
	arrslice.DefineArr()
	arrslice.DefineAndAppendSlice()

	fmt.Println("--maps--")
	mapex.DefineMapVar()
	mapex.DefineScoreMap()

	fmt.Println("--struct--")
	person := structure.Person{Name: "Ko", Age: 40}
	fmt.Println(person.Name, person.Age)
	rect := structure.Rectangle{Width: 10, Height: 5}
	fmt.Println(rect.Area())

	fmt.Println("--goroutine--")
	// 1
	c1 := make(chan string)
	go goroutine.HelloGoroutine(c1)
	fmt.Println(<-c1)
	// 2
	c2 := make(chan int)
	go goroutine.SendNum(10, c2)
	for range 10 {
		fmt.Print(<-c2, " ")
	}
	fmt.Println()
	// 3
	c3 := make(chan int, 5)
	go goroutine.MakeBufferedNum(c3)
	for range 100 {
		goroutine.PrintBufferedNum(c3)
	}
	fmt.Println()
	// 4

}

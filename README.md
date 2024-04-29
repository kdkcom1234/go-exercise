아래는 Go 언어의 기초 문법을 다루는 30개의 코딩 연습 문제입니다. 문제들은 패키지/imports, 변수/상수, 함수, 반복문, 조건문, 포인터, 배열/슬라이스, 맵, 구조체까지 다양한 주제를 다룹니다.

### 코딩 연습 문제

#### 패키지/Imports

1. `fmt` 패키지를 import하고, "Hello, World!"를 출력하는 프로그램을 작성하세요.
2. `math` 패키지를 사용하여, `math.Sqrt(16)`의 결과를 출력하세요.

#### 변수/상수

3. 정수형 변수 `num`을 선언하고, 20을 할당한 다음 출력하세요.
4. 문자열 상수 `AppName`을 "GoApp"으로 정의하고 출력하세요.
5. 변수 `x`와 `y`를 선언하고, 각각 10과 20을 할당한 후 두 변수의 합을 출력하세요.

#### 함수

6. 두 정수를 입력받아 더한 결과를 반환하는 함수 `add`를 작성하고, 이 함수를 사용하여 5와 7의 합을 출력하세요.
7. 두 정수의 곱을 반환하는 함수 `multiply`를 작성하고, 이 함수를 사용하여 3과 4의 곱을 출력하세요.

#### 반복문

8. 1부터 10까지 출력하는 for 루프를 작성하세요.
9. 0부터 9까지의 숫자를 각각 제곱하여 출력하는 for 루프를 작성하세요.
10. 문자열 슬라이스 `days`가 ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]일 때, 각 요소를 출력하는 for-range 루프를 작성하세요.

#### 조건문

11. 정수 `age`가 18 이상이면 "adult"를, 그렇지 않으면 "minor"를 출력하는 if 문을 작성하세요.
12. 문자열 `grade`가 "A", "B", "C" 중 하나에 따라 각각 "Excellent", "Good", "Fair"를 출력하고, 그 외의 경우 "Poor"를 출력하는 switch 문을 작성하세요.

#### 포인터

13. 정수형 변수 `a`를 선언하고, 그 포인터를 사용하여 `a`의 값을 50으로 설정한 후 출력하세요.
14. 함수 안에서 정수형 변수 `num`의 값을 100으로 변경하는 함수 `setTo100`을 작성하고, 변수 값의 변경을 확인하세요.

#### 배열/슬라이스

15. 정수형 배열 `arr`를 선언하고, {1, 2, 3, 4, 5}로 초기화한 후 배열의 모든 요소를 출력하세요.
16. 정수형 슬라이스 `s`를 선언하고, 1부터 5까지의 값을 append 함수를 사용해 추가한 후 슬라이스의 모든 요소를 출력하세요.
17. 슬라이스 `numbers`에서 홀수만 필터링하여 새 슬라이스 `oddNumbers`에 저장하고, 이를 출력하세요.

#### 맵

18. `string`을 키로 하고 `int`를 값으로 하는 맵 `m`을 선언하고 초기화하여, 몇 가지 요소를 추가한 후, 모든 키-값 쌍을 출력하세요.
19. 맵 `scoreMap`에 이름을 키로, 점수를 값으로 추가하고, 특정 이름의 점수를 출력하세요.

#### 구조체

20. `Person` 구조체를 정의하고, 이름(name)과 나이(age) 필드를 포함하게 하세요. 이 구조체를 사용하여 사람 객체를 생성하고, 그 필드를 출력하세요.
21. 구조체 `Rectangle`을 정의하고, 너비(width)와 높이(height) 필드를 포함하게 하세요. 이 구조체를

사용하여 사각형의 면적을 계산하는 메소드 `Area`를 작성하세요.

#### 고루틴과 채널

문제 1: 기본적인 고루틴 생성
고루틴을 생성하여 "Hello, World!"를 출력하는 Go 프로그램을 작성하세요.

문제 2: 채널을 통한 데이터 전송
정수를 보내고 받는 간단한 고루틴과 채널을 사용하는 프로그램을 작성하세요. 메인 고루틴에서는 1부터 10까지의 수를 채널에 보내고, 생성된 고루틴에서는 채널에서 수를 받아 출력하세요.

문제 3: 버퍼 있는 채널
버퍼를 가진 채널을 생성하고, 버퍼가 꽉 찰 때까지 숫자를 채널에 보내는 프로그램을 작성하세요. 채널의 버퍼 크기는 5로 설정하세요. 고루틴에서는 채널로부터 숫자를 받아 출력하세요.

문제 4: 채널을 이용한 고루틴 동기화
두 개의 고루틴을 생성하여, 하나는 'Ping', 다른 하나는 'Pong'을 출력하도록 하세요. 이 두 고루틴이 서로 동기화되어 번갈아가면서 출력할 수 있도록 채널을 사용하세요.

문제 5: 채널을 이용한 모든 고루틴의 완료 기다리기
5개의 고루틴을 생성하여 각각 다른 시간(1초, 2초, 3초, 4초, 5초) 동안 대기한 후, 완료 메시지를 출력하도록 하세요. 모든 고루틴이 완료될 때까지 메인 고루틴이 종료되지 않도록 하세요. 각 고루틴의 완료를 알리는 신호로 채널을 사용하세요.

### 답안

1.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

2.

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.Sqrt(16))
}
```

3.

```go
package main

import "fmt"

func main() {
    var num int = 20
    fmt.Println(num)
}
```

4.

```go
package main

import "fmt"

func main() {
    const AppName string = "GoApp"
    fmt.Println(AppName)
}
```

5.

```go
package main

import "fmt"

func main() {
    var x int = 10
    var y int = 20
    fmt.Println(x + y)
}
```

6.

```go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    fmt.Println(add(5, 7))
}
```

7.

```go
package main

import "fmt"

func multiply(a int, b int) int {
    return a * b
}

func main() {
    fmt.Println(multiply(3, 4))
}
```

8.

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
}
```

9.

```go
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        fmt.Println(i * i)
    }
}
```

10.

```go
package main

import "fmt"

func main() {
    days := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
    for _, day := range days {
        fmt.Println(day)
    }
}
```

11.

```go
package main

import "fmt"

func main() {
    var age int = 20
    if age >= 18 {
        fmt.Println("adult")
    } else {
        fmt.Println("minor")
    }
}
```

12.

```go
package main

import "fmt"

func main() {
    var grade string = "B"
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
```

13.

```go
package main

import "fmt"

func main() {
    var a int
    var p *int = &a
    *p = 50
    fmt.Println(a)
}
```

14.

```go
package main

import "fmt"

func setTo100(num *int) {
    *num = 100
}

func main() {
    var myNum int = 25
    setTo100(&myNum)
    fmt.Println(myNum) // Output: 100
}
```

15.

```go
package main

import "fmt"

func main() {
    var arr [5]int = [5]int{1, 2, 3, 4, 5}
    for _, value := range arr {
        fmt.Println(value)
    }
}
```

16.

```go
package main

import "fmt"

func main() {
    s := make([]int, 0)
    for i := 1; i <= 5; i++ {
        s = append(s, i)
    }
    fmt.Println(s)
}
```

17.

```go
package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    var oddNumbers []int
    for _, num := range numbers {
        if num%2 != 0 {
            oddNumbers = append(oddNumbers, num)
        }
    }
    fmt.Println(oddNumbers)
}
```

18.

```go
package main

import "fmt"

func main() {
    m := make(map[string]int)
    m["Alice"] = 23
    m["Bob"] = 34
    for key, value := range m {
        fmt.Printf("%s: %d\n", key, value)
    }
}
```

19.

```go
package main

import "fmt"

func main() {
    scoreMap := make(map[string]int)
    scoreMap["Alice"] = 95
    scoreMap["Bob"] = 82
    fmt.Println("Alice's score:", scoreMap["Alice"])
}
```

20.

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "John", Age: 30}
    fmt.Printf("%s is %d years old.\n", p.Name, p.Age)
}
```

21.

```go
package main

import "fmt"

type Rectangle struct {
    Width  int
    Height int
}

func (r Rectangle) Area() int {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    fmt.Println("Area of rectangle:", rect.Area())
}
```

이 문제들을 통해 Go 언어의 기초를 연습할 수 있으며, 각각의 답안 코드는 해당 문제에 대해 올바르게 구현된 예시를 제공합니다.

네, Go 언어의 goroutine과 channel을 사용하는 연습문제를 만들어 드리겠습니다. 각 문제는 goroutine과 channel의 기본적인 사용 방법을 이해하고 실습할 수 있도록 구성하겠습니다. 답안은 각 문제 뒤에 제공하겠습니다.

### 연습문제 1: 기본적인 Goroutine 사용

Go 프로그램을 작성하여, 100까지 숫자를 출력하는 데에 5개의 goroutine을 사용하세요. 각 goroutine은 20개의 숫자를 차례대로 출력해야 합니다.

#### 답안

```go
package main

import (
	"fmt"
	"sync"
)

func printNumbers(start int, wg *sync.WaitGroup) {
	defer wg.Done()
	end := start + 20
	for i := start; i < end; i++ {
		fmt.Println(i)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go printNumbers(i*20, &wg)
	}
	wg.Wait()
}
```

### 연습문제 2: Channel 기본

정수들의 슬라이스를 받아 모든 정수의 합을 반환하는 함수를 goroutine과 channel을 이용하여 구현하세요.

#### 답안

```go
package main

import (
	"fmt"
)

func sum(slice []int, c chan int) {
	total := 0
	for _, v := range slice {
		total += v
	}
	c <- total
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan int)
	go sum(slice, c)
	total := <-c
	fmt.Println("Total:", total)
}
```

### 연습문제 3: 멀티 채널 사용

3개의 goroutine을 생성하여 각각 1에서 100, 101에서 200, 201에서 300까지의 숫자를 각각 더하는 프로그램을 작성하세요. 각 결과를 메인 함수에서 받아 최종 합을 출력하세요.

#### 답안

```go
package main

import (
	"fmt"
)

func sum(start, end int, c chan int) {
	total := 0
	for i := start; i <= end; i++ {
		total += i
	}
	c <- total
}

func main() {
	c := make(chan int)
	go sum(1, 100, c)
	go sum(101, 200, c)
	go sum(201, 300, c)

	result1, result2, result3 := <-c, <-c, <-c
	fmt.Println("Final Total:", result1 + result2 + result3)
}
```

### 연습문제 4: Select 문 사용

2개의 channel을 생성하고 각각 다른 goroutine에서 데이터를 보냅니다. 메인 goroutine에서는 select 문을 사용하여 어느 channel로부터 데이터가 먼저 도착하는지 확인하고 출력하세요.

#### 답안

```go
package main

import (
	"fmt"
	"time"
)

func send(c chan int, delay time.Duration, value int) {
	time.Sleep(delay)
	c <- value
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go send(c1, 2*time.Second, 1)
	go send(c2, 1*time.Second, 2)

	select {
	case res := <-c1:
		fmt.Println("Received from c1:", res)
	case res := <-c2:
		fmt.Println("Received from c2:", res)
	}
}
```

### 연습문제 5: Buffer된 Channel

100개의 정수를 buffer된 channel을 통해 보내고 받는 프로그램을 작성하세요. 받은 데이터를 출력하세요.

#### 답안

```go
package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 100)

	for i := 1; i <= 100; i++ {
		c <- i
	}
	close(c)

	for i := range c {
		fmt.Println(i)
	}
}
```

이 연습문제들은 Go 언어의 동시성 패턴을 이해하고 실

습하는 데 도움이 될 것입니다.

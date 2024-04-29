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

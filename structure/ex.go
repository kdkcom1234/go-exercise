package structure

type Person struct {
	Name string
	Age  int
}

type Rectangle struct {
	Width  int
	Height int
}

func (rect *Rectangle) Area() int {
	return rect.Width * rect.Height
}

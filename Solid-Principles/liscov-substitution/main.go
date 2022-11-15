package main

import "fmt"

type Shape interface {
	GetHeight() int
	GetWidth() int
	SetHeight(height int)
	SetWidth(width int)
}

type Rectangle struct {
	height, width int
}

func (r *Rectangle) GetHeight() int {
	return r.height
}
func (r *Rectangle) GetWidth() int {
	return r.width
}
func (r *Rectangle) SetHeight(height int) {
	r.height = height
}
func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func Useit(shape Shape) {

	width := shape.GetWidth()
	shape.SetHeight(10)
	fmt.Println("Expected area is : ", 10*width)
	fmt.Println("Area is : ", shape.GetHeight()*shape.GetWidth())
}

func main() {

	rect := Rectangle{5, 5}
	Useit(&rect)
}

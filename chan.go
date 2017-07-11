package main

import (
	"fmt"
	"math"
	"sort"
	"reflect"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

type Circle struct {
	Shape
}

type Shape struct {
	Radius float32
}

func (s Shape) Area() float32{
	return s.Radius*s.Radius*math.Pi
}

func main() {
	a := []int{4,3,67,89}
	sort.Ints(a)
	fmt.Println(a)

	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Circle{Shape{5}}      // Area() of Square needs a pointer
	// shapes := []Shaper{Shaper(r), Shaper(q)}
	// or shorter
	shapes := []Shaper{r, q}
	fmt.Println(shapes[0]==r)
	if _,ok:=shapes[0].(Shaper);ok{
		fmt.Println("yes")
	}
	fmt.Println("haha",reflect.ValueOf(r).Kind())
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}
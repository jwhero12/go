package main

import "fmt"
import "math"

type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct{
	radius float64
}

func (r Rect) area() float64{
	return r.width * r.height
}

func (r Rect) perimeter() float64{
	return 2*r.width + 2*r.height
}

func (c Circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64{
	return 2 * math.Pi * c.radius
}

func measure(s Shape){
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func showArea(shapes ...Shape){
	for _,shape:= range shapes{
		fmt.Println(shape.area())
	}
}

func printIf(itf interface{}){
	fmt.Println(itf)
}

func main(){
	var x interface{}
	r:=Rect{width:5, height:2}
	c:=Circle{radius:10}

	x = 55
	x = "test"

	printIf(x)
	showArea(c,r)

	measure(r)
	measure(c)
}


package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	height, width float64
}

type Cricle struct {
	redius float64
}

func (r Rectangle) Area() float64 {
	return 2 * r.height * r.height
}

func (c Cricle) Area() float64 {
	return 2 * 3.1415 * c.redius
}

func (c Cricle) Area2() float64 {
	return 2 * 3.1415 * c.redius
}

func CalcuteArea(s Shape) float64 {
	return s.Area()
}

func main() {
	rec := Rectangle{
		height: 10.0,
		width:  2.0,
	}
	cir := Cricle{
		redius: 5,
	}
	fmt.Println(cir.Area2())
	fmt.Println("okkkkkkkkkk")
	fmt.Println(CalcuteArea(rec))
	fmt.Println(CalcuteArea(cir))

}

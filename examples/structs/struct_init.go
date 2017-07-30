package main

import "fmt"

type Rect struct {
	x, y          float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

func main() {
	//init
	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{width: 100, height: 200}

	fmt.Println(rect1) //&{0 0 0 0}
	fmt.Println(rect2) //&{0 0 0 0}
	fmt.Println(rect3) //&{0 0 100 200}
	fmt.Println(rect4) //&{0 0 100 200}
}

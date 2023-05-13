package chapter_11

import "fmt"

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

func (r *Rectangle) Area() float32 {
	return r.width * r.length
}

func itf() {
	r := &Rectangle{5, 3}
	s := &Square{5}

	shapes := []Shaper{r, s}
	for i := range shapes {
		fmt.Println(shapes[i].Area())
	}

}

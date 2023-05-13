package chapter_11

type Shaper2 interface {
	Area() float32
	Perimeter() float32
}

type Square2 struct {
	side float32
}

func (s *Square2) Area() float32 {
	return s.side * s.side
}

func (s Square2) Perimeter() float32 {
	return s.side * 4
}

func mthds() {
	//s1 := Square2{2}
	//s2 := &Square2{3}
	//
	//var shape Shaper2
	// Cannot use 's1' (type Square2) as the type Shaper2
	// Type does not implement 'Shaper2' as the
	// 'Area' method has a pointer receiver
	// shape = s1
	//shape = s2

}

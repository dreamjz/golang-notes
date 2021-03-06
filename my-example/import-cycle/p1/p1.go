package p1

import (
	"fmt"
	"import-cycle-example/p2"
)

type PP1 struct {}

func New() *PP1 {
	return &PP1{}
}

func (p *PP1) HelloFromP1(){
	fmt.Println("Hello from p1")
}

func (p *PP1) HelloFromP2Side(){
	pp2 := p2.New(p)
	pp2.HelloFromP2()
}
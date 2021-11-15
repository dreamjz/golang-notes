package p2

import (
	"fmt"
)

type Pi interface {
	HelloFromP1()
}

type PP2 struct{
	PP1 Pi
}

func New(pp1 Pi) *PP2 {
	return &PP2{
		PP1: pp1,
	}
}

func (p *PP2) HelloFromP2(){
	fmt.Println("Hello from P2")
}

func (p *PP2) HelloFromP1Side() {
	p.PP1.HelloFromP1()
}

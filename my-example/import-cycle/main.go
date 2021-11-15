package main

import (
	"import-cycle-example/p1"
	"import-cycle-example/p2"
)

func main () {
	pp1 := &p1.PP1{}
	pp1.HelloFromP1()
	pp1.HelloFromP2Side()
	pp2 := &p2.PP2{
		PP1: pp1,
	}
	pp2.HelloFromP2()
	pp2.HelloFromP1Side()
}

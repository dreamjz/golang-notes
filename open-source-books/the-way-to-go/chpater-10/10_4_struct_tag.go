package chpater_10

import (
	"fmt"
	"reflect"
)

type TagType struct {
	f1 int  "Field01"
	f2 bool "Field02"
}

func structag() {
	tt := TagType{}
	refTag(tt, 0)
	refTag(tt, 1)
}

func refTag(tt TagType, i int) {
	ttType := reflect.TypeOf(tt)
	f := ttType.Field(i)
	fmt.Println(f.Tag)
}

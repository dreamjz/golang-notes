package utils

import (
	"log"
	"reflect"
)

func PrintRecord(record ...interface{}) {
	for _, r := range record {
		printSingle(r)
	}
}

func printSingle(record interface{}) {
	v := reflect.ValueOf(record)
	k := v.Kind()
	switch k {
	case reflect.Struct:
		log.Printf("Record: %+v", v)
	case reflect.Slice:
		log.Printf("Records: %+v", v)
	default:
		log.Printf("%#v", v)
	}
}

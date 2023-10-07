package _2_datastructure

import (
	"reflect"
	"testing"
)

func BenchmarkGetObjUsingNew(b *testing.B) {
	var res *Config
	for i := 0; i < b.N; i++ {
		res = new(Config)
	}
	_ = res
}

func BenchmarkGetObjUsingReflect(b *testing.B) {
	var res *Config
	t := reflect.TypeOf(Config{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res, _ = reflect.New(t).Interface().(*Config)
	}
	_ = res
}

func BenchmarkSet(b *testing.B) {
	cfg := &Config{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cfg.Name = "Name"
		cfg.IP = "IP"
		cfg.URL = "URL"
		cfg.Timeout = "Timeout"
	}
}

func BenchmarkSetUsingReflectByFieldName(b *testing.B) {
	typ := reflect.TypeOf(Config{})
	ins := reflect.New(typ).Elem()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ins.FieldByName("Name").SetString("Name")
		ins.FieldByName("IP").SetString("IP")
		ins.FieldByName("URL").SetString("URL")
		ins.FieldByName("Timeout").SetString("Timeout")
	}
}

func BenchmarkSetUsingReflectByFieldNameWithCache(b *testing.B) {
	typ := reflect.TypeOf(Config{})
	fCache := make(map[string]int, typ.NumField())

	for i := 0; i < typ.NumField(); i++ {
		fCache[typ.Field(i).Name] = i
	}

	ins := reflect.New(typ).Elem()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ins.Field(fCache["Name"]).SetString("Name")
		ins.Field(fCache["IP"]).SetString("IP")
		ins.Field(fCache["URL"]).SetString("URL")
		ins.Field(fCache["Timeout"]).SetString("Timeout")
	}
}

func BenchmarkSetUsingReflectByFieldIdx(b *testing.B) {
	typ := reflect.TypeOf(Config{})
	ins := reflect.New(typ).Elem()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ins.Field(0).SetString("Name")
		ins.Field(1).SetString("IP")
		ins.Field(2).SetString("URL")
		ins.Field(3).SetString("Timeout")
	}
}

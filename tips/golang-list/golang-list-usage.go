package main

import (
	"container/list"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("------------------------ 初始化 ------------------------")
	// 使用 list.New() 函数
	list1 := list.New()
	// 声明变量
	var list2 list.List
	list3 := new(list.List)

	fmt.Printf("Var: %+v \n", *list1)
	fmt.Printf("Var: %+v \n", list2)
	fmt.Printf("Var: %+v \n", *list3)

	fmt.Println("------------------------ 插入 ------------------------")
	var list4 list.List
	list4.PushBack(1)
	list4.PushBack("Hello")
	fmt.Println(joinList(list4))

	fmt.Println("------------------------ 删除 ------------------------")
	var list5 list.List
	list5.PushBack("A")
	list5.PushBack("B")
	list5.PushBack("C")
	fmt.Println(joinList(list5))
	e := list5.Front()
	list5.Remove(e)
	fmt.Println(joinList(list5))
}

func joinList(l list.List) string {
	var sb strings.Builder
	sb.WriteString("[")
	for e := l.Front(); e != nil; e = e.Next() {
		sb.WriteString(fmt.Sprintf("%v", e.Value))
		sb.WriteString(", ")
	}
	sb.WriteString("]")
	return sb.String()
}

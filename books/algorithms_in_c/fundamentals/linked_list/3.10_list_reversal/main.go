package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	item int
	next *Node
}

func (n *Node) String() string {
	var builder strings.Builder

	builder.WriteString("[")
	for x := n; x != nil; x = x.next {
		builder.WriteString(fmt.Sprintf("%d ", x.item))
	}
	builder.WriteString("]")
	return builder.String()
}

func (n *Node) Reverse() *Node {
	var r *Node
	for x := n; x != nil; {
		t := x.next
		x.next = r
		r = x
		x = t
	}
	return r
}

func NewLinkedList(len int) *Node {
	head := &Node{}

	x := head
	for i := 1; i < len; i++ {
		node := &Node{item: i}

		node.next = x.next
		x.next = node
		x = x.next
	}
	return head
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: ListReversal list_length")
	}
	len, _ := strconv.Atoi(os.Args[1])

	list := NewLinkedList(len)
	fmt.Println(list)
	reverseList := list.Reverse()
	fmt.Println(reverseList)
}

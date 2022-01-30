package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	item int
	next *Node
}

func (n *Node) InsertionSort() *Node {
	sortedHead := &Node{}

	for x := n.next; x != nil; {
		tmp := x.next
		var y *Node
		for y = sortedHead; y.next != nil; y = y.next {
			if x.item < y.next.item {
				break
			}
		}
		x.next = y.next
		y.next = x
		x = tmp
	}
	return sortedHead
}

func (n Node) String() string {
	var builder strings.Builder

	builder.WriteString("[")
	for x := n.next; x != nil; x = x.next {
		builder.WriteString(fmt.Sprintf("%d ", x.item))
	}
	builder.WriteString("]")

	return builder.String()
}

func NewRandomList(len int) *Node {
	head := &Node{}
	t := head
	for i := 0; i < len; i++ {
		n := &Node{item: rand.Intn(100)}
		t.next = n
		t = t.next
	}

	return head
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage ListInsertionSort list_length")
	}
	len, _ := strconv.Atoi(os.Args[1])

	list := NewRandomList(len)
	fmt.Println(list)
	sortedList := list.InsertionSort()
	fmt.Println(sortedList)
}

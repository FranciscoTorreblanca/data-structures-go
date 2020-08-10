package main

import (
	"fmt"

	"github.com/FranciscoTorreblanca/data-structures-go/structures"
)

func main() {
	// Linked List
	ll := structures.NewLinkedList(1)
	ll.Add(2)
	ll.Add(3)
	ll.Remove()
	fmt.Println(ll.First())
	fmt.Println(ll.Last())
	fmt.Println(ll.Length())

	// Queue
	q := structures.NewQueue(1)
	q.Add(2)
	q.Add(3)
	q.Remove()
	fmt.Println(q)

	// Stack
	s := structures.NewStack(1)
	s.Add(2)
	s.Add(3)
	s.Remove()
	fmt.Println(s)

}

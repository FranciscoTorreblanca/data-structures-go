package structures

import "fmt"

func NewLinkedList(v interface{}) *LinkedList {
	ll := new(LinkedList)
	ll.root = &LinkedListNode{value: v, next: nil}
	return ll
}

type LinkedListNode struct {
	value interface{}
	next  *LinkedListNode
}

func (n *LinkedListNode) String() string {
	return fmt.Sprintf("%v", n.value)
}
func (n *LinkedListNode) Next() *LinkedListNode {
	return n.next
}

func (ll *LinkedList) First() *LinkedListNode {
	return ll.root
}
func (ll *LinkedList) Last() *LinkedListNode {
	curr := ll.First()
	next := curr.next
	for next != nil {
		curr = next
		next = curr.next
	}
	return curr
}
func (ll *LinkedList) Length() (len int) {
	if ll.root != nil {
		len++
	}
	curr := ll.root
	next := curr.next
	for next != nil {
		len++
		curr = next
		next = curr.next
	}
	return
}
func (ll *LinkedList) Add(v interface{}) {
	lastn := ll.Last()
	newn := LinkedListNode{value: v}
	lastn.next = &newn
}

func (ll *LinkedList) Remove() {
	curr := ll.root
	prev := curr
	next := curr.next
	for next != nil {
		prev = curr
		curr = next
		next = curr.next
	}

	// Delete value
	curr.value = nil

	// Unlink node
	prev.next = nil
}

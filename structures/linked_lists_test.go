package structures

import (
	"testing"
)

const (
	lv1 = 1
	lv2 = "2"
	lv3 = false
)

func TestLinkedList(t *testing.T) {
	ll := NewLinkedList(lv1)
	if ll.root.value != lv1 {
		t.Errorf("NewLinkedList(%v): %v; want %v", lv1, ll.root.value, lv1)
	}
	if fn := ll.First(); fn.value != lv1 {
		t.Errorf("First(): %v; want %v", fn.value, lv1)
	}

	ll.Add(lv2)
	ll.Add(lv3)
	if nn := ll.root.Next(); nn.value != lv2 {
		t.Errorf("Next(): %v; want %v", nn.value, lv2)
	}
	if ln := ll.Last(); ln.value != lv3 {
		t.Errorf("Last(): %v; want %v", ln.value, lv3)
	}
	if l := ll.Length(); l != 3 {
		t.Errorf("Lenght(): %v; want %v", l, 3)
	}
	ll.Remove()

	if l := ll.Length(); l != 2 {
		t.Errorf("Length(): %v; want %v", l, 2)
	}
}

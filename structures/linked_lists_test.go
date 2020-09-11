package structures

import (
	"testing"
)

const (
	v1 = 1
	v2 = 2
)

func TestLinkedList(t *testing.T) {
	ll := NewLinkedList(v1)
	if ll.root.value != v1 {
		t.Errorf("NewLinkedList(%v): %v; want %v", v1, ll.root.value, v1)
	}
	if fn := ll.First(); fn.value != v1 {
		t.Errorf("First(): %v; want %v", fn.value, v1)
	}

	ll.Add(v2)
	if ln := ll.Last(); ln.value != v2 {
		t.Errorf("Last(): %v; want %v", ln.value, v2)
	}
	if l := ll.Length(); l != 2 {
		t.Errorf("Lenght(): %v; want %v", l, 2)
	}
	ll.Remove()

	if l := ll.Length(); l != 1 {
		t.Errorf("Length(): %v; want %v", l, 1)
	}
}

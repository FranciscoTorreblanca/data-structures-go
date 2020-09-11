package structures

import (
	"testing"
)

const (
	qv1 = "1"
	qv2 = 2
	qv3 = true
)

func TestQueue(t *testing.T) {
	q := NewQueue(qv1)
	if q.items[0] != qv1 {
		t.Errorf("NewQueue(%v): %v; want %v", qv1, q.items[0], qv1)
	}
	if fi := q.First(); fi != qv1 {
		t.Errorf("First(): %v; want %v", fi, qv1)
	}

	// q.Add(qv2)
	// q.Add(qv3)
	// if li := q.Last(); li != qv3 {
	// 	t.Errorf("Last(): %v; want %v", li, qv3)
	// }
	// if l := q.Length(); l != 3 {
	// 	t.Errorf("Lenght(): %v; want %v", l, 3)
	// }
	// q.Remove()

	// if l := q.Length(); l != 2 {
	// 	t.Errorf("Length(): %v; want %v", l, 2)
	// }

	// q.Remove()
	// q.Remove()
	// if l := q.Last(); l != nil {
	// 	t.Errorf("Last(): %v; want %v", l, nil)
	// }
}

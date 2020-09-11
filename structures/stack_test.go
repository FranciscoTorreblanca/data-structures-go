package structures

import (
	"testing"
)

const (
	sv1 = "1"
	sv2 = 2
	sv3 = true
)

func TestStack(t *testing.T) {
	s := NewStack(sv1)
	if s.items[0] != sv1 {
		t.Errorf("NewStack(%v): %v; want %v", sv1, s.items[0], sv1)
	}
	if fi := s.First(); fi != sv1 {
		t.Errorf("First(): %v; want %v", fi, sv1)
	}

	s.Add(sv2)
	s.Add(sv3)
	// if li := s.Last(); li != sv1 {
	// 	t.Errorf("Last(): %v; want %v", li, sv1)
	// }
	// if l := s.Length(); l != 3 {
	// 	t.Errorf("Lenght(): %v; want %v", l, 3)
	// }
	s.Remove()

	// if l := s.Length(); l != 2 {
	// 	t.Errorf("Length(): %v; want %v", l, 2)
	// }

	// s.Remove()
	// s.Remove()
	// if l := s.Last(); l != nil {
	// 	t.Errorf("Last(): %v; want %v", l, nil)
	// }
}

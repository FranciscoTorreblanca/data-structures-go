package structures

func NewStack(v interface{}) *Stack {
	s := new(Stack)
	s.items = append(s.items, v)
	return s
}

func (s *Stack) First() interface{} {
	return s.items[0]
}
func (s *Stack) Last() interface{} {
	l := s.Length()
	if l == 0 {
		return nil
	}
	return s.items[l-1]
}
func (s *Stack) Length() int {
	return len(s.items)
}
func (s *Stack) Add(v interface{}) {
	var sl []interface{}
	sl = append(sl, v)
	s.items = append(sl, s.items...)
}

func (s *Stack) Remove() {
	l := s.Length()
	s.items[0] = nil
	s.items = s.items[1:l]
}

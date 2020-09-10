package structures

type Queue struct {
	items []interface{}
}

func NewQueue(v interface{}) *Queue {
	q := new(Queue)
	q.items = append(q.items, v)
	return q
}

func (q *Queue) First() interface{} {
	return q.items[0]
}
func (q *Queue) Last() interface{} {
	l := q.Length()
	if l == 0 {
		return nil
	}
	return q.items[l-1]
}
func (q *Queue) Length() int {
	return len(q.items)
}
func (q *Queue) Add(v interface{}) {
	q.items = append(q.items, v)
}

func (q *Queue) Remove() {
	l := q.Length()
	q.items[l-1] = nil
	q.items = q.items[:l-1]
}

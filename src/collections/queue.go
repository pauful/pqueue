package collections

type Item struct {
	next *Item
	val  interface{}
}

type Queue struct {
	first, last *Item
	len         int
}

func NewQueue() *Queue {
	return &Queue{len: 0}
}

func (q *Queue) Len() int {
	return q.len
}

func (q *Queue) Add(item interface{}) {
	itemToInsert := &Item{val: item}
	if q.len == 0 {
		q.first = itemToInsert
	} else {
		q.last.next = itemToInsert
	}
	q.last = itemToInsert
	q.len++
}

func (q *Queue) Remove() interface{} {
	if q.len == 0 {
		return nil
	}
	first := q.first
	q.first = q.first.next
	q.len--
	return first.val
}

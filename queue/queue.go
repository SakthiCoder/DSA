package queue

type Queue struct {
	items []interface{}
}

func (q *Queue) Push(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Pop() (interface{}, bool) {

	if len(q.items) == 0 {
		return -1, false
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Peek() interface{} {

	if len(q.items) == 0 {
		return -1
	}

	return q.items[0]
}

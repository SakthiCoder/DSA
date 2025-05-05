package queue

type Queue struct {
	items []any
}

func (q *Queue) Push(item any) {
	q.items = append(q.items, item)
}

func (q *Queue) Pop() (any, bool) {

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

func (q *Queue) Peek() any {

	if len(q.items) == 0 {
		return -1
	}

	return q.items[0]
}

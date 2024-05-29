package queue

type Queue[T any] struct {
	buffer []*T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		buffer: make([]*T, 0),
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.buffer = append(q.buffer, &item)
}

func (q *Queue[T]) Dequeue() (*T, bool) {
	if len(q.buffer) == 0 {
		return nil, false
	}
	item := q.buffer[0]
	q.buffer = q.buffer[1:]
	return item, true
}

func (q *Queue[T]) Size() int {
	return len(q.buffer)
}

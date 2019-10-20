package queue

import "fmt"

// ArrayQueue define a queue
type ArrayQueue struct {
	data     []interface{}
	capacity int
	head     int
	tail     int
}

// NewArrayQueue is init a array queue
func NewArrayQueue(c int) *ArrayQueue {
	return &ArrayQueue{
		data:     make([]interface{}, c),
		capacity: c,
		head:     0,
		tail:     0,
	}
}

// Enqueue is append a element
func (q *ArrayQueue) Enqueue(v interface{}) bool {
	if q.tail == q.capacity {
		if q.head == 0 {
			return false
		}
		for i := q.head; i < q.tail; i++ {
			q.data[i-q.head] = q.data[i]
		}

		q.data[q.tail] = v
		q.tail++
		return false
	}

	q.data[q.tail] = v
	q.tail++
	return true
}

// Dequeue is pop a element
func (q *ArrayQueue) Dequeue() (interface{}, bool) {
	if q.head == q.tail {
		return 0, false
	}
	v := q.data[q.head]
	q.head++
	return v, true
}

// Print is list all data of queue
func (q *ArrayQueue) Print() {
	for i := q.head; i < q.tail; i++ {
		fmt.Println(q.data[i])
	}
}

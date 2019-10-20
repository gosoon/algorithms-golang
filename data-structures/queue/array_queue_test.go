package queue

import (
	"testing"
)

func TestNewArrayQueue(t *testing.T) {
	q := NewArrayQueue(5)
	if ok := q.Enqueue(3); !ok {
		t.Log("enqueue failed")
	}
	if ok := q.Enqueue(5); !ok {
		t.Log("enqueue failed")
	}

	q.Print()

	if _, ok := q.Dequeue(); !ok {
		t.Log("failed dequeue")
	}
	q.Print()
}

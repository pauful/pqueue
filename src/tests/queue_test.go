package queue_tests

import (
	"collections"
	"testing"
)

func checkQueueLen(t *testing.T, q *collections.Queue, expected int) {
	if q.Len() != expected {
		t.Errorf("Queue len is '%d' instead of '%d'", q.Len(), expected)
	}
}

func TestQueue_New(t *testing.T) {
	queue := collections.New()
	if queue == nil {
		t.Errorf("Queue New returns a nil value")
	}
}

func TestQueue_New_Empty(t *testing.T) {
	queue := collections.New()
	checkQueueLen(t, queue, 0)
}

func TestQueue_Add(t *testing.T) {
	queue := collections.New()
	queue.Add(1)
	checkQueueLen(t, queue, 1)
}

func TestQueue_RemoveReturnsLastElement(t *testing.T) {
	queue := collections.New()
	queue.Add(1)
	lastElement := queue.Remove()
	if lastElement != 1 {
		t.Errorf("Element value '%d' not expectd '%d'", lastElement, 1)
	}
	checkQueueLen(t, queue, 0)
}

func TestQueue_RemoveEmptyListResultsNil(t *testing.T) {
	queue := collections.New()
	lastElement := queue.Remove()
	if lastElement != nil {
		t.Errorf("Element value '%v' not expectd '%v'", lastElement, nil)
	}
	checkQueueLen(t, queue, 0)
}

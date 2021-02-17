package collections

import (
	"time"
)

type record struct {
	Data  []byte
	Stamp int64
}

type QueuesManager struct {
	Queues map[string]*Queue
}

func (q *QueuesManager) Exists(name string) bool {
	_, ok := q.Queues[name]
	return ok
}

func (q *QueuesManager) Push(name string, data []byte) {
	newItem := &record{Data: data, Stamp: time.Now().Unix()}
	if q.Exists(name) {
		q.Queues[name].Add(newItem)
	} else {
		q.Queues[name] = NewQueue()
		q.Queues[name].Add(newItem)
	}
}

func (q *QueuesManager) Len(name string) int {
	return q.Queues[name].Len()
}

func (q *QueuesManager) Pop(name string) []byte {
	if q.Exists(name) {
		front := q.Queues[name].Remove()
		return front.(*record).Data
	} else {
		return nil
	}
}

func NewQueuesManager() *QueuesManager {
	qm := new(QueuesManager)
	qm.Queues = make(map[string]*Queue)
	return qm
}

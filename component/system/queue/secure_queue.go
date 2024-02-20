package queue

import (
	"container/list"
	"sync"
)

type SecureQueue struct {
	data *list.List
	size int
	lock sync.Mutex
}

func MakeDefault() *SecureQueue {
	q := new(SecureQueue)
	q.init()
	return q
}

func (q *SecureQueue) init() {
	q.data = list.New()
}

func (q *SecureQueue) Size() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.size
}

func (q *SecureQueue) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.size == 0
}

func (q *SecureQueue) Front() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.data.Front().Value
}

func (q *SecureQueue) Back() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.data.Back().Value
}

func (q *SecureQueue) Push(value ...interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	for _, single := range value {
		q.data.PushBack(single)
	}
	q.size++
}

func (q *SecureQueue) Pop() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.size > 0 {
		tmp := q.data.Back()
		q.data.Remove(tmp)
		q.size--
		return tmp
	} else {
		return nil
	}

}

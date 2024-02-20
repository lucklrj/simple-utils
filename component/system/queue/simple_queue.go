package quenue

import (
	"errors"
	"sync"
)

type SimpleQueue struct {
	Data []interface{}
	lock sync.Mutex
	len  int
}

func (q *SimpleQueue) Dequeue() (interface{}, error) {
	if q.GetLength() == 0 {
		return nil, errors.New("empty slice")
	}
	q.lock.Lock()
	defer func() {
		q.Data = q.Data[1:]
		q.lock.Unlock()
	}()

	return q.Data[0], nil
}

func (q *SimpleQueue) Enqueue(source ...interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.Data = append(q.Data, source...)
}
func (q *SimpleQueue) GetLength() int {
	q.lock.Lock()
	defer func() {
		q.lock.Unlock()
	}()
	return len(q.Data)
}
func (q *SimpleQueue) Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.Data = make([]interface{}, 0)
}
func (q *SimpleQueue) Insert(point int, objs ...interface{}) error {
	if point > q.GetLength() {
		return errors.New("slice bounds out of range")
	}
	if point < 1 {
		point = 0
	}

	q.lock.Lock()
	defer q.lock.Unlock()

	left := append([]interface{}{}, q.Data[0:point]...)
	right := append([]interface{}{}, q.Data[point:]...)

	q.Data = append(append(append(left, objs...)), right...)
	return nil
}

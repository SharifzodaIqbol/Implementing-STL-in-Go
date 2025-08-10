package main

import (
	"errors"
	"fmt"
)

type Queue[T any] struct {
	items []T
}

// Enqueue добавляет элемент в конец очереди (аналог C++ push).
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue удаляет и возвращает первый элемент (аналог C++ pop).
func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if len(q.items) == 0 {
		return zero, errors.New("queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}

// Front возвращает первый элемент без удаления (аналог C++ front).
func (q *Queue[T]) Front() (T, error) {
	var zero T
	if len(q.items) == 0 {
		return zero, errors.New("queue is empty")
	}
	return q.items[0], nil
}

// Back возвращает последний элемент (аналог C++ back).
func (q *Queue[T]) Back() (T, error) {
	var zero T
	if len(q.items) == 0 {
		return zero, errors.New("queue is empty")
	}
	return q.items[len(q.items)-1], nil
}

// IsEmpty проверяет, пуста ли очередь (аналог C++ empty).
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Size возвращает количество элементов (аналог C++ size).
func (q *Queue[T]) Size() int {
	return len(q.items)
}

func main() {
	q := Queue[int]{}
	q.Enqueue(10)
	q.Enqueue(20)
	fmt.Println(q.Front()) // 10, nil
	fmt.Println(q.Back())  // 20, nil
	val, _ := q.Dequeue()
	fmt.Println(val)      // 10
	fmt.Println(q.Size()) // 1
}

package main

import (
	"fmt"
)

type Stack[T any] []T

// Push добавляет элемент на вершину стека.
func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

// Pop удаляет и возвращает верхний элемент стека. Если стек пуст, возвращает ошибку.
func (s *Stack[T]) Pop() error {
	if len(*s) == 0 {
		panic("stack is empty")
	}
	*s = (*s)[:len(*s)-1]
	return nil
}

// Top возвращает верхний элемент стека без удаления.
func (s Stack[T]) Top() T {
	if len(s) == 0 {
		panic("stack is empty")
	}
	return s[len(s)-1]
}

// Size возвращает количество элементов в стеке.
func (s Stack[T]) Size() int {
	return len(s)
}

// Empty проверяет, пуст ли стек.
func (s Stack[T]) Empty() bool {
	return len(s) == 0
}
func main() {
	var stack Stack[int]

	stack.Push(4)
	stack.Push(5)
	stack.Push(6)

	fmt.Println(stack.Top()) // 6

	stack.Pop() // удаляем 6

	fmt.Println(stack.Top())   // 5
	fmt.Println(stack.Size())  // размер стека = 2
	fmt.Println(stack.Empty()) // false
}

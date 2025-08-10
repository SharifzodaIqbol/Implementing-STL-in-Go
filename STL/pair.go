package main

import (
	"fmt"
)

type Pair[T comparable, U comparable] struct {
	first  T
	second U
}

// NewPair — аналог std::make_pair, создает новую пару
func NewPair[T comparable, U comparable](first T, second U) Pair[T, U] {
	return Pair[T, U]{first, second}
}

// First — возвращает первый элемент (аналог .first в C++)
func (p Pair[T, U]) First() T {
	return p.first
}

// Second — возвращает второй элемент (аналог .second в C++)
func (p Pair[T, U]) Second() U {
	return p.second
}

// SetFirst — устанавливает первый элемент
func (p *Pair[T, U]) SetFirst(val T) {
	p.first = val
}

// SetSecond — устанавливает второй элемент
func (p *Pair[T, U]) SetSecond(val U) {
	p.second = val
}

// String — строковое представление (аналог operator<<)
func (p Pair[T, U]) String() string {
	return fmt.Sprintf("(%v %v)", p.first, p.second)
}

// Equals — проверка на равенство (аналог operator==)
func (p Pair[T, U]) Equals(other Pair[T, U]) bool {
	return p.first == other.first && p.second == other.second
}

// Less — проверка "меньше" (аналог operator<)
func (p Pair[T, U]) Less(other Pair[T, U]) bool {
	switch {
	case p.first != other.first:
		return fmt.Sprintf("%v", p.first) < fmt.Sprintf("%v", other.first)
	default:
		return fmt.Sprintf("%v", p.second) < fmt.Sprintf("%v", other.second)
	}
}
func main() {
	pr1 := NewPair(18, "age")
	fmt.Println(pr1)

	fmt.Println(pr1.First())  // 18
	fmt.Println(pr1.Second()) // age

	pr1.SetFirst(42)
	pr1.SetSecond("world")
	fmt.Println(pr1) // (42, world)

	// Проверка на равенство
	pr2 := NewPair(42, "world")
	fmt.Println(pr1.Equals(pr2)) // true

	// Сравнение
	p3 := NewPair(1, "apple")
	p4 := NewPair(2, "banana")
	fmt.Println(p3.Less(p4)) // true
}

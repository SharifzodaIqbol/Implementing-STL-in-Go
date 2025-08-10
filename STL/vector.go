package main

import (
	"fmt"
	"strings"
)

type Vector[T any] []T

func (v *Vector[T]) Push_back(val T) {
	*v = append(*v, val)
}
func (v *Vector[T]) Size() int {
	return len(*v)
}
func (v *Vector[T]) Empty() bool {
	return len(*v) == 0
}
func (v *Vector[T]) Clear() {
	*v = Vector[T]{}
}
func (v *Vector[T]) Insert(position int, value T) error {
	if position < 0 || position > len(*v) {
		return fmt.Errorf("Index out of range")
	}
	*v = append(*v, value)
	copy((*v)[position+1:], (*v)[position:])
	return nil
}
func (v *Vector[T]) Capacity() int {
	return cap(*v)
}

func (v Vector[T]) String() string {
	var builder strings.Builder
	for i, value := range v {
		if i > 0 {
			builder.WriteString(", ")
		}
		fmt.Fprintf(&builder, "%v", value)
	}
	return builder.String()
}
func main() {
	myVector := Vector[string]{"Hello"}
	myVector.Push_back("World!")

	fmt.Println(myVector)                  // сам вектор
	fmt.Println("Size:=", myVector.Size()) // размер

	myVector.Clear() // очищаем вектор
	if myVector.Empty() {
		fmt.Println("Vector Empty")
		fmt.Println("Size:=", myVector.Size()) // размер
	} else {
		fmt.Println(myVector)
		fmt.Println("Vector not Empty")
		fmt.Println("Size:=", myVector.Size())
	}

	myVector.Push_back("Hello")
	myVector.Push_back("World!")

	if myVector.Empty() {
		fmt.Println("Vector Empty")
		fmt.Println("Size:=", myVector.Size())
	} else {
		fmt.Println(myVector)
		fmt.Println("Vector not Empty")
		fmt.Println("Size:=", myVector.Size())
	}
	myVector.Insert(2, "Accept") /// Вставим Accept на 2-й индекс
	fmt.Println(myVector)

	fmt.Println("Capacity :=", myVector.Capacity()) /// проверим Capacity(вместимость)
}

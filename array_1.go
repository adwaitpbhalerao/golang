package main

import (
	"fmt"
)

// Stack :
type Stack struct {
	arr []int
}

//type stackInterface interface {
///	push(int) string
//	pop() string
//	//setSize()
//}

// Push :
func (s *Stack) Push(newVal int) string {
	// fmt.Println("Push")

	for i := 0; i < len(s.arr); i++ {

		if s.arr[i] != 0 && i == cap(s.arr)-1 {

			fmt.Println("array full!!")
			return fmt.Sprintf("%v", s.arr)

		}
		if s.arr[i] == 0 {

			s.arr[i] = newVal
			return fmt.Sprintf("%v", s.arr)

		}
	}

	return fmt.Sprintf("%v", s.arr)

}

// Pop :
func (s *Stack) Pop() string {
	// fmt.Println("Pop")

	if s.arr[0] == 0 {

		fmt.Println("Array Empty")
		return fmt.Sprintf("%v", nil)

	}

	for i := 0; i < len(s.arr); i++ {

		if s.arr[i] == 0 {

			s.arr[i-1] = 0
			return fmt.Sprintf("%v", s.arr)

		}
	}

	return fmt.Sprintf("%v", s.arr)

}

//func (s *Stack) setSize(size int) {
	// fmt.Println("setSize")

	//s.arr = make([]int, size)

//}

func main() {
	//size := Size(10)
  //ask producers and set the below to instead of 5
	s := Stack{arr: make([]int, 5)}
	//s.setSize(10)
	fmt.Println(s.Push(20))
	fmt.Println(s.Push(22))
	fmt.Println(s.Push(245))
	fmt.Println(s.Push(67))
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

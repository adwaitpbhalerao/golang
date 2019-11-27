// necassary functions to implement a stack in go


package main 
import "fmt"

// stack struct

type Stack struct{
	top *Node
	length int
}

// function to return length of stack
func (s *Stack) Length() int{
	return s.length
}

// Push function
func (s *Stack) Push(val interface{}){
	n := &Node{value: val}
	if s.top == nil {
		s.top = n
	} else {
		n.SetPrevious(s.top)
		s.top = n
	}
	s.length += 1
}

// Pop function
func (s *Stack) Pop() interface{}{
	if s.top == nil {
		return nil
	}
	currentTop := s.top
	s.top = currentTop.Previous()
	currentTop.SetPrevious(nil)
	s.length -= 1

	return currentTop.Value()
}

// print function for stack
func (s *Stack) Print(){
	for n :=s.top; n != nil; n = n.Previous(){
		fmt.Println(n.Value())
	}
}
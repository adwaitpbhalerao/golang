package main

import "fmt"

// CirCircularQueue
//   capacity is the numbers of this queue's ability (capacity - 1)
type CircularQueue struct {
  capacity int
  elements []interface{}
  first    int
  end      int
}

// NewCircularQueue create new CircularQueue passing a integer as its size
// and return its pointer
func NewCircularQueue(size int) *CircularQueue {
  cq := CircularQueue{capacity: size + 1, first: 0, end: 0}
  cq.elements = make([]interface{}, cq.capacity)
  return &cq
}

// IsEmpty return if this queue is empty
func (c CircularQueue) IsEmpty() bool {
  return c.first == c.end
}

// IsFull return if this queue is full
func (c CircularQueue) IsFull() bool {
  return c.first == (c.end+1)%c.capacity
}

// Push pushing a element to this queue
// note: if pushing into a full queue, it will panic
func (c *CircularQueue) Push(e interface{}) {
  if c.IsFull() {
    panic("Queue is full")
  }
  c.elements[c.end] = e
  c.end = (c.end + 1) % c.capacity
}

// Shift shift a element witch pushed earlist
// note: if will return nil if this queue is empty
func (c *CircularQueue) Shift() (e interface{}) {
  if c.IsEmpty() {
    return nil
  }
  e = c.elements[c.first]
  c.first = (c.first + 1) % c.capacity
  return
}

func main() {
  fmt.Println("Hello World")
  cq := NewCircularQueue(3)
  cq.Push(1)
  cq.Push(2)
  cq.Push(3)
  cq.Shift()
  cq.Push(3)
  fmt.Println(*cq)
}
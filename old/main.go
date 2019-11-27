package main

import "fmt"


func main() {
	var l LinkedList
	l.Insert("element 1")
	l.Insert("element 2")
	l.Insert("element 3")
	l.Insert("element 4")
	l.Insert("element 5")
	l.Print()

	fmt.Println(l.Length())
	
  for i := 1; i < 6; i ++{
    l.Delete(1)
    fmt.Println(l.Length())
    l.Print()
  }
  l.Print()
  fmt.Println(l.Length())
  
	// l.Delete(3)
	// l.Print()
	// l.Delete(4)
	// l.Print()
	// l.Delete(10)
	// l.Print()
}
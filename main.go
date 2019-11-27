// to build and run this 
// go build main.go linkedList.go stack.go node.go
// .\main

package main

import "fmt"
import "time"
import "strconv"

func main() {
	//-----------------------LINKED LIST --------------//
	var l LinkedList

	var number int
	fmt.Println("Enter in number of producers/consumers: ")
    fmt.Scanf("%d", &number)

	start := time.Now()

	fmt.Println("Producing")
	for i:= 1; i < number+1; i++ {
		str := strconv.Itoa(i)
		str = "element " + str
		l.Insert(str)
		l.Print()
	}
	

	fmt.Println("Consuming")
	for k:= 1; k < number+1; k++ {
		l.Pop()
		l.Print()
	}
	fmt.Println("Printing empty Linked List: ")
	l.Print()
	
	elapsed := time.Since(start)

	fmt.Printf("linked list took %s", elapsed)

	
	// l.Circular()
	// l.Insert("element 6")
	// fmt.Println(l.Length())
	// for i := 1; i < 9; i ++{
	// 	fmt.Println(l.Get(i))
	// }
	
	

	
	//-------------------STACK-------------------------//
	// start := time.Now()
	// fmt.Println("Testing the stack")
	// var s Stack
	// s.Push("element 1")
	// s.Push("element 2")
	// s.Push("element 3")
	// s.Push("element 4")
	// s.Print()

	// fmt.Println("stack length: ", s.Length())
	// s.Pop()
	// s.Pop()
	// s.Print()
	// fmt.Println("stack length: ", s.Length())
	
	// elapsed := time.Since(start)
	// fmt.Printf("Stack took %s", elapsed)

}



/*
Measure time
start := time.Now()

elapsed := time.Since(start)
	log.Printf("Stack took %s", elapsed)
*/
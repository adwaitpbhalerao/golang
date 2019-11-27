package main

import "fmt"
import "time"

func main() {

    fmt.Println("This is a Stack Implementation using Array")

    var stack []int
  
	start := time.Now()
    prod := make(chan []int) 
    cons := make(chan []int)
    go push(&stack, prod) // calling push function
    <- prod
    fmt.Println("Producing", stack)
    //stack = pop(stack)
    go pop(stack, cons)
    <- cons
    //time.Sleep(2 * time.Second)
    // fmt.Println(stack[0])
    fmt.Println("Consuming", stack)
    elapsed := time.Since(start)
    fmt.Printf("linked list took %s", elapsed)
}

func push(stack *[]int, prodc chan []int ) {
var number int
fmt.Println("Enter in number of producers/consumers: ")
fmt.Scanf("%d", &number)
producer := number
//producer := number
i := 0
for len(*stack) < producer {
*stack = append(*stack, 1)
    //fmt.Println(*stack)
i++
}
prodc <- *stack 
}

func pop(stack []int, consm chan []int) {
	j := len(stack)
  for j > 0 {
		stack[j-1] = 0 
		//fmt.Println(stack)
		j--
		}
	  //return stack
    consm <- stack
	}
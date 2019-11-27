package main

import "fmt"
import "time"

func main() {

    fmt.Println("This is a Stack Implementation using Array")

    var stack []int
  
	start := time.Now()
    prod := make(chan []int) // creating channel for producer
    cons := make(chan []int) // creating channel for consumer
    go push(&stack, prod) // calling go routine on push function
    <- prod //
    fmt.Println("Producing")
    go pop(stack, cons) // calling go routine on pop function
    <- cons // 
    fmt.Println("Consuming")
    elapsed := time.Since(start)
    fmt.Printf("linked list took %s", elapsed)
}

func push(stack *[]int, prodc chan []int ) {
	var number int
	fmt.Println("Enter in number of producers/consumers: ")
	fmt.Scanf("%d", &number)
	producer := number
	i := 0
	for len(*stack) < producer {
		*stack = append(*stack, 1)
    	i++
		}
	prodc <- *stack 
}

func pop(stack []int, consm chan []int) {
	j := len(stack)
  	for j > 0 {
		stack[j-1] = 0 
		j--
		}
	consm <- stack
}
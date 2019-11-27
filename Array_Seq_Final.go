package main

import "fmt"
import "time"

func main() {

    fmt.Println("This is a Stack Implementation using Array")

    var stack []int
	
	start := time.Now()
    
	push(&stack) // calling push function
	fmt.Println("Producing")
    pop(stack) // calling pop function
    fmt.Println("Consuming")
	elapsed := time.Since(start)
    fmt.Printf("linked list took %s", elapsed)
}

func push(stack *[]int) {
	var number int
	fmt.Println("Enter in number of producers/consumers: ")
	fmt.Scanf("%d", &number)
		producer := number
		i := 0
		for len(*stack) < producer {
		*stack = append(*stack, 1)
  		  i++
		}
}

func pop(stack []int) {
	j := len(stack)
	for j > 0 {
   		stack[j-1] = 0 
        j--
    }
}
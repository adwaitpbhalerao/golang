package main

import "fmt"

func main() {

    fmt.Println("This is a Stack Implementation using Array")

    var stack []string
	
	start := time.Now()
    
    push(&stack) // calling push function
    fmt.Println("After Producer Acts", stack)
    stack = pop(stack)
    // fmt.Println(stack[0])
    fmt.Println("After Consumer Acts",stack)
	elapsed := time.Since(start)
    fmt.Printf("linked list took %s", elapsed)
}

func push(stack *[]string) {
var number int
fmt.Println("Enter in number of producers/consumers: ")
fmt.Scanf("%d", &number)
producer := number
producer := number
i := 0
for len(*stack) < producer {
*stack = append(*stack, "e")
    fmt.Println(*stack)
i++
}
}

func pop(stack []string) []string {
  for len(stack) > 0 {
    stack = stack[:len(stack)-1] 
    fmt.Println(stack)
    }
  return stack
}

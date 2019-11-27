package main

import "fmt"
import "time"

func main() {

    fmt.Println("This is a Stack Implementation using Array")

    var stack []int
	
	start := time.Now()
    
    push(stack) // calling push function
    fmt.Println("After Producer Acts", stack)
    //stack = pop(stack)
    pop(stack)
    // fmt.Println(stack[0])
    fmt.Println("After Consumer Acts",stack)
	elapsed := time.Since(start)
    fmt.Printf("linked list took %s", elapsed)
}

//func push(stack *[]int) {
//var number int
//fmt.Println("Enter in number of producers/consumers: ")
//fmt.Scanf("%d", &number)
//producer := number
//producer := number
//i := 0
//for len(*stack) < producer {
//*stack = append(*stack, 1)
//    fmt.Println(*stack)
//i++
//}
//}

//func pop(stack []string) []string {
//  for len(stack) > 0 {
//    stack = stack[:len(stack)-1] 
//    fmt.Println(stack)
//    }
////  return stack
//}

func push(stack []int) {
var number int
fmt.Println("Enter in number of producers/consumers: ")
fmt.Scanf("%d", &number)
producer := number
stack = make([]int, producer)
i := 0
for i < producer {
stack[i] = 1
fmt.Println(stack)
i++
}
}


func pop(stack []int) {
j := len(stack)
for j > 0 {
    stack[j-1] = 0 
    fmt.Println(stack)
    j--
    }
  //return stack
}
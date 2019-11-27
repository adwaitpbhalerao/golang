package main

import (
	"fmt"
)

type circularBuffer struct {
	dataArr []int
}

type cbInterface interface {
	push(int)
	pop() int
	setSize(int)
}

func (cbList *circularBuffer) push(newData int) {

	if cbList == nil {
		cbList.dataArr = append(cbList.dataArr, newData)
		fmt.Printf("First value added")
	}

	for i := 0; i < len(cbList.dataArr); i++ {

		if cbList.dataArr[i] == 0 {
			cbList.dataArr[i] = newData
			break
		}
	}
}

func (cbList *circularBuffer) pop() (poppedValue int) {

	if cbList == nil {
		fmt.Printf("Buffer empty")
	}

	for i := 0; i < len(cbList.dataArr); i++ {

		if cbList.dataArr[i] != 0 {
			poppedValue = cbList.dataArr[i]
			cbList.dataArr[i] = 0
			break
		}
	}
	if poppedValue == 0 {
		fmt.Println("Buffer empty")
	}

	return
}

func (cbList *circularBuffer) setSize(size int) {

	newArr := make([]int, size)
	if len(cbList.dataArr) > 0 {

		for i := 0; i < len(cbList.dataArr); i++ {

			// newArr = append(newArr, cbList.dataArr[i])
			newArr[i] = cbList.dataArr[i]
		}
	}
	cbList.dataArr = newArr
}

func main() {
	cbList := circularBuffer{make([]int, 5)}

	cbList.push(50)
	cbList.push(45)
	cbList.push(67)
	fmt.Println(cbList.pop())
	fmt.Println(cbList.pop())
	// cbList.pop()
	// cbList.pop()
	cbList.setSize(10)
	fmt.Println(cbList)
}

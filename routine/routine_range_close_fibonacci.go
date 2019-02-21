package routine

import "fmt"

// MyFibInterface is the interface to be implemented
type MyFibInterface interface {
	GenerateFibonacciSeq(fibNum int) []int
}

// myFib is the concrete struct that implements the interface
type myFib struct {
}

// fibonacci contains the actual logic of the Fibonacci sequence
func (mf *myFib) fibonacci(fibNum int, channel chan int) {
	x, y := 0, 1
	for i := 0; i < fibNum; i++ {
		channel <- x
		x, y = y, x+y
	}
	defer close(channel)
}

// GenerateFibonacciSeq generates the Fibonacci sequence
func (mf *myFib) GenerateFibonacciSeq(fibNum int) []int {
	channel := make(chan int, fibNum)

	go mf.fibonacci(cap(channel), channel)

	var fibSeq []int
	for value := range channel {
		fmt.Println("GenerateFibonacciSeq value:", value)
		fibSeq = append(fibSeq, value)
	}

	return fibSeq
}

// NewMyFibonacci is the factory function that instantiates the myCounter through its constructor
func NewMyFibonacci() MyFibInterface {
	var myFibInterface MyFibInterface = &myFib{}
	return myFibInterface
}

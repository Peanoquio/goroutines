package main

import (
	"fmt"
	"time"

	"github.com/Peanoquio/goroutines/routine"
)

func main() {
	// NewMyCounter
	fmt.Println("===== NewMyCounter =====")
	myCounterObj := routine.NewMyCounter(0)
	for i := 0; i < 100; i++ {
		go myCounterObj.Inc()
	}
	time.Sleep(time.Second * 1)
	fmt.Println("final counter value:", myCounterObj.Get())

	// NewMyFibonacci
	fmt.Println("===== NewMyFibonacci =====")
	myFibObj := routine.NewMyFibonacci()
	var fibSeq []int = myFibObj.GenerateFibonacciSeq(10)

	for i, v := range fibSeq {
		fmt.Printf("index: %d value: %d \n", i, v)
	}

	// NewMyRandomStuff
	fmt.Println("===== NewMyRandomStuff =====")
	myRandomStuffObj := routine.NewMyRandomStuff()
	textList := []string{"a", "b", "c", "x", "y", "z"}
	myRandomStuffObj.StoreText(textList)
}

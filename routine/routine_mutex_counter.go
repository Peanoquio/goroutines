package routine

import (
	"fmt"
	"sync"
)

// MyCounterInterface is the exported interface that needs to be implemented
type MyCounterInterface interface {
	Inc() int
	Get() int
}

// MyCounter struct that will contain the atomic counter
type myCounter struct {
	counter int
	// data enclosed between Mutex Lock and Unlock will be atomic (only one Go Routine can access and modify it at a time)
	// similar to synchronize() for Java threads
	mux sync.Mutex
}

// Inc will increment the counter and then return the value
func (mc *myCounter) Inc() int {
	mc.mux.Lock()
	fmt.Println("myCounter Inc current counter value:", mc.counter)
	mc.counter++
	defer mc.mux.Unlock()
	return mc.counter
}

// Get will get the current value of the counter
func (mc *myCounter) Get() int {
	mc.mux.Lock()
	defer mc.mux.Unlock()
	return mc.counter
}

// NewMyCounter is the factory function that instantiates the myCounter through its constructor
func NewMyCounter(counter int) MyCounterInterface {
	var myCounterInterface MyCounterInterface = &myCounter{counter: 0}
	return myCounterInterface
}

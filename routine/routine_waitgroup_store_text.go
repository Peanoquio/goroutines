package routine

import (
	"fmt"
	"sync"
)

// MyRandomStuffInterface is the interface to be implemented
type MyRandomStuffInterface interface {
	StoreText(textList []string)
}

// myRandomStuff is the concrete struct that implements the interface
type myRandomStuff struct {
	textList []string
	// data enclosed between Mutex Lock and Unlock will be atomic (only one Go Routine can access and modify it at a time)
	// similar to synchronize() for Java threads
	mux *sync.Mutex
}

// addTextToList will add the text to the list
func (mrs *myRandomStuff) addTextToList(text string) {
	mrs.mux.Lock()
	mrs.textList = append(mrs.textList, text)
	fmt.Println("addTextToList textList:", mrs.textList)
	defer mrs.mux.Unlock()
}

// StoreText will store the text
func (mrs *myRandomStuff) StoreText(textList []string) {
	// wait group is similar to NodeJS Promise.all()
	wg := sync.WaitGroup{}

	for _, value := range textList {
		// must add each Go routine count to the wait group
		wg.Add(1)
		// closure as a Go routine
		go func(text string) {
			defer wg.Done()
			mrs.addTextToList(text)
		}(value)
	}

	// wait for all go routines to complete
	wg.Wait()

	fmt.Println("StoreText final textList:", mrs.textList)
}

// NewMyRandomStuff is the factory function that instantiates the myRandomStuff through its constructor
func NewMyRandomStuff() MyRandomStuffInterface {
	var myRandomStuffInterface MyRandomStuffInterface = &myRandomStuff{mux: &sync.Mutex{}}
	return myRandomStuffInterface
}

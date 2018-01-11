package main

import (
	"sync"
	"fmt"
	"sync/atomic"
	"runtime"
)

var (
	counter int64

	wg sync.WaitGroup
)

func main()  {
	wg.Add(2)
	go intcre(counter)
	go intcre(counter)
	wg.Wait()

	fmt.Println("Final count:",counter)
}

func intcre(id int64)  {
	defer  wg.Done()

	for count := 0;count<2;count++{
		atomic.AddInt64(&counter,1)

		runtime.Gosched()//exit the thread,return to quene
	}
	fmt.Println(counter)
}

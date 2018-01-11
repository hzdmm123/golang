package main

import (
	"sync"
	"fmt"
	"runtime"
)

var (
	counter1 int
	wg1 sync.WaitGroup
	mutex sync.Mutex
)

func main()  {
	wg1.Add(2)

	go test(counter1)
	go test(counter1)

	wg1.Wait()

	fmt.Println("final count:",counter1)
}

func test(i int)  {
	defer wg1.Done()

	for count:=0;count<2 ;count++  {
		mutex.Lock()
		{
			value := counter1

			runtime.Gosched()

			value++

			counter1 = value
		}
		mutex.Unlock()//临界区加锁（信号量）
	}
}






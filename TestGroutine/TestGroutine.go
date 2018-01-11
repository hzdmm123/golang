package main

import (
	"runtime"
	"sync"
	"fmt"
)

func main()  {
	runtime.GOMAXPROCS(2)//一个逻辑调度器和两个逻辑调度器是不一样的
	var wg sync.WaitGroup

	wg.Add(2)

	fmt.Println("start goroutine")

	go func() {
		defer wg.Done()//通知main函数这已经完成

		for count:=0;count<3 ;count++  {
			fmt.Println(count)
		}
	}()

	go func() {
		defer  wg.Done()//释放一个值

		for count:=5;count<10 ;count++  {
			fmt.Println(count)
		}
	}()

	fmt.Println("waiting finished")

	wg.Wait()//只要wg的值>0 这个方法就会阻塞等待

	fmt.Println("Done")
}
//为什么这两个不是交叉答打印的，因为执行的太快了



package main

import (
	"fmt"
	"time"
	"reflect"
)

func main()  {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	cur := time.Now()
	test,_:= time.ParseDuration("239h")
	test1 := cur.Add(test)
	sumM := cur.Sub(test1)
	fmt.Println(sumM)
	fmt.Println(reflect.TypeOf(sumM))
	fmt.Println(sumM.Hours()/24)

}

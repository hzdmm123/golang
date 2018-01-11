package main

import "fmt"

type A struct {
	Id int
	Name string `json:"name"`
} 

type B struct {
	A
	Name string `json:"name"`
}

func main()  {
	var a B
	a.Id = 1
	fmt.Print(a)
}

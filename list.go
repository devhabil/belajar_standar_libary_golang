package main

import (
	"container/list"
	"fmt"
)

func main() {
	// manual
	var data *list.List = list.New()

	data.PushBack("Muhammad")
	data.PushBack("Habil")
	data.PushBack("Arifin")

	var head *list.Element = data.Front()
	fmt.Println(head.Value)

	next := head.Next() 
	fmt.Println(next.Value)

	next = next.Next()
	fmt.Println(next.Value)

	//perulangan
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
package main

import (
	"fmt"
	"go-ds/doublylinkedlist"
)

func main() {
    list := doublylinkedlist.New()

    list.PushTail(1)
    list.PushTail(2)
    list.PushHead(3)

    list.Print()

    fmt.Println(list.PopHead())
    fmt.Println(list.PopTail())
    fmt.Println(list.PopTail())
    fmt.Println(list.PopTail())
}

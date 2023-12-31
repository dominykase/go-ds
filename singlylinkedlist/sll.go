package singlylinkedlist

import (
	"errors"
	"fmt"
)

type SinglyLinkedList struct {
    Head *Node
    Tail *Node
}

type Node struct {
    Data int
    Next *Node
}

func CreateSinglyLinkedList(values []int) (*SinglyLinkedList, error) {
    if len(values) == 0 {
        return nil, errors.New(
            "SinglyLinkedList::CreateSinglyLinkedList(values []int) must take a non-empty slice of ints 'values'",
        )
    }

    list := SinglyLinkedList{Head: nil, Tail: nil}
    list.Head = &Node{Data: values[0], Next: nil}
    list.Tail = list.Head

    for i, value := range values {
        if i == 0 {
            continue
        }

        list.Tail.Next = &Node{Data: value, Next: nil}
        list.Tail = list.Tail.Next
    }
    
    return &list, nil
}

func Get(list *SinglyLinkedList, id int) (int, error) {
    if id < 0 {
        return 0, errors.New(fmt.Sprintf("SinglyLinkedList::Get(id int), id=%d index out of bounds.", id))
    }

    current := list.Head
    i := 0

    for ; i < id; i++ {
        if (current.Next != nil) {
            current = current.Next
        } else {
            return 0, errors.New(fmt.Sprintf("SinglyLinkedList::Get(id int), id=%d index out of bounds.", id))
        }
    }

    return current.Data, nil
}

func Add(list *SinglyLinkedList, value int) {
    list.Tail.Next = &Node{Data: value, Next: nil}
    list.Tail = list.Tail.Next
}

func Length(list *SinglyLinkedList) int {
    length := 0
    current := list.Head

    for ; current.Next != nil; length++ {
        current = current.Next
    }

    length++

    return length
}

func PrintSinglyLinkedList(list *SinglyLinkedList) {
    current := list.Head

    for ; current.Next != nil; {
        fmt.Printf("%d ", current.Data)
        current = current.Next
    }

    fmt.Printf("%d\n", current.Data)
}

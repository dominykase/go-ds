package doublylinkedlist

import (
	"errors"
	"fmt"
)

type DoublyLinkedList struct {
    head *Node
    tail *Node
}

type Node struct {
    value int
    next *Node
    prev *Node
}

func New() *DoublyLinkedList {
    return &DoublyLinkedList{}
}

func (list *DoublyLinkedList) PushHead(value int) {
    if list.head == nil && list.tail == nil {
        list.head = &Node{value: value, next: nil, prev: nil}
        list.tail = list.head
        
        return
    }

    list.head = &Node{value: value, next: list.head, prev: nil}
}

func (list *DoublyLinkedList) PushTail(value int) {
    if list.head == nil && list.tail == nil {
        list.head = &Node{value: value, next: nil, prev: nil}
        list.tail = list.head
        
        return
    }

    list.tail.next = &Node{value: value, next: nil, prev: list.tail}
    list.tail = list.tail.next
}

func (list *DoublyLinkedList) PopHead() (int, error) {
    if list.head == nil && list.tail == nil {
        return 0, errors.New("Cannot pop from an empty list")
    }

    if list.head == list.tail {
        value := list.head.value
        list.head = nil
        list.tail = nil

        return value, nil
    }

    value := list.head.value
    list.head = list.head.next

    return value, nil
}

func (list *DoublyLinkedList) PopTail() (int, error) {
    if list.head == nil && list.tail == nil {
        return 0, errors.New("Cannot pop from an empty list")
    }

    if list.head == list.tail {
        value := list.tail.value
        list.head = nil
        list.tail = nil
        
        return value, nil
    }

    value := list.tail.value
    list.tail = list.tail.prev

    return value, nil
}

func (list *DoublyLinkedList) PeekHead() (int, error) {
    if list.head == nil && list.tail == nil {
        return 0, errors.New("Cannot peek from an empty list")
    }

    return list.head.value, nil
}

func (list *DoublyLinkedList) PeekTail() (int, error) {
    if list.head == nil && list.tail == nil {
        return 0, errors.New("Cannot peek from an empty list")
    }

    return list.tail.value, nil
}

func (list *DoublyLinkedList) IsEmpty() bool {
    return list.head == nil && list.tail == nil
}

func (list *DoublyLinkedList) Length() int {
    length := 0

    for node := list.head; node != nil; node = node.next {
        length++
    }

    return length
}

func (list *DoublyLinkedList) Get(index int) (int, error) {
    if index < 0 || index >= list.Length() {
        return 0, errors.New("Index out of bounds")
    }

    node := list.head

    for i := 0; i < index; i++ {
        node = node.next
    }

    return node.value, nil
}

func (list *DoublyLinkedList) Print() {
    for node := list.head; node != nil; node = node.next {
        fmt.Printf("%d ", node.value)
    }

    fmt.Printf("\n")
}

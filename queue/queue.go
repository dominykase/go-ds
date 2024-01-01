package queue

import (
	"errors"
	"fmt"
)

type Queue struct {
    front *Node
    back *Node
}

type Node struct {
    value int
    next *Node
    prev *Node
}

func New() *Queue {
    return &Queue{front: nil, back: nil}
}

func (queue *Queue) Push(value int) {
    if queue.IsEmpty() {
        queue.front = &Node{value: value, next: nil, prev:nil}
        queue.back = queue.front
    } else {
        queue.back.prev = &Node{value: value, next: queue.back, prev: nil}
        queue.back = queue.back.prev
    }
}

func (queue *Queue) Pop() (int, error) {
    val := 0

    var err error
    err = nil

    if queue.IsEmpty() {
        err = errors.New("Cannot pop from an empty queue.")
    } else if queue.Length() > 1 {
        val = queue.front.value

        queue.front = queue.front.prev
        queue.front.next = nil
    } else {
        val = queue.front.value

        queue.front = nil
        queue.back = nil
    }

    return val, err 
}

func (queue *Queue) IsEmpty() bool {
    return queue.front == nil && queue.back == nil 
}

func (queue *Queue) Length() int {
    if queue.IsEmpty() {
        return 0
    }

    current := queue.back

    i := 0

    for ; current.next != nil; i++ {
        current = current.next
    }

    i++
    
    return i
}

func (queue *Queue) ToString() string {
    if queue.IsEmpty() {
        return "The queue is empty"
    }

    current := queue.back

    s := ""

    for ; current.next != nil; {
        s += fmt.Sprintf("%d->", current.value)
        current = current.next
    }

    s += fmt.Sprintf("%d\n", current.value)

    return s
}

func (queue *Queue) Print() {
    if queue.IsEmpty() {
        fmt.Println("The queue is empty")
    }

    current := queue.back

    for ; current.next != nil; {
        fmt.Printf("%d->", current.value)
        current = current.next
    }

    fmt.Printf("%d\n", current.value)
}

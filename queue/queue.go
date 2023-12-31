package queue

import (
	"errors"
	"fmt"
)

type Queue struct {
    Front *Node
    Back *Node
}

type Node struct {
    Value int
    Next *Node
    Prev *Node
}

func CreateQueue() *Queue {
    return &Queue{Front: nil, Back: nil}
}

func Push(queue *Queue, value int) {
    if IsEmpty(queue) {
        queue.Front = &Node{Value: value, Next: nil, Prev:nil}
        queue.Back = queue.Front
    } else {
        queue.Back.Prev = &Node{Value: value, Next: queue.Back, Prev: nil}
        queue.Back = queue.Back.Prev
    }
}

func Pop(queue *Queue) (int, error) {
    val := 0

    var err error
    err = nil

    if IsEmpty(queue) {
        err = errors.New("Cannot pop from an empty queue.")
    } else if Length(queue) > 1 {
        val = queue.Front.Value

        queue.Front = queue.Front.Prev
        queue.Front.Next = nil
    } else {
        val = queue.Front.Value

        queue.Front = nil
        queue.Back = nil
    }

    return val, err 
}

func IsEmpty(queue *Queue) bool {
    return queue.Front == nil && queue.Back == nil 
}

func Length(queue *Queue) int {
    if IsEmpty(queue) {
        return 0
    }

    current := queue.Back

    i := 0

    for ; current.Next != nil; i++ {
        current = current.Next
    }

    i++
    
    return i
}

func PrintQueue(queue *Queue) {
    if IsEmpty(queue) {
        fmt.Println("The queue is empty")
    }

    current := queue.Back

    for ; current.Next != nil; {
        fmt.Printf("%d->", current.Value)
        current = current.Next
    }

    fmt.Printf("%d\n", current.Value)
}

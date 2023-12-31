package stack

import "fmt"

type Stack struct {
    Top *Node
}

type Node struct {
    Value int
    Next *Node
    Prev *Node
}

func CreateStack() *Stack {
    return &Stack{nil}
}

func Push(stack *Stack, value int) {
    if stack.Top == nil {
        stack.Top = &Node{Value: value, Next: nil, Prev: nil}
    } else {
        stack.Top.Next = &Node{Value: value, Next: nil, Prev: stack.Top}
        stack.Top = stack.Top.Next
    }
}

func Pop(stack *Stack) int {
    val := stack.Top.Value

    stack.Top = stack.Top.Prev
    
    if stack.Top != nil {
        stack.Top.Next = nil
    }

    return val
}

func Length(stack *Stack) int {
    if stack.Top == nil {
        return 0
    }

    i := 0
    current := stack.Top
    
    for ; current.Prev != nil; i++ {
        current = current.Prev
    }

    i++;

    return i
}

func IsEmpty(stack *Stack) bool {
    return stack.Top == nil
}

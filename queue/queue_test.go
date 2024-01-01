package queue

import "testing"

func TestPushToQueue(t *testing.T) {
    queue := New()

    queue.Push(1)
    queue.Push(2)
    queue.Push(3)

    if queue.Length() != 3 {
        t.Error("Queue length should be 3.")
    }

    actual := queue.ToString()

    if actual != "3->2->1\n" {
        t.Error("Queue should be 3->2->1.")
    }
}

func TestPopFromQueue(t *testing.T) {
    queue := New()

    queue.Push(1)
    queue.Push(2)

    actual, err := queue.Pop()

    if actual != 1 || err != nil {
        t.Errorf("Pop should return 1. Error: %v", err)
    }
}

func TestPopFromEmptyQueue(t *testing.T) {
    queue := New()

    _, err := queue.Pop()

    if err == nil {
        t.Error("Pop should return error.")
    }
}

func TestIsEmpty(t *testing.T) {
    queue := New()

    if !queue.IsEmpty() {
        t.Error("Queue should be empty #1.")
    }

    queue.Push(1)
    queue.Pop()

    if !queue.IsEmpty() {
        t.Error("Queue should be empty #2.")
    }
}

func TestLength(t *testing.T) {
    queue := New()

    if queue.Length() != 0 {
        t.Error("Queue length should be 0. #1")
    }

    queue.Push(1)

    if queue.Length() != 1 {
        t.Error("Queue length should be 1. #1")
    }

    queue.Push(2)

    if queue.Length() != 2 {
        t.Error("Queue length should be 2.")
    }

    queue.Pop()

    if queue.Length() != 1 {
        t.Error("Queue length should be 1. #2")
    }

    queue.Pop()

    if queue.Length() != 0 {
        t.Error("Queue length should be 0. #2")
    }
}

func TestToString(t *testing.T) {
    queue := New()

    if queue.ToString() != "The queue is empty" {
        t.Error("Queue should be empty.")
    }

    queue.Push(1)
    queue.Push(2)
    queue.Push(3)

    if queue.ToString() != "3->2->1\n" {
        t.Error("Queue should be 3->2->1.")
    }
}

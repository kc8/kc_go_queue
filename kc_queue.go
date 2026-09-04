package kc_queue

import (
	"bytes"
	"fmt"
    "sync"
)

// Implementation with be an array of nodes using circular buffer
// Insert at the tail, return at the head
type Queue[T any] struct { 
    lock sync.Mutex
    nodes []T
    head int 
    tail int
    length int // available slots
    size int // total size of queue
}

func New[T any]() *Queue[T] {
    result := &Queue[T]{}
    result.create()
    return result
}

func (q* Queue[T]) create() *Queue[T] {
    initalSize := 1
    q.nodes = make([]T, initalSize)
    q.head = 0
    q.tail = 0
    q.size = initalSize
    q.length = 0
    return q
}

func (q* Queue[T]) Deque() T {
    q.lock.Lock()
    defer q.lock.Unlock()
    if q.length == 0 {
        var result T
        return result // returns the default type for anytype
    }
    result := q.nodes[q.head]
    q.head = (q.head + 1) % q.size
    q.length -= 1 
    return result
}

func (q* Queue[T]) IsFull() bool {
    q.lock.Lock()
    defer q.lock.Unlock()
    if q.length == q.size {
        return true 
    }
    return false
}

func (q* Queue[T]) IsEmpty() bool {
    q.lock.Lock()
    defer q.lock.Unlock()
    if q.length == 0 {
        return true 
    }
    return false
}

func (q* Queue[T]) Enqueue(node T) {
    q.lock.Lock()
    defer q.lock.Unlock()
    q.resize()
    q.nodes[q.tail] = node;
    q.tail = (q.tail + 1) % q.size
    q.length += 1
}

// Return error or the new size of the queue
func (q* Queue[T]) resize() (int, error) {
    if (q.length == q.size) {
        newSize := q.size * 2
        tempNodes := make([]T, newSize)
        //TODO we could try built in copy w/ array splicing
        for j := 0; j < q.size; j++ {
            // var pos int = q.head + j % q.size
            // fmt.Printf("pos %d, j %d, head %d, size %d, newsize %d\n" ,pos, j, q.head, q.size, newSize)
            if (q.head == 0) {
                tempNodes[j] = q.nodes[q.head + j % q.size];
            } else {
                tempNodes[j] = q.nodes[q.head % q.size];
            }
        }
        q.size = newSize
        q.head = 0
        q.tail = q.length
        q.nodes = tempNodes;
        return newSize , nil
    }
    return 0, nil
		for j := 0; j < q.length; j++ {
			tempNodes[j] = q.nodes[(q.head+j)%q.size]
		}
}

func (q *Queue[T]) ToString() string {
    // TODO this does not actually display the order of the queue
    var result  bytes.Buffer
    result.WriteString("CONTENTS OF QUEUE ARE NOT IN QUEUE ORDER")
    for i, obj := range q.nodes {
        result.WriteString(fmt.Sprintf("pos %d | obj %v \n", i, obj))
    }
    return result.String()
}

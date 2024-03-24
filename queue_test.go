package kc_queue

import (
	"testing"
)

func TestEmptyQueueShouldNotFail(t *testing.T) {
    var q *Queue[int]= New[int]()

    nothing := q.Deque();
    if (nothing != 0) {
        t.Errorf("Failed with none nil value\n")
    }
}

func TestNewAndEnqueue(t *testing.T) {
    var q *Queue[int] = New[int]()
    q.Enqueue(1000)
    if q.IsEmpty() == true {
        t.Errorf("Queue should have at least one elemenet, got isEmpty == true")
    }

    if q.IsEmpty() == true {
        t.Errorf("Queue should have at least one elemenet, got isEmpty == true")
    }
}

func TestQueueEmpties(t *testing.T) {
    var q *Queue[int] = New[int]()
    q.Enqueue(0)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    q.Deque()
    q.Deque()
    q.Deque()
    q.Deque()
    q.Deque()

    nothing := q.Deque()
    if (nothing != 0) {
        t.Errorf("q.Deque got value with none nil value %v\n", nothing)
    }
}

func TestQueueAddRemove(t *testing.T) {
    var q *Queue[int] = New[int]()
    q.Enqueue(0)
    q.Enqueue(1)
    q.Enqueue(2)
    q.Enqueue(3)

    zero := q.Deque()
    if (zero != 0) {
        t.Errorf("q.Deque got %d instead of %d", zero, 0)
    }

    one := q.Deque()
    if (one != 1) {
        t.Errorf("q.Deque got %d instead of %d", one, 1)
    }
    q.Enqueue(4)
    q.Enqueue(5)

    two := q.Deque()
    if (two != 2) {
        t.Errorf("q.Deque got %d instead of %d", two, 2)
    }

    three := q.Deque()
    four := q.Deque()
    if (three != 3) {
        t.Errorf("q.Deque got %d instead of %d", three, 3)
    }

    if (four != 4) {
        t.Errorf("q.Deque got %d instead of %d", four, 4)
    }
}

func TestMesageStruct(t* testing.T) {
    var q *Queue[int]= New[int]()
    q.Enqueue(0)
    q.Enqueue(1)
    q.Enqueue(2)

    var zero = q.Deque()
    if (zero != 0) {
        t.Errorf("q.Deque got %d instead of %d", zero, 0)
    }

    q.Enqueue(3)
    // --------------
    q.Enqueue(4)
    q.Enqueue(5)
    q.Enqueue(6)
    var one = q.Deque()

    q.Enqueue(7)
    if (one != 1) {
        t.Errorf("q.Deque got %d instead of %d", one, 1)
    }
}

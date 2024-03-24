package kc_queue

import "testing"

func TestNewAndEnqueue_Prority(t *testing.T) {
	var q *PriorityQueue[int] = NewPriorityQueue[int]()
	q.ProrityInsert(1000, 100)
	if q.IsEmpty() == true {
		t.Errorf("Queue should have at least one element, got isEmpty == true")
	}
	if q.IsEmpty() == true {
		t.Errorf("Queue should have at least one element, got isEmpty == true")
	}
}

func TestAddItemAndGetBack_Prority(t *testing.T) {
	var q *PriorityQueue[int] = NewPriorityQueue[int]()
	q.ProrityInsert(1000, 100)
    item := q.GetNetxtPriorityItem()
    if item != 1000 {
		t.Errorf("Wanted to get back %d but got %d instead", 1000, item)
    }
}

func TestAddItemAndGetBackCorrect_Prority(t *testing.T) {
	var q *PriorityQueue[int] = NewPriorityQueue[int]()
	q.ProrityInsert(1000, 100)
	q.ProrityInsert(2000, 10)
	q.ProrityInsert(3000, 1000)
	q.ProrityInsert(4000, 0)
    item := q.GetNetxtPriorityItem()
    if item != 3000 {
		t.Errorf("Wanted to get back %d but got %d instead", 3000, item)
    }
    item = q.GetNetxtPriorityItem()
    if item != 1000 {
		t.Errorf("Wanted to get back %d but got %d instead", 1000, item)
    }
}

func TestAddItemAndGetBackCorrecAfterAppend_Prority(t *testing.T) {
	var q *PriorityQueue[int] = NewPriorityQueue[int]()
	q.ProrityInsert(1000, 100)
	q.ProrityInsert(2000, 10)
	q.ProrityInsert(3000, 1000)
	q.ProrityInsert(4000, 0)
    item := q.GetNetxtPriorityItem()
    if item != 3000 {
		t.Errorf("Wanted to get back %d but got %d instead", 3000, item)
    }
	q.ProrityInsert(9089, 1)
	q.ProrityInsert(9099, 9000)
    item = q.GetNetxtPriorityItem()
    if item != 9099 {
		t.Errorf("Wanted to get back %d but got %d instead", 9099, item)
    }
}

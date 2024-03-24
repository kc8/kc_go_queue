package kc_queue

import (
	"errors"
)

// NOTE this is a very naive implementation

type node[T any] struct {
	data    T
	prority int
}

type PriorityQueue[T any] struct {
	nodes        []node[T]
	size         int
	validEntries int
}

func (pq *PriorityQueue[T]) ProrityInsert(rawData T, prority int) error {
	if prority < 0 {
		return errors.New("[ERROR] priority for this queue must be above -1 (negative values reserved for naive solution implementation")
	}
	newNode := node[T]{
		data:    rawData,
		prority: prority,
	}

	pq.nodes = append(pq.nodes, newNode)
	pq.validEntries++
	return nil
}

func NewPriorityQueue[T any]() *PriorityQueue[T] {
	nodes := make([]node[T], 0)
	result := &PriorityQueue[T]{
		nodes: nodes,
	}
	return result
}

func (pq *PriorityQueue[T]) GetNetxtPriorityItem() T {
	var highestPriority = -1
	var highest T
	var highestIndex = -1
	for index, node := range pq.nodes {
		if node.prority > highestPriority {
			highest = node.data
            highestPriority = node.prority
			highestIndex = index
		}
	}
	if highestIndex < len(pq.nodes) {
		pq.nodes[highestIndex].prority = -100
	}
	pq.validEntries--
	return highest
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	if pq.validEntries <= 0 {
		return true
	}
	return false
}

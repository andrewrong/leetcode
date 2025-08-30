// leetcode/problems/lc_00347_top_k_frequent_elements/go/solution.go
package main

import "container/heap"

type NodeHeap []*Node

type Node struct {
	Counter int
	Value   int
}

func (n NodeHeap) Len() int {
	return len(n)
}

func (n NodeHeap) Less(i, j int) bool {
	return n[i].Counter < n[j].Counter
}

func (n NodeHeap) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n *NodeHeap) Push(x any) {
	*n = append(*n, x.(*Node))
}

func (h *NodeHeap) Pop() any {
	size := h.Len()
	x := (*h)[size-1]
	*h = (*h)[:size-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	// step 1 统计次数
	counters := make(map[int]*Node)
	for _, value := range nums {
		v, existed := counters[value]
		if existed {
			v.Counter += 1
		} else {
			counters[value] = &Node{
				Value:   value,
				Counter: 1,
			}
		}
	}
	minHeap := &NodeHeap{}
	heap.Init(minHeap)
	for _, v := range counters {
		if minHeap.Len() < k {
			heap.Push(minHeap, v)
			continue
		}

		min := heap.Pop(minHeap).(*Node)
		if min.Counter > v.Counter {
			heap.Push(minHeap, min)
		} else {
			heap.Push(minHeap, v)
		}
	}

	result := make([]int, k)
	for i := 0; i < k; i += 1 {
		result[i] = heap.Pop(minHeap).(*Node).Value
	}

	return result
}

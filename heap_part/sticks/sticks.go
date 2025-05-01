package main

import (
	"container/heap"
	"fmt"
	"studying/heap_part/types"
)

func countSticks(sticks []int) int {
	//
	minHeap := &types.MinHeap{}
	heap.Init(minHeap)

	for _, stick := range sticks {
		heap.Push(minHeap, stick)
	}
	fmt.Println(minHeap)

	cost := 0

	for minHeap.Len() > 1 {
		//fmt.Println(minHeap)
		stick1 := heap.Pop(minHeap).(int)
		stick2 := heap.Pop(minHeap).(int)

		//fmt.Println(stick1, stick2)
		combined := stick1 + stick2

		cost += combined

		heap.Push(minHeap, combined)

	}

	return cost
}

func main() {
	sticks := []int{1, 8, 2, 5}

	output := countSticks(sticks)

	fmt.Println(output)

}

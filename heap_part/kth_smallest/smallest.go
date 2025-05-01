package main

import (
	"container/heap"
	"fmt"
	"studying/heap_part/types"
)

func FindKthSmallest(arr []int, k int) []int {
	maxHeap := &types.MaxHeap{}
	heap.Init(maxHeap)

	index := 0
	for ; index < k; index++ {
		heap.Push(maxHeap, arr[index])
	}

	for ; index < len(arr); index++ {
		if arr[index] < (maxHeap.MinHeap)[0] {
			heap.Pop(maxHeap)
			heap.Push(maxHeap, arr[index])
		}
	}

	result := make([]int, k)

	for i := 0; i < k; i++ {
		result[i] = heap.Pop(maxHeap).(int)
	}

	return result
}

func main() {
	arr := []int{3, 1, 5, 12, 2, 11}
	output := FindKthSmallest(arr, 3)
	fmt.Println(output)

}

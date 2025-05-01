package main

import (
	"container/heap"
	"fmt"
	"studying/heap_part/types"
)

func findKLargest(nums []int, k int) []int {

	minHeap := &types.MinHeap{}
	heap.Init(minHeap)

	index := 0
	for ; index < k; index++ {
		heap.Push(minHeap, nums[index])
	}

	fmt.Println(minHeap)

	for ; index < len(nums); index++ {

		fmt.Println(nums[index], minHeap)

		if nums[index] > (*minHeap)[0] {
			//need to call
			heap.Pop(minHeap)
			heap.Push(minHeap, nums[index])
		}
	}

	//fmt.Println(minHeap)

	result := make([]int, k)
	for i := 0; i < len(result); i++ {
		result[i] = heap.Pop(minHeap).(int)
	}
	return result
}

func main() {
	arr := []int{3, 1, 5, 12, 2, 11}
	output := findKLargest(arr, 3)
	fmt.Println(output)

}

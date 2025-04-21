package main

import (
	"fmt"
	"studying/stacks"
)

//https://www.youtube.com/watch?v=vsfcpSoqJDc&list=PLMPSjUOg8tP2wSulRK6aai19kdFJelT-P&index=3

const NOT_VISITED int = 0
const VISITED int = 2

func main() {

	graph := map[int][]int{
		1: []int{2, 3, 4},
		2: []int{3, 5},
	}

	fmt.Println(topoOrder(graph))
}

func topoOrder(graph map[int][]int) []int {
	//find verticies
	stack := stacks.NewStack[int]()

	verts := make(map[int]struct{})
	for k, v := range graph {
		verts[k] = struct{}{}
		for _, j := range v {
			verts[j] = struct{}{}
		}
	}
	fmt.Println("Verts")
	for k, _ := range verts {
		fmt.Println("\t", k)
	}
	fmt.Println("end Verts")

	status := make(map[int]int)

	for k, _ := range verts {
		if status[k] == NOT_VISITED {
			fmt.Println("\t", k)
			fmt.Println("need to do something")
			dfs(graph, stack, status, k)
		}
	}

	fmt.Println(stack.Stack) //need to revers this

	result := []int{}
	//now to result...
	for stack.IsEmpty() == false {
		v, _ := stack.Pop()
		result = append(result, v)
	}

	return result
}

func dfs(graph map[int][]int, stack *stacks.Stack[int], status map[int]int, vert int) {
	status[vert] = VISITED
	fmt.Printf("Setting %d to VISITED\n", vert)

	for _, neighbor := range graph[vert] {
		if status[neighbor] == NOT_VISITED {
			dfs(graph, stack, status, neighbor)
		}
	}
	stack.Push(vert)
}

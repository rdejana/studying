package main

import (
	"fmt"
	"studying/binary_tree"
)

func main() {
	//let's build a simple tree
	root := binary_tree.NewTreeWithKey(100)
	root.Insert(50)
	root.Insert(150)
	root.Insert(69)
	root.Insert(96)
	bt := binary_tree.Tranverse(root)

	for _, v := range bt {
		fmt.Println(v)
	}

	found := root.Search(151)
	fmt.Println(found)

}

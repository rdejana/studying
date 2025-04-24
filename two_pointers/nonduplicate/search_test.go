package nonduplicate

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	input := []int{2, 3, 3, 3, 6, 9, 9}
	output := moveElements(input)
	fmt.Println(output)
}

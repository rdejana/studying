package dutch

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	input := []int{1, 0, 2, 1, 0}
	dutchFlag(input)

	fmt.Println(input)
}

func Test2(t *testing.T) {
	input := []int{2, 2, 0, 1, 2, 0}
	dutchFlag(input)

	fmt.Println(input)
}

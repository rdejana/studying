package window

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	input := []int{1, 3, 2, 0, -1, 7, 10}
	output := sort(input)
	fmt.Println(output)
}

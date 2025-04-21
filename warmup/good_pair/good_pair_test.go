package good_pair

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 1}
	pairs := NumberOfGoodPairs(nums)

	fmt.Println(pairs)

}

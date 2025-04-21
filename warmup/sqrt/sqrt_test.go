package sqrt

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {

	input := 4
	output := mySqrt(input)
	fmt.Println(input, output)
}

func Test2(t *testing.T) {

	input := 8
	output := mySqrt(input)
	fmt.Println(input, output)
}

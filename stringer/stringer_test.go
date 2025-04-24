package stringer

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	input := "xyz##"

	s := NewStack(input)
	for s.IsEmpty() == false {
		v, _ := s.PopValild()
		fmt.Println(v)
	}
}

func Test2(t *testing.T) {
	input := "xyz##"
	//input = "xy#z"
	input = "xywrrmu#p"
	output := myStringer(input)
	fmt.Println(input)
	fmt.Println(output)
}

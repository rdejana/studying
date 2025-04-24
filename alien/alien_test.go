package alien

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	input := []string{"kaa", "akcd", "akca", "cak", "cad"}
	output := alien(input)
	fmt.Println(output)
}

func Test2(t *testing.T) {
	input := []string{"b", "a"}
	output := alien(input)
	fmt.Println(output)
}

func Test3(t *testing.T) {
	input := []string{"ab", "a", "b"}
	output := alien(input)
	fmt.Println(output)
}

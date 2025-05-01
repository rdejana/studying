package intro

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	people := []int{10, 55, 70, 20, 90, 85}
	limit := 100
	total := numberOfBoats(people, limit)
	expected := 4
	fmt.Println("Number of Boats: ", total)
	if total != expected {
		t.Errorf("Number of Boats: %d, expected: %d", total, expected)
	}
}

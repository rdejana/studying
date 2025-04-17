package main

import (
	"fmt"
	"strings"
)

func hack(str1, str2 string) bool {
	//not optimized..
	str1R := processString(str1)
	str2R := processString(str2)

	return str1R == str2R

}

func processString(str1 string) string {
	str1Index := len(str1) - 1

	builder := strings.Builder{}

	for str1Index >= 0 {
		newIndex := findIndex(str1, str1Index)
		builder.WriteByte(str1[newIndex])
		str1Index = newIndex - 1 //??
	}

	//string is now backwards///
	return builder.String()
}

func findIndex(str string, index int) int {
	//how many back spaces?
	backspaces := 0
	for index >= 0 {
		if str[index] == '#' {
			//we have a backspace
			backspaces++
		} else if backspaces > 0 {
			backspaces--
		} else { //backspace == 0
			break //valid and no more back spaces
		}

		index--
	}

	return index
}

func main() {
	str1 := "xy#z"
	str2 := "xzz#"
	str1 = "xy#z"
	str2 = "xyz#"
	str1 = "xp#"
	str2 = "xyz##"

	str1 = "xywrrmp"
	str2 = "xywrrmu#p"

	r := hack(str1, str2)
	fmt.Println(r)
	//processString("xy#z")
}

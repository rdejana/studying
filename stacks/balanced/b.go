package main

import "fmt"

// we only have ()[]{}
func isValid(s1 string) bool {
	stack := []rune{}

	for _, c := range s1 {
		//fmt.Println(string(c))
		//opening...
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else {
			//this means closing...
			if len(stack) == 0 {
				//stack is empty and we have a closing...
				return false
			}
			//pop
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if c == ')' && top != '(' { //we got the wrong one back
				return false
			}

			if c == ']' && top != '[' { //we got the wrong one back
				return false
			}

			if c == '}' && top != '{' { //we got the wrong one back
				return false
			}

		}

	}

	return true
}

func main() {
	s1 := "{[()]}"
	valid := isValid(s1)

	fmt.Printf("Is %s valid: %t\n", s1, valid)

	s1 = "{[}]"
	valid = isValid(s1)

	fmt.Printf("Is %s valid: %t\n", s1, valid)

	s1 = "(]"
	valid = isValid(s1)

	fmt.Printf("Is %s valid: %t\n", s1, valid)

}

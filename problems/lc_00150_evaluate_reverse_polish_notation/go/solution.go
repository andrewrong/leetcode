// leetcode/problems/lc_00150_evaluate_reverse_polish_notation/go/solution.go
package main

import (
	kit "leetcode-go/kit/go"
	"strconv"
)

func evalRPN(tokens []string) int {
	// TODO: User implementation
	stack := kit.NewStack[string]()
	for _, token := range tokens {
		switch token {
		case "+", "*", "-", "/":
			{
				rightV, _ := stack.Pop()
				leftV, _ := stack.Pop()
				left, _ := strconv.ParseInt(leftV, 10, 32)
				right, _ := strconv.ParseInt(rightV, 10, 32)

				switch token {
				case "+":
					{
						stack.Push(strconv.FormatInt(left+right, 10))
					}
				case "*":
					{
						stack.Push(strconv.FormatInt(left*right, 10))
					}
				case "-":
					{
						stack.Push(strconv.FormatInt(left-right, 10))
					}
				case "/":
					{
						stack.Push(strconv.FormatInt(left/right, 10))
					}
				}
			}
		default:
			{
				stack.Push(token)
			}
		}
	}

	value, _ := stack.Pop()

	result, _ := strconv.ParseInt(value, 10, 32)
	return int(result)
}

// leetcode/problems/lc_00020_valid_parentheses/go/solution.go
package main

import kit "leetcode-go/kit/go"

func isValid(s string) bool {
	stack := kit.NewStack[byte]() // TODO: User implementation

	size := len(s)
	for i := 0; i < size; i += 1 {
		switch s[i] {
		case '(':
			{
				stack.Push(s[i])
			}
		case '{':
			{
				stack.Push(s[i])
			}
		case '[':
			{
				stack.Push(s[i])
			}
		case ')':
			{
				v, result := stack.Pop()
				if !result {
					return false
				}
				if v != '(' {
					return false
				}
			}
		case '}':
			{
				v, result := stack.Pop()
				if !result {
					return false
				}
				if v != '{' {
					return false
				}
			}
		case ']':
			{
				v, result := stack.Pop()
				if !result {
					return false
				}
				if v != '[' {
					return false
				}
			}
		default:
			{
				return false
			}
		}
	}
	if stack.Size() != 0 {
		return false
	}

	return true
}

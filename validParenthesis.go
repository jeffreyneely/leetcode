package main

func IsValidParenthesis(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := []rune{}
	m := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	for _, r := range s {
		switch r {
		case '(', '{', '[':
			stack = append(stack, r)
		case ')', '}', ']':
			l := len(stack) - 1
			if l+1 == 0 || stack[l] != m[r] {
				return false
			}
			stack = stack[:l]
		}
	}
	return len(stack) == 0
}

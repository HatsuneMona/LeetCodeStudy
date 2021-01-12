/**
 * @Author HatsuneMona
 * @Date  2020-11-13 15:26
 * @Description //TODO
 **/
package Q20

import "container/list"

func isValid(s string) bool {
	if s == "" {
		return true
	}
	strRune := []rune(s)
	stack := list.New()
	for _, r := range strRune {
		if r == '{' || r == '[' || r == '(' {
			stack.PushFront(r)
		} else {
			//防止出现错误：“}}}]])))”
			if stack.Len() == 0 {
				return false
			}
			need := stack.Front().Value.(rune)
			switch r {
			case '}':
				if need != '{' {
					return false
				}
			case ']':
				if need != '[' {
					return false
				}
			case ')':
				if need != '(' {
					return false
				}
			}
			stack.Remove(stack.Front())
		}
	}
	//防止出现错误：“{{((([”
	if stack.Len() == 0 {
		return true
	} else {
		return false
	}

}

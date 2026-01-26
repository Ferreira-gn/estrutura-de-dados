package balancingofexpression

import (
	"strings"

	"github.com/ferreira-gn/estrutura-de-dados/LeetCode/ValidParentheses/stack"
)

func IsValid(expression string) bool {
	leftExpression := stack.StackConstructor()
	rightExpression := stack.StackConstructor()

	mapLeftExpression := "({["
	mapRightExpression := ")}]"
	
	expressionCorresponding := map[rune]rune{
		'(' : ')',
		'{' : '}',
		'[' : ']',
	}

	for _, value := range expression {
		if strings.ContainsRune(mapLeftExpression, value) {
			leftExpression.Push(value)
			continue
		}

		if strings.ContainsRune(mapRightExpression, value) {
			rightExpression.Push(value)
			continue
		}
	}
	
	expressionLeftSize := leftExpression.Len()

	for i := 0 ; i < expressionLeftSize ; i++ {
		leftExpressionTop := leftExpression.ViewTheTop()
		rightExpressionTop := rightExpression.ViewTheTop()
		
		if leftExpression.Len() == 0 || rightExpression.Len() == 0 {
			return  false
		} 
		
		if  expressionCorresponding[leftExpressionTop] == rightExpressionTop {
			leftExpression.Pop()
			rightExpression.Pop()
			continue
		}
		
		return false		
	}

	
	return  true
}

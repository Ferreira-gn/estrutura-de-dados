package main

import (
	"fmt"

	balancingofexpression "github.com/ferreira-gn/estrutura-de-dados/LeetCode/ValidParentheses/balancing_of_expression"
)

func main() {
	var expression string = "[(]){}[(]){(})"
	fmt.Printf("%v", balancingofexpression.IsValid(expression))
}

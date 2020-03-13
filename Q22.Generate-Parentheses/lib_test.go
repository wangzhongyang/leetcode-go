package Q22_Generate_Parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  int
	output []string
}{
	{
		input: 3,
		output: []string{
			"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()"},
	},
}

func Test_mergeTwoLists(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := generateParenthesis(item.input)
		a.Equal(item.output, output)
	}
}

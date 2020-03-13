package Q20_Valid_Parentheses

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  string
	output bool
}{
	{
		input:  "()",
		output: true,
	},
	{
		input:  "()[]{}",
		output: true,
	},
	{
		input:  "(]",
		output: false,
	},
	{
		input:  "([)]",
		output: false,
	},
	{
		input:  "{[]}",
		output: true,
	},
}

func Test_isValid(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := isValid(item.input)
		if !a.Equal(item.output, output) {
			fmt.Println(item.input)
		}
	}
}

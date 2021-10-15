package Q151_Reverse_Words_in_a_String

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  string
	output string
}{
	{
		input:  "  hello world  ",
		output: "world hello",
	},
	{
		input:  "a good   example",
		output: "example good a",
	},
	{
		input:  "  Bob    Loves  Alice   ",
		output: "Alice Loves Bob",
	},
	{
		input:  "Alice does not even like bob",
		output: "bob like even not does Alice",
	},
}

func Test_reverseWords(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := reverseWords(item.input)
		a.Equal(item.output, output)
	}
}

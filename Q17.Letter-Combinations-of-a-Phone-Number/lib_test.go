package Q17_Letter_Combinations_of_a_Phone_Number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  string
	output []string
}{
	{
		input:  "23",
		output: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
	},
}

func Test_letterCombinations(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := letterCombinations(item.input)
		a.Equal(item.output, output)
	}
}

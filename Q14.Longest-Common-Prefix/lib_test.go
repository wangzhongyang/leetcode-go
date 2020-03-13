package Q14_Longest_Common_Prefix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  []string
	output string
}{
	{
		input:  []string{"flower", "flow", "flight"},
		output: "fl",
	},
	{
		input:  []string{"dog", "racecar", "car"},
		output: "",
	},
	{
		input:  []string{},
		output: "",
	},
}

func Test_longestCommonPrefix(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := longestCommonPrefix(item.input)
		a.Equal(item.output, output)
	}
}

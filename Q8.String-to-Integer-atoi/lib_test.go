package Q8_String_to_Integer_atoi

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  string
	output int
}{
	{
		input:  "42",
		output: 42,
	},
	{
		input:  "   -42",
		output: -42,
	},
	{
		input:  "4193 with words",
		output: 4193,
	},
	{
		input:  "words and 987",
		output: 0,
	},
	{
		input:  "-91283472332",
		output: -2147483648,
	},
}

func Test_myAtoi(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := myAtoi(item.input)
		fmt.Println(item.input, "|", output)
		a.Equal(item.output, output)
	}
}

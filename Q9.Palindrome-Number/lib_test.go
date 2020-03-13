package Q9_Palindrome_Number

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  int
	output bool
}{
	{
		input:  121,
		output: true,
	},
	{
		input:  -121,
		output: false,
	},
	{
		input:  10,
		output: false,
	},
}

func Test_isPalindrome(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := isPalindrome(item.input)
		fmt.Println(item.input, "|", output)
		a.Equal(item.output, output)
	}
}

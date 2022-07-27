package Q5_Longest_Palindromic_Substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	Input  string
	Output string
}{
	{
		Input:  "babad",
		Output: "bab",
	},
	{
		Input:  "cbbd",
		Output: "bb",
	},
	{
		Input:  "ccc",
		Output: "ccc",
	},
	{
		Input:  "abb",
		Output: "bb",
	},
	{
		Input:  "aacabdkacaa",
		Output: "aca",
	},
}

func Test_longestPalindrome(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := longestPalindrome(item.Input)
		a.Equal(item.Output, output)
	}
}

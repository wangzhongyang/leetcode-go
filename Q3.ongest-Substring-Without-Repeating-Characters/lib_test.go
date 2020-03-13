package Q3_ongest_Substring_Without_Repeating_Characters

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  string
	output int
	res    string
}{
	{
		input:  "abcabcbb",
		output: 3,
		res:    "abc",
	},
	{
		input:  "bbbbb",
		output: 1,
		res:    "b",
	},
	{
		input:  "pwwkew",
		output: 3,
		res:    "wke",
	},
	{
		input:  " ",
		output: 1,
		res:    " ",
	},
	{
		input:  "dvdf",
		output: 3,
		res:    "vdf",
	},
	{
		input:  "tmmzuxt",
		output: 5,
		res:    "mzuxt",
	},
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		fmt.Println("input:	", item.input)
		a.Equal(item.output, lengthOfLongestSubstring(item.input))
	}
	//fmt.Println(lengthOfLongestSubstring(" "))
	//fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
}

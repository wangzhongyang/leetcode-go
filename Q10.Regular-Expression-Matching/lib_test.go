package Q10_Regular_Expression_Matching

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	s      string
	p      string
	output bool
}{
	{
		s:      "aa",
		p:      "a",
		output: false,
	},
	{
		s:      "aa",
		p:      "a*",
		output: true,
	},
	{
		s:      "ab",
		p:      ".*",
		output: true,
	},
	{
		s:      "aab",
		p:      "c*a*b",
		output: true,
	},
	{
		s:      "mississippi",
		p:      "mis*is*p*.",
		output: false,
	},
}

func Test_isMatch(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := isMatch(item.s, item.p)
		fmt.Println(item.s, "|", item.p, "|", output)
		a.Equal(item.output, output)
	}
}

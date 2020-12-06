package Q567_Permutation_in_String

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	s1     string
	s2     string
	output bool
}{
	{
		s1:     "ab",
		s2:     "eidbaooo",
		output: true,
	},
	{
		s1:     "ab",
		s2:     "eidboaoo",
		output: false,
	},
	{
		s1:     "a",
		s2:     "ab",
		output: true,
	},
	{
		s1:     "ab",
		s2:     "ab",
		output: true,
	},
	{
		s1:     "hello",
		s2:     "ooolleoooleh",
		output: false,
	},
}

func Test_checkInclusion(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := checkInclusion(item.s1, item.s2)
		a.Equal(item.output, output)
	}
}

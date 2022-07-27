package Q62_Unique_Paths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	m      int
	n      int
	output int
}{
	{
		m:      3,
		n:      7,
		output: 28,
	},
	{
		m:      3,
		n:      2,
		output: 3,
	},
	{
		m:      3,
		n:      3,
		output: 6,
	},
}

func Test_uniquePaths(t *testing.T) {
	a := assert.New(t)
	for _, v := range tcs {
		res := uniquePaths(v.m, v.n)
		a.Equal(v.output, res)
	}
}

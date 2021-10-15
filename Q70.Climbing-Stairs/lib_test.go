package Q70_Climbing_Stairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  int
	output int
}{
	{
		input:  3,
		output: 3,
	},
}

func Test_climbStairs(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := climbStairs(item.input)
		a.Equal(item.output, output)
	}
}

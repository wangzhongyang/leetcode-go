package Q7_Reverse_Integer

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  int
	output int
}{
	{
		input:  123,
		output: 321,
	},
	{
		input:  -123,
		output: -321,
	},
	{
		input:  120,
		output: 21,
	},
	{
		input:  1534236469,
		output: 0,
	},
}

func Test_reverse(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := reverse(item.input)
		fmt.Println(item.input, output)
		a.Equal(item.output, output)
	}
	fmt.Println("max:", math.MaxInt32, 123/10)
}

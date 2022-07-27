package Q121_Best_Time_to_Buy_and_Sell_Stock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  []int
	output int
}{
	{
		input:  []int{7, 1, 5, 3, 6, 4},
		output: 5,
	},
}

func Test_maxProfit(t *testing.T) {
	a := assert.New(t)
	for _, v := range tcs {
		res := maxProfit(v.input)
		a.Equal(v.output, res)
	}
}

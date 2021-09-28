package Q42_Trapping_Rain_Water

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	nums []int
	res  int
}{
	{
		nums: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		res:  6,
	},
}

func Test_trap(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := trap(item.nums)
		a.Equal(output, item.res)
	}
}

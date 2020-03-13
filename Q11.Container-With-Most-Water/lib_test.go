package Q11_Container_With_Most_Water

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  []int
	output int
}{
	{input: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, output: 49},
}

func Test_maxArea(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := maxArea(item.input)
		a.Equal(item.output, output)
	}
}

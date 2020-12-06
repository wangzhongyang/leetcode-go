package Q415_Add_Strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	num1   string
	num2   string
	output string
}{
	{
		num1:   "1",
		num2:   "19",
		output: "20",
	},
	{
		num1:   "123",
		num2:   "234",
		output: "357",
	},
	{
		num1:   "1",
		num2:   "9",
		output: "10",
	},
}

func Test_addStrings(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := addStrings(item.num1, item.num2)
		a.Equal(item.output, output)
	}
}

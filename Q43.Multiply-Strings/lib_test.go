package Q43_Multiply_Strings

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
		output: "19",
	},
	{
		num1:   "123",
		num2:   "456",
		output: "56088",
	},
	{
		num1:   "0",
		num2:   "0",
		output: "0",
	},
	{
		num1:   "123456789",
		num2:   "987654321",
		output: "121932631112635269",
	},
	{
		num1:   "498828660196",
		num2:   "840477629533",
		output: "419254329864656431168468",
	},
	//{
	//	num1:   "215049099293599071909520290620044758659898177276367431",
	//	num2:   "80742066461429320157075733499567563769505940892",
	//	output: "0",
	//},
}

func Test_addStrings(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := multiply(item.num1, item.num2)
		a.Equal(item.output, output)
	}
}

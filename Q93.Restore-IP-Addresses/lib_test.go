package Q93_Restore_IP_Addresses

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	input  string
	output []string
}{
	{
		input: "25525511135",
		output: []string{
			"255.255.11.135",
			"255.255.111.35",
		},
	},
	{
		input: "0000",
		output: []string{
			"0.0.0.0",
		},
	},
	{
		input: "010010",
		output: []string{
			"0.10.0.10",
			"0.100.1.0",
		},
	},
}

func Test_restoreIpAddresses(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := restoreIpAddresses(item.input)
		fmt.Println(output, len(output), len(item.output))
		a.Equal(item.output, output)
	}
}

package Q79_Word_Search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	board  [][]byte
	word   string
	output bool
}{
	{
		board: [][]byte{
			[]byte{'A', 'B', 'C', 'E'},
			[]byte{'S', 'F', 'C', 'S'},
			[]byte{'A', 'D', 'E', 'E'},
		},
		word:   "ABCCED",
		output: true,
	},
	{
		board: [][]byte{
			[]byte{'A', 'B', 'C', 'E'},
			[]byte{'S', 'F', 'C', 'S'},
			[]byte{'A', 'D', 'E', 'E'},
		},
		word:   "SEE",
		output: true,
	},
	{
		board: [][]byte{
			[]byte{'A', 'B', 'C', 'E'},
			[]byte{'S', 'F', 'C', 'S'},
			[]byte{'A', 'D', 'E', 'E'},
		},
		word:   "SEE",
		output: true,
	},
	{
		board: [][]byte{
			[]byte{'A', 'B', 'C', 'E'},
			[]byte{'S', 'F', 'C', 'S'},
			[]byte{'A', 'D', 'E', 'E'},
		},
		word:   "ZZZZ",
		output: false,
	},
	{
		board: [][]byte{
			[]byte{'A', 'B', 'C', 'E'},
			[]byte{'S', 'F', 'E', 'S'},
			[]byte{'A', 'D', 'E', 'E'},
		},
		word:   "ABCESEEEFS",
		output: true,
	},
}

func Test_threeSum(t *testing.T) {
	a := assert.New(t)
	for _, item := range tcs {
		output := exist(item.board, item.word)
		a.Equal(item.output, output)
	}
}

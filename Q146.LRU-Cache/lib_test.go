package Q146_LRU_Cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	in1    []string
	in2    [][]int
	output []int
}{
	//{
	//	in1:    []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
	//	in2:    [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
	//	output: []int{0, 0, 0, 1, 0, -1, 0, -1, 3, 4},
	//},
	{
		in1:    []string{"LRUCache", "put", "put", "put", "put", "get", "get"},
		in2:    [][]int{{2}, {2, 1}, {1, 1}, {2, 3}, {4, 1}, {1}, {2}},
		output: []int{0, 0, 0, 0, 0, 1, -1},
	},
}

func Test_Constructor(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Println()
	lru.Put(2, 2)
	lru.Println()
	lru.Put(1, 3)
	lru.Println()
	//lru.Put(4, 4)
	//lru.Println()
	return
	a := assert.New(t)
	for _, item := range tcs {
		var lruCache LRUCache
		for k, v := range item.in1 {
			if k == 0 { // init LRUCache
				lruCache = Constructor(item.in2[0][0])
				continue
			}
			if v == "put" {
				lruCache.Put(item.in2[k][0], item.in2[k][1])
			} else {
				a.Equal(item.output[k], lruCache.Get(item.in2[k][0]))
			}
			lruCache.Println()
		}
	}
}

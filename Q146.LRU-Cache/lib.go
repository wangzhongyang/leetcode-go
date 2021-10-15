package Q146_LRU_Cache

import "fmt"

type LRUCache struct {
	len        int
	cap        int
	m          map[int]*List
	head, tail *List
}

type List struct {
	prev, next *List
	key, value int
}

// Constructor 在LeetCode上不稳定
//Runtime: 556 ms, faster than 84.32% of Go online submissions for LRU Cache.
//Memory Usage: 85.5 MB, less than 27.73% of Go online submissions for LRU Cache.
func Constructor(capacity int) LRUCache {
	head, tail := initList(0, 0), initList(0, 0)
	head.next = tail
	tail.prev = head
	return LRUCache{
		len:  capacity,
		cap:  0,
		m:    make(map[int]*List),
		head: head,
		tail: tail,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.m[key]
	if ok {
		// 重新拼接双链表
		node.prev.next = node.next
		node.next.prev = node.prev
		// 将该节点放到双链表的头部
		node.next = this.head.next
		this.head.next.prev = node
		node.prev = this.head
		this.head.next = node
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.len == 0 {
		return
	}
	node, ok := this.m[key]
	if ok {
		node.value = value
		// 重新拼接双链表
		node.prev.next = node.next
		node.next.prev = node.prev
	} else {
		// 如果已满，删除最后一个
		if this.cap == this.len {
			delete(this.m, this.tail.prev.key)
			last2 := this.tail.prev.prev // 倒数第二个
			last2.next = this.tail
			this.tail.prev = last2
		} else {
			this.cap++
		}
		node = initList(key, value)
		this.m[key] = node
	}
	// 将该节点放到双链表的头部
	secondNode := this.head.next
	this.head.next = node
	node.prev = this.head
	secondNode.prev = node
	node.next = secondNode
}

func initList(k, v int) *List {
	return &List{
		key:   k,
		value: v,
	}
}

func (this *LRUCache) Println() {
	head := this.head.next
	for head != nil && head != this.tail {
		fmt.Print("[", head.key, ",", head.value, "]")
		head = head.next
	}
	fmt.Println("")
}

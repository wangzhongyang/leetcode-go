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

func (this *LRUCache) Get(key int) int {
	item, ok := this.m[key]
	if !ok {
		return -1
	}
	// 拼接当前节点的前后节点
	this.rmItem(item)
	// 当前节点放到队首
	this.toHead(item)
	return item.value
}

func (this *LRUCache) Put(key int, value int) {
	item, ok := this.m[key]
	if ok {
		item.value = value
		// 拼接当前节点的前后节点
		this.rmItem(item)
		// 当前节点放到队首
		this.toHead(item)
	} else {
		item = &List{
			key:   key,
			value: value,
		}
		this.m[key] = item
		this.cap++

		// 当前节点放到队首
		this.toHead(item)
	}
	// 是否超过长度
	if this.cap > this.len {
		// 删除最后一个
		last := this.tail.prev
		delete(this.m, last.key)
		this.tail.prev = last.prev
		last.prev.next = this.tail
		this.cap--
	}
}

func (this *LRUCache) toHead(item *List) {
	headNext := this.head.next
	this.head.next = item
	item.next = headNext
	item.prev = this.head
	headNext.prev = item
}

func (this *LRUCache) rmItem(item *List) {
	item1, item2 := item.prev, item.next
	item1.next = item2
	item2.prev = item1
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

func initList(k, v int) *List {
	return &List{
		key:   k,
		value: v,
	}
}

func (this *LRUCache) Println() {
	fmt.Print("缓存是：")
	head := this.head.next
	for head != nil && head != this.tail {
		fmt.Print("[", head.key, ",", head.value, "]")
		head = head.next
	}
	fmt.Println("")
}

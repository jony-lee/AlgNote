package leetcode

/*
标题 : LRU Cache
链接 : https://leetcode.cn/problems/lru-cache/
难度 : 一般
解题思路:涉及操作缓存O1查找,O1移动,O1删除
O1查找使用字典,O1移动使用链表

注意事项:

*/
type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}
type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     &DLinkedNode{key: 0, value: 0},
		tail:     &DLinkedNode{key: 0, value: 0},
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if !ok {
		init_node := &DLinkedNode{key: key, value: value}
		this.cache[key] = init_node
		this.addToHead(init_node)
		this.size++
		if this.size > this.capacity {
			pop := this.removeTail()
			delete(this.cache, pop.key)
			this.size--
		}
	} else {
		node.value = value
		this.moveToHead(node)
	}
}

// 移动需要移除并添加
func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

// 淘汰缓存,需要让尾部节点前指，然后移除节点，需要注意不移除节点会导致内存泄漏
func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

// 双链表把节点添加到头部
func (this *LRUCache) addToHead(node *DLinkedNode) {
	// head和tail 是一个哨兵节点
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

// 双链表移除节点，让当前节点上下游互指
func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
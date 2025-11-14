package lrucache

type LRUCache struct {
	capacity int
	head     *Node
	tail     *Node
	items    map[int]*Node
}

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		items:    make(map[int]*Node, capacity),
		head:     head,
		tail:     tail,
	}
}

func (c *LRUCache) Get(i int) int {
	if node, ok := c.items[i]; ok {
		c.moveToHead(node)
		return node.value
	}

	return -1
}

func (c *LRUCache) Put(i int, value int) int {
	if node, ok := c.items[i]; ok {
		node.value = value
		c.moveToHead(node)
		return node.key
	}

	if len(c.items) == c.capacity {
		c.evitLRUNode()
	}

	node := &Node{
		key:   i,
		value: value,
	}
	c.items[i] = node
	c.insertHead(node)
	return node.value
}

func (c *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) moveToHead(node *Node) {
	c.remove(node)
	c.insertHead(node)
}

func (c *LRUCache) insertHead(node *Node) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

func (c *LRUCache) evitLRUNode() {
	lruNode := c.tail.prev
	if lruNode == c.head {
		return
	}

	c.remove(lruNode)
	delete(c.items, lruNode.key)

}

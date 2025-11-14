package implementqueueusingstacks

type MyQueue struct {
	waitQueue  []int
	serveQueue []int
}

func Constructor() MyQueue {
	return MyQueue{
		waitQueue:  []int{},
		serveQueue: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.waitQueue = append(this.waitQueue, x)
}

func (this *MyQueue) Pop() int {
	if len(this.serveQueue) == 0 {
		this.fillServeQueue()
	}
	n := this.serveQueue[len(this.serveQueue)-1]
	this.serveQueue = this.serveQueue[:len(this.serveQueue)-1]
	return n
}

func (this *MyQueue) Peek() int {
	if len(this.serveQueue) == 0 {
		this.fillServeQueue()
	}
	return this.serveQueue[len(this.serveQueue)-1]

}

func (this *MyQueue) Empty() bool {
	return len(this.waitQueue) == 0 && len(this.serveQueue) == 0
}

func (this *MyQueue) fillServeQueue() {
	for len(this.waitQueue) > 0 {
		n := this.waitQueue[len(this.waitQueue)-1]
		this.waitQueue = this.waitQueue[:len(this.waitQueue)-1]
		this.serveQueue = append(this.serveQueue, n)
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

package examples

import "fmt"

type MyCircularQueue struct {
	front    int
	rear     int
	capacity int
	data     []interface{}
}

/** Initialize your data structure here. Set the size of the queue to be k. */
// front指向队头元素, rear指向队尾元素的下一个位置,所以会存在一个空余的位置留给rear,所以我们实际make的数组长度为k+1
func Constructor(k int) MyCircularQueue {
	myQueue := MyCircularQueue{
		capacity: k + 1, //这里要注意capacity的长度是k+1,因为我们要留一个位置给rear,rear指向队尾的下一个元素,里面是不存储实际内容的,所以我们要存储数据为k时,我们的数组长度需要为k+1
		data:     make([]interface{}, k+1),
	}

	return myQueue
}

func (this MyCircularQueue) Length() int {
	length := (this.rear - this.front + this.capacity) % (this.capacity)
	return length
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		fmt.Println("myqueue is full!")
		return false
	}

	if this.IsEmpty() {
		this.front = 0
		this.rear = 1
		this.data[0] = value
		return true
	}

	result := this.push(value)
	return result
}

func (this *MyCircularQueue) push(value int) bool {
	this.data[this.rear] = value
	this.moveRear()
	return true
}

func (this *MyCircularQueue) moveRear() {
	this.rear = (this.rear + 1) % this.capacity
}

func (this *MyCircularQueue) moveFront() {
	this.front = (this.front + 1) % this.capacity
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		fmt.Println("the Queue is empty!")
		return false
	}
	this.data[this.front] = nil
	this.moveFront()
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		fmt.Println("the Queue is empty!")
		return -1
	}
	return this.data[this.front].(int)
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		fmt.Println("the Queue is empty!")
		return -1
	}
	return this.data[(this.rear+this.capacity-1)%this.capacity].(int)
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.front == this.rear
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if (this.rear+1)%this.capacity == this.front {
		return true
	}
	return false
}

func (this *MyCircularQueue) printq() {
	fmt.Printf("front:%v,rear:%v:len:%v,capacity:%v,data=%v\n", this.front, this.rear, this.Length(), this.capacity, this.data)
}

/**
* Your MyCircularQueue object will be instantiated and called as such:
* obj := Constructor(k);
* param_1 := obj.EnQueue(value);
* param_2 := obj.DeQueue();
* param_3 := obj.Front();
* param_4 := obj.Rear();
* param_5 := obj.IsEmpty();
* param_6 := obj.IsFull();
 */

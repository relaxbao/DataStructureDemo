package examples

import (
	"fmt"
)

func ExampleMyCircularQueue() {
	circularQueue := Constructor(3) // set the size to be 3
	p1 := circularQueue.EnQueue(1)  // return true
	p2 := circularQueue.EnQueue(2)  // return true
	p3 := circularQueue.EnQueue(3)  // return true
	p4 := circularQueue.EnQueue(4)  // return false, the queue is full
	p5 := circularQueue.Rear()      // return 3
	p6 := circularQueue.IsFull()    // return true
	p7 := circularQueue.DeQueue()   // return true
	p8 := circularQueue.EnQueue(4)  // return true
	p9 := circularQueue.Rear()      // return 4
	fmt.Println(p1, ",", p2, ",", p3, ",", p4, ",", p5, ",", p6, ",", p7, ",", p8, ",", p1, ",", p9)
	//output:
	// myqueue is full!
	// true , true , true , false , 3 , true , true , true , true , 4
}

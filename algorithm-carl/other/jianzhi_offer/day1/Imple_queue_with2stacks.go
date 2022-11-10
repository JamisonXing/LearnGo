package day1

type CQueue struct {
	instack, outstack []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (q *CQueue) AppendTail(value int) {
	q.instack = append(q.instack, value)
}

func (q *CQueue) DeleteHead() int {
	if len(q.outstack) == 0 {
		if len(q.instack) == 0 {
			return -1
		}
		//instack不为0
		q.in2out()
	}
	value := q.outstack[len(q.outstack)-1]
	q.outstack = q.outstack[:len(q.outstack)-1]
	return value
}

func (q *CQueue) in2out() {
	for len(q.instack) > 0 {
		q.outstack = append(q.outstack, q.instack[len(q.instack)-1])
		q.instack = q.instack[:len(q.instack)-1]
	}
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

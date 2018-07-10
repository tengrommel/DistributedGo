package implement_stack_using_queues

type MyStack struct {
	a, b *Queue
}

func Constructor() MyStack {
	return MyStack{a: NewQueue(), b: NewQueue()}
}

func (ms *MyStack) Push(x int) {
	if ms.a.Len() == 0 {
		ms.a, ms.b = ms.b, ms.a
	}
	ms.a.Push(x)
}

func (ms *MyStack) Pop() int {
	if ms.a.Len() == 0 {
		ms.a, ms.b = ms.b, ms.a
	}

	for ms.a.Len() > 1 {
		ms.b.Push(ms.a.Pop())
	}

	return ms.a.Pop()
}

func (ms *MyStack) Top() int {
	res := ms.Pop()
	ms.Push(res)
	return res
}

func (ms *MyStack) Empty() bool {
	return (ms.a.Len() + ms.b.Len()) == 0
}

type Queue struct {
	nums []int
}

func NewQueue() *Queue {
	return &Queue{nums: []int{}}
}

func (q *Queue)Push(n int)  {
	q.nums = append(q.nums, n)
}

func (q *Queue)Pop() int {
	res := q.nums[0]
	q.nums = q.nums[1:]
	return res
}

func (q *Queue)Len() int {
	return len(q.nums)
}

func (q *Queue)IsEmpty() bool {
	return q.Len() == 0
}
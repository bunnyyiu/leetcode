type Node struct {
    Val int
    MinVal int
    Next *Node
    Pre *Node
}

type MinStack struct {
    cur *Node
}

func min(a, b int) int {
    if (a < b) {
        return a
    }
    return b
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(x int)  {
    if this.cur == nil {
        this.cur = &Node{
            Val: x,
            MinVal: x,
            Next: nil,
            Pre: nil,
        }
    } else {
        newNode := &Node{
            Val: x,
            MinVal: min(x, this.cur.MinVal),
            Next:nil,
            Pre: this.cur,
        }
        this.cur.Next = newNode
        this.cur = this.cur.Next
    }
}


func (this *MinStack) Pop()  {
    if this.cur != nil {
        tmp := this.cur.Pre
        this.cur.Next = nil
        this.cur = tmp
    }
}


func (this *MinStack) Top() int {
    return this.cur.Val
}


func (this *MinStack) GetMin() int {
    return this.cur.MinVal
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
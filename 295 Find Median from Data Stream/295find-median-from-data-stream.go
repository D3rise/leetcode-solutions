type MedianFinder struct {
    LeftHeap *MaxHeap
    RightHeap *MinHeap
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Peek() int         { return (*h)[0] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Peek() int         { return (*h)[0] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Constructor() MedianFinder {
    rightHeap := &MinHeap{}
    leftHeap := &MaxHeap{}
    heap.Init(rightHeap)
    heap.Init(leftHeap)

    return MedianFinder{
        RightHeap: rightHeap,
        LeftHeap: leftHeap,
    }
}


func (this *MedianFinder) AddNum(num int)  {
    heap.Push(this.LeftHeap, num)
    heap.Push(this.RightHeap, heap.Pop(this.LeftHeap))

    if len(*this.RightHeap) > len(*this.LeftHeap) {
        heap.Push(this.LeftHeap, heap.Pop(this.RightHeap))
    }
}


func (this *MedianFinder) FindMedian() float64 {
    curLen := len(*this.RightHeap) + len(*this.LeftHeap)

    if curLen % 2 == 0 {
        mean := float64((*this.LeftHeap)[0] + (*this.RightHeap)[0]) / 2.0
        return mean
    }

    return float64((*this.LeftHeap)[0])
}


/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


func removeValueFromHeap(h *IntHeap, val int) {
    for i, num := range *h {
        if num == val {
            heap.Remove(h, i)
            return
        }
    }
}

func maxSlidingWindow(nums []int, k int) []int {
    begin := 0
    maxH := &IntHeap{}
    heap.Init(maxH)
    var result []int

    for end, num := range nums {
        heap.Push(maxH, num)

        if end - begin + 1 == k {
            result = append(result, (*maxH)[0])
            removeValueFromHeap(maxH, nums[begin])
            begin++
        }
    }

    return result
}
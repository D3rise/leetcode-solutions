import "container/heap"

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

func findAndRemoveFromMinHeap(h *MinHeap, elem int) bool {
    for i, num := range *h {
        if num == elem {
            heap.Remove(h, i)
            return true
        }
    }
    return false
}

func findAndRemoveFromMaxHeap(h *MaxHeap, elem int) bool {
    for i, num := range *h {
        if num == elem {
            heap.Remove(h, i)
            return true
        }
    }
    return false
}

func medianSlidingWindow(nums []int, k int) []float64 {
	leftHeap := &MaxHeap{}
    rightHeap := &MinHeap{}
    heap.Init(leftHeap)
    heap.Init(rightHeap)

    var result []float64

    begin := 0
    for end, num := range nums {
        heap.Push(leftHeap, num)
        heap.Push(rightHeap, heap.Pop(leftHeap))

        if len(*rightHeap) > len(*leftHeap) {
            heap.Push(leftHeap, heap.Pop(rightHeap))
        }

        if end - begin + 1 == k {
            if k % 2 == 0 {
                mean := float64((*leftHeap)[0] + (*rightHeap)[0]) / 2.0
                result = append(result, mean)
            } else {
                result = append(result, float64((*leftHeap)[0]))
            }

            if !findAndRemoveFromMinHeap(rightHeap, nums[begin]) {
                findAndRemoveFromMaxHeap(leftHeap, nums[begin])
            }

            if len(*rightHeap) > len(*leftHeap) {
                heap.Push(leftHeap, heap.Pop(rightHeap))
            }

            begin++
        }
    }

    return result
}
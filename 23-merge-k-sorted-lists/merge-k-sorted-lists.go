/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {	*h = append(*h, x.(int)) }

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
    h := &IntHeap{}
    heap.Init(h)
    dummy := &ListNode{}

    for _, head := range lists {
        current := head
        for current != nil {
            heap.Push(h, current.Val)
            current = current.Next
        }
    }

    current := dummy
    for h.Len() > 0 {
        newNode := &ListNode{
            Val: heap.Pop(h).(int),
        }

        current.Next = newNode
        current = current.Next
    }

    return dummy.Next
}
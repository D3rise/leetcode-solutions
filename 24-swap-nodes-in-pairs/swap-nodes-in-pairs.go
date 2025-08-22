/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    dummy := &ListNode{
        Next: head,
    }

    current := dummy
    for current.Next != nil && current.Next.Next != nil {
        first := current.Next
        second := current.Next.Next

        first.Next = second.Next
        second.Next = first

        current.Next = second
        current = first
    }

    return dummy.Next
}
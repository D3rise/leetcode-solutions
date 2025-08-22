/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    current := head
    for current != nil {
        next := current.Next
        if next != nil && next.Val == current.Val {
            for next != nil && next.Val == current.Val {
                next = next.Next
            }
        }
        current.Next = next
        current = next
    }

    return head
}
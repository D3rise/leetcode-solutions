/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }

    fast := head
    slow := head
    for k != 0 {
        if fast == nil {
            fast = head
        }
        fast = fast.Next
        k--
    }

    if fast == nil {
        return head
    }

    for fast.Next != nil {
        fast = fast.Next
        slow = slow.Next
    }

    fast.Next = head
    head = slow.Next
    slow.Next = nil

    return head
}
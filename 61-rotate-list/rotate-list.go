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

    // if fast is nil by the end of the cycle,
    // then the k is equal to the length of the
    // linked list and there is nothing to rotate
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
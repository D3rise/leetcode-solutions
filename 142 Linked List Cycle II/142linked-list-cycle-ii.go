func detectCycle(head *ListNode) *ListNode {
    fast := head
    slow := head

    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next

        if fast == slow {
            break
        }
    }

    if fast == nil || fast.Next == nil {
        return nil
    }

    fast = head
    for fast != slow {
        fast = fast.Next
        slow = slow.Next
    }

    return fast
}
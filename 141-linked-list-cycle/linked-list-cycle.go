func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }

    p1, p2 := head, head.Next
    for p2.Next != nil && p2.Next.Next != nil {
        if p1 == p2 {
            return true
        }

        p1 = p1.Next
        p2 = p2.Next.Next
    }

    return false
}
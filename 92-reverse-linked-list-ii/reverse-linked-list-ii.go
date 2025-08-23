func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if head == nil || left == right {
        return head
    }

    dummy := &ListNode{
        Next: head,
    }

    prev := dummy
    for i := 0; i < left - 1; i++ {
        prev = prev.Next
    }

    start := prev.Next
    current := prev.Next
    var prevHead *ListNode
    for i := 0; i <= right - left; i++ {
        next := current.Next
        current.Next = prevHead
        prevHead = current
        current = next
    }

    start.Next = current
    prev.Next = prevHead

    return dummy.Next
}
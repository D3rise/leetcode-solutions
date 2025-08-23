func reverseList(head *ListNode) *ListNode {
    current := head
    var prevHead *ListNode

    for current != nil {
        next := current.Next
        current.Next = prevHead
        prevHead = current
        current = next
    }

    return prevHead
}
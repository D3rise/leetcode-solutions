func reverseList(head *ListNode) *ListNode {
    var prevNode *ListNode
    current := head

    for current != nil {
        oldNext := current.Next
        current.Next = prevNode
        prevNode = current
        current = oldNext
    }

    return prevNode
}
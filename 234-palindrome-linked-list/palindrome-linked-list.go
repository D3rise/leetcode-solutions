func isPalindrome(head *ListNode) bool {
    slow, fast := head, head

    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        fmt.Println(slow)
    }

    var prevNode *ListNode
    current := slow
    for current != nil {
        oldNext := current.Next
        current.Next = prevNode
        prevNode = current
        current = oldNext
    }

    p1, p2 := head, prevNode
    for p1 != nil && p2 != nil {
        fmt.Println(p1.Val, p2.Val)
        if p1.Val != p2.Val {
            return false
        }

        p1 = p1.Next
        p2 = p2.Next
    }

    return true
}
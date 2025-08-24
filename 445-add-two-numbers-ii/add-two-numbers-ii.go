func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    stack1 := []int{}
    stack2 := []int{}

    p1 := l1
    p2 := l2
    for p1 != nil || p2 != nil {
        if p1 != nil { 
            stack1 = append(stack1, p1.Val)
            p1 = p1.Next 
        }

        if p2 != nil {
            stack2 = append(stack2, p2.Val)
            p2 = p2.Next
        }
    }

    var prevHead *ListNode
    remainder := 0
    for len(stack1) > 0 || len(stack2) > 0 || remainder > 0 {
        x := 0
        if len(stack1) > 0 {
            x = stack1[len(stack1)-1]
            stack1 = stack1[:len(stack1)-1]
        }

        y := 0
        if len(stack2) > 0 {
            y = stack2[len(stack2)-1]
            stack2 = stack2[:len(stack2)-1]
        }

        sum := x + y + remainder
        remainder = sum / 10
        sum = sum % 10

        current := &ListNode{
            Val: sum,
            Next: prevHead,
        }
        prevHead = current
    }

    return prevHead
}
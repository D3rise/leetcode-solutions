/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    p1 := l1
    p2 := l2

    dummy := &ListNode{}
    remainder := 0

    current := dummy
    for p1 != nil || p2 != nil || remainder > 0 {
        x, y := 0, 0
        if p1 != nil { x = p1.Val }
        if p2 != nil { y = p2.Val }

        sum := x + y + remainder
        remainder = sum / 10
        current.Next = &ListNode{
            Val: sum % 10,
        }
        current = current.Next
        
        if p1 != nil { p1 = p1.Next }
        if p2 != nil { p2 = p2.Next }
    }

    return dummy.Next
}
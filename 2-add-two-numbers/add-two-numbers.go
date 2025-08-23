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
    var remainder bool

    current := dummy
    for p1 != nil && p2 != nil {
        sum := p1.Val + p2.Val
        if remainder {
            sum++
            remainder = false
        }

        if sum >= 10 {
            sum = sum % 10
            remainder = true
        }

        current.Next = &ListNode{
            Val: sum,
        }
        current = current.Next
        p1 = p1.Next
        p2 = p2.Next
    }

    for p1 != nil {
        val := p1.Val
        if remainder {
            val++
            remainder = false
        }

        if val >= 10 {
            val = val % 10
            remainder = true
        }

        current.Next = &ListNode{
            Val: val,
        }
        current = current.Next
        p1 = p1.Next
    }

    for p2 != nil {
        val := p2.Val
        if remainder {
            val++
            remainder = false
        }

        if val >= 10 {
            val = val % 10
            remainder = true
        }

        current.Next = &ListNode{
            Val: val,
        }
        current = current.Next
        p2 = p2.Next
    }

    if remainder {
        current.Next = &ListNode{
            Val: 1,
        }
    }

    return dummy.Next
}
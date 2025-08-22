/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    p1 := list1
    p2 := list2

    current := dummy
    for p1 != nil && p2 != nil {
        if p1.Val < p2.Val {
            current.Next = p1
            p1 = p1.Next
        } else {
            current.Next = p2
            p2 = p2.Next
        }

        current = current.Next
    }

    if p1 != nil {
        current.Next = p1
    }

    if p2 != nil {
        current.Next = p2
    }

    return dummy.Next
}
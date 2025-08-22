/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{
        Next: head,
    }

    current := dummy
    for current.Next != nil && current.Next.Next != nil {
        // dummy -> 1 -> 2 -> 3
        //          |    |
        //        first  second
        first := current.Next
        second := current.Next.Next

        // now:    2 -> 1 -> 3
        //              ^
        //      dummy is still pointing to 1
        first.Next = second.Next
        second.Next = first

        // dummy -> 2 -> 1 -> 3
        current.Next = second
        current = first
    }

    return dummy.Next
}
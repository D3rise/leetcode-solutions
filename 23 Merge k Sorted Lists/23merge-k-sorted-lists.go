/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func push(headNode *ListNode, val int) *ListNode {
	if headNode == nil {
		return &ListNode{
			Val:  val,
			Next: nil,
		}
	}

	if val < headNode.Val {
		temp := headNode
		newNode := ListNode{
			Val:  val,
			Next: temp,
		}
		return &newNode
	}

	temp := headNode
	for temp.Next != nil {
		if temp.Next.Val >= val {
			break
		}
		temp = temp.Next
	}

	newNode := ListNode{
		Val:  val,
		Next: temp.Next,
	}
	temp.Next = &newNode

	return headNode
}

func mergeKLists(lists []*ListNode) *ListNode {
    var newList *ListNode

    for _, l := range lists {
        h := l
        for h != nil {
            newList = push(newList, h.Val)
            h = h.Next
        }
    }

    return newList
}
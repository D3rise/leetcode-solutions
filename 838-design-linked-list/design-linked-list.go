type MyLinkedList struct {
    head *ListNode
}


func Constructor() MyLinkedList {
    return MyLinkedList{}
}


func (this *MyLinkedList) Get(index int) int {
    if index > 0 && this.head == nil {
        return -1
    }

    current := this.head
    for i := 0; i < index; i++ {
        if current == nil {
            return -1
        }

        current = current.Next
    }

    if current == nil {
        return -1
    }

    return current.Val
}


func (this *MyLinkedList) AddAtHead(val int)  {
    newNode := ListNode{
        Val: val,
        Next: this.head,
    }

    this.head = &newNode
}


func (this *MyLinkedList) AddAtTail(val int)  {
    current := this.head
    newNode := ListNode{
        Val: val,
    }

    if current == nil {
        this.head = &newNode
        return    
    }

    for current.Next != nil {
        current = current.Next
    }

    current.Next = &newNode
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    newNode := ListNode{
        Val: val,
    }

    if index == 0 {
        newNode.Next = this.head
        this.head = &newNode
        return
    }

    if this.head == nil {
        return
    }

    current := this.head
    for i := 0; i < index-1; i++ {
        if current == nil {
            return
        }
        current = current.Next
    }

    if current == nil {
        return
    }

    if current.Next != nil {
        newNode.Next = current.Next
    }

    current.Next = &newNode
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if this.head == nil {
        return
    }

    if index == 0 {
        this.head = this.head.Next
    }

    current := this.head
    for i := 0; i < index-1; i++ {
        if current == nil {
            return
        }

        current = current.Next
    }

    if current == nil || current.Next == nil {
        return
    }

    current.Next = current.Next.Next
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
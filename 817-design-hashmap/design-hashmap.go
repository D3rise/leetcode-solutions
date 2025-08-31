type Node struct {
    Key int
    Value int
    Next *Node
}

type LinkedList struct {
    Head *Node
}

func (l *LinkedList) Put(key int, val int) {
    current := l.Head
    for current != nil {
        if current.Key == key {
            current.Value = val
            return
        }
        current = current.Next
    }

    newHead := &Node{
        Key: key,
        Value: val,
        Next: l.Head,
    }

    l.Head = newHead
}

func (l LinkedList) Get(key int) int {
    current := l.Head

    for current != nil {
        if current.Key == key {
            return current.Value
        }
        current = current.Next
    }

    return -1
}

func (l *LinkedList) Remove(key int) {
    if (l.Head == nil) {
        return
    }

    if (l.Head.Key == key) {
        l.Head = l.Head.Next
        return
    }

    current := l.Head
    for current.Next != nil {
        if current.Next.Key == key {
            current.Next = current.Next.Next
            return
        }
        current = current.Next
    }
}

type MyHashMap struct {
    n int
    buckets []*LinkedList
}

func Constructor() MyHashMap {
    n := 991
    buckets := make([]*LinkedList, n)

    return MyHashMap{
        n: n,
        buckets: buckets,
    }
}

func (this *MyHashMap) hash(key int) int {
    return key % this.n
}

func (this *MyHashMap) Put(key int, value int)  {
    n := this.hash(key)
    if ok := this.buckets[n]; ok == nil {
        this.buckets[n] = &LinkedList{}
    }
    this.buckets[n].Put(key, value)
}


func (this *MyHashMap) Get(key int) int {
    n := this.hash(key)
    if ok := this.buckets[n]; ok == nil {
        this.buckets[n] = &LinkedList{}
    }
    return this.buckets[n].Get(key)
}

func (this *MyHashMap) Remove(key int)  {
    n := this.hash(key)
    if ok := this.buckets[n]; ok == nil {
        this.buckets[n] = &LinkedList{}
    }
    this.buckets[n].Remove(key)
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
package lru

import "fmt"

type Node struct {
        Key   interface{}
        Value interface{}
        Prev  *Node
        Next  *Node
}

type LRUCache struct {
        Head, Tail *Node
        Capacity   int
        Map        map[interface{}]*Node
}

func NewLRUCache(capacity int) *LRUCache {
        l := &LRUCache{}
        l.Capacity = capacity
        l.Head = &Node{}
        l.Tail = &Node{}
        l.Head.Next = l.Tail
        l.Tail.Prev = l.Head
        l.Head.Prev = nil
        l.Tail.Next = nil
        l.Map = make(map[interface{}]*Node)
        return l
}

//分离节点
func (l *LRUCache) detach(n *Node) {
        n.Prev.Next = n.Next
        n.Next.Prev = n.Prev
}

//节点插入头部
func (l *LRUCache) attach(n *Node) {
        n.Prev = l.Head
        n.Next = l.Head.Next
        l.Head.Next = n
        n.Next.Prev = n
}

func (l *LRUCache) Put(k interface{}, v interface{}) {
        oldV, ok := l.Map[k]
        if ok {
                l.detach(oldV)
                oldV.Value = v
        } else {
                var n *Node
                if len(l.Map) >= l.Capacity {
                        n = l.Tail.Prev
                        l.detach(n)
                        delete(l.Map, n.Key)
                } else {
                        n = new(Node)
                }
                n.Key = k
                n.Value = v
                l.Map[k] = n
                l.attach(n)
        }
}

func (l *LRUCache) Get(k interface{}) interface{} {
        v, ok := l.Map[k]
        if ok {
                l.detach(v)
                l.attach(v)
                return v.Value
        }
        return -1
}

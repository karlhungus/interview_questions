package main

const hash = 512

type node struct {
	val  int
	next *node
}

type MyHashSet struct {
	h []*node
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{h: make([]*node, hash)}
}

func (this *MyHashSet) Add(key int) {
	h := key % hash
	n := this.h[h]
	if n == nil {
		this.h[h] = &node{val: key}
		return
	} else if n.val == key {
		return
	}
	for n.next != nil {
		if n.next.val == key {
			return
		}
		n = n.next
	}
	n.next = &node{val: key}
}

func (this *MyHashSet) Remove(key int) {
	h := key % hash
	n := this.h[h]
	if n == nil {
		return
	}
	if n.val == key {
		this.h[h] = n.next
		return
	}
	for n.next != nil {
		if n.next.val == key {
			n.next = n.next.next
			return
		}
		n = n.next
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	h := key % hash
	n := this.h[h]
	if n == nil {
		return false
	} else if n.val == key {
		return true
	}
	for n.next != nil {
		if n.next.val == key {
			return true
		}
		n = n.next
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

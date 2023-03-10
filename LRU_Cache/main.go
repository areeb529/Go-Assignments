package main

import (
	"fmt"
)

// Node structure
type Node struct {
	prev *Node
	next *Node
	val  int
}

var mp1 = make(map[int]int)
var mp2 = make(map[int]*Node)
var currCapacity = 0

// Doubly linked list structure
type DLL struct {
	head *Node
	tail *Node
}

// Reciever type constructor for doubly linked list
func (d *DLL) DLLConstructor() {
	d = &DLL{
		head: &Node{
			val: -1,
		},
		tail: &Node{
			val: -1,
		},
	}
	d.head.next = d.tail
	d.tail.prev = d.head
}

// Return type Constructor for doubly linked list
func DLLConstructor() *DLL {
	d := &DLL{}
	h := &Node{
		val: -1,
	}
	t := &Node{
		val: -1,
	}
	d.head = h
	d.tail = t
	d.head.next = d.tail
	d.tail.prev = d.head
	return d
}

// Push most recently used value at the back
func (d *DLL) push_back(v int) {
	newNode := &Node{
		prev: d.tail.prev,
		next: d.tail,
		val:  v,
	}

	d.tail.prev.next = newNode
	d.tail.prev = newNode
	mp2[v] = newNode
	currCapacity++
}

// Remove least recently used value from the front
func (d *DLL) pop_front() {
	d.head.next.next.prev = d.head
	d.head.next = d.head.next.next
	currCapacity--
}

// LRU cache structure
type LRUCache struct {
	l        *DLL
	capacity int
}

// LRU Cache constructor
func (lru *LRUCache) LRUCacheConstructor(capacity int) {

	//lru.l.DLLConstructor()   // reciever type call
	lru.l = DLLConstructor()   // return type call
	lru.capacity = capacity
}

// delete node at a particular location
func (n *Node) deleteFromList() {
	n.prev.next = n.next
	n.next.prev = n.prev
	currCapacity--
}

// get value corresponding to a key
func (lru *LRUCache) get(key int) int {
	val, ok := mp1[key]
	if ok {
		mp2[key].deleteFromList()
		delete(mp2, key)
		lru.l.push_back(key)
		return val
	} else {
		return -1
	}
}

// put key, value pair into the lru cache
func (lru *LRUCache) put(key int, value int) {
	_, ok := mp1[key]
	if ok {
		mp2[key].deleteFromList()
		delete(mp2, key)
		lru.l.push_back(key)
		mp1[key] = value
	} else {
		mp1[key] = value
		if currCapacity < lru.capacity {
			lru.l.push_back(key)
		} else {
			delete(mp1, lru.l.head.next.val)
			delete(mp2, lru.l.head.next.val)
			lru.l.pop_front()
			lru.l.push_back(key)
		}
	}
}

func main() {
	var lru LRUCache
	lru.LRUCacheConstructor(2)
	lru.put(1, 1)
	lru.put(2, 2)
	fmt.Println(lru.get(1))
	fmt.Println(lru.l.head.next.val)
	lru.put(3, 3)
	fmt.Println(lru.get(2))
	lru.put(4, 4)
	fmt.Println(lru.get(1))
	fmt.Println(lru.get(3))
	fmt.Println(lru.get(4))
}

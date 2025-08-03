// Package linkedlist implements a linked list type supporting operations
// including Append, Insert, Get, Remove.
package linkedlist

import (
	"fmt"
	"go-book/debug"
	"strconv"
)

type LinkedList struct {
	sentinel *Node
	*debug.Debugger
}

type Node struct {
	key  int64
	val  int64
	next *Node
}

func ConstructLinkedList(debugActive bool) *LinkedList {
	sentinel := &Node{key: 0, val: 0, next: nil}
	return &LinkedList{sentinel: sentinel, Debugger: &debug.Debugger{Active: debugActive}}
}

func (l *LinkedList) Append(key, val int64) {
	l.DebugLog("\nAPPEND: key = %v, val = %v, linked list = %v\n", key, val, l)
	node := l.sentinel
	for node.next != nil {
		node = node.next
	}
	node.next = &Node{key: key, val: val, next: nil}
	l.DebugLog("result: linked list = %v\n", l)
}

func (l *LinkedList) AppendMany(keys []int64, values []int64) error {
	if len(keys) != len(values) {
		return fmt.Errorf("AppendMany: keys and values slices must be of the same length")
	}
	l.DebugLog("\nAPPEND: keys = %v, values = %v, linked list = %v\n", keys, values, l)
	node := l.sentinel
	for node.next != nil {
		node = node.next
	}
	for i := range keys {
		node.next = &Node{key: keys[i], val: values[i], next: nil}
		node = node.next
	}
	l.DebugLog("result: linked list = %v\n", l)
	return nil
}

func (l *LinkedList) Insert(key int64, val int64, index uint) {
	l.DebugLog("\nINSERT: key = %v, val = %v, index = %v, linked list = %v\n", key, val, index, l)
	node := l.sentinel
	for i := uint(0); i < index; i++ {
		if node.next == nil {
			return
		}
		node = node.next
	}
	newNode := &Node{key: key, val: val, next: node.next}
	node.next = newNode
	l.DebugLog("result: linked list = %v\n", l)
}

func (l *LinkedList) Get(key int64) (int64, bool) {
	if l == nil {
		return 0, false
	}
	l.DebugLog("\nGet: key = %v, linked list = %v\n", key, l)
	node := l.sentinel
	for node != nil {
		if node.key == key {
			l.DebugLog("found %v: value = %v\n", node.key, node.val)
			return node.val, true
		}
		node = node.next
	}
	l.DebugLog("result: false\n")
	return 0, false
}

// Items
// Returns a slice containing keys and a slice containing values.
// TODO: Replace usage of this in the hashmap with an iterator.
func (l *LinkedList) Items() ([]int64, []int64) {
	l.DebugLog("\nItems: linked list = %v\n", l)
	var (
		keys   []int64
		values []int64
	)
	node := l.sentinel.next
	for node != nil {
		keys = append(keys, node.key)
		values = append(values, node.val)
		node = node.next
	}
	l.DebugLog("keys = %v, values = %v\n", keys, values)
	return keys, values
}

func (l *LinkedList) Remove(key int64) {
	l.DebugLog("\nREMOVE: key = %v, linked list = %v\n", key, l)
	node := l.sentinel
	for node.next != nil {
		if node.next.key == key {
			node.next = node.next.next
			break
		}
		node = node.next
	}
	l.DebugLog("result: linked list = %v\n", l)
}

func (l *LinkedList) String() string {
	node := l.sentinel.next
	var buf []string
	for node != nil {
		s := fmt.Sprintf("(%v, %v)", strconv.FormatInt(node.key, 10), strconv.FormatInt(node.val, 10))
		buf = append(buf, s)
		node = node.next
	}
	return fmt.Sprintf("%v", buf)
}

func (l *LinkedList) IsEmpty() bool {
	return l.sentinel.next == nil
}

func (l *LinkedList) Length() int {
	node := l.sentinel.next
	length := 0
	for node != nil {
		node = node.next
		length++
	}
	l.DebugLog("\nLENGTH: linked list = %v, length = %v\n", l, length)
	return length
}

func (l *LinkedList) Clear() {
	l.sentinel.next = nil
	l.DebugLog("\nCLEAR: linked list = %v\n", l)
}

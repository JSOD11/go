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
	val  int64
	next *Node
}

func ConstructLinkedList(debugActive bool) *LinkedList {
	sentinel := &Node{val: 0, next: nil}
	return &LinkedList{sentinel: sentinel, Debugger: &debug.Debugger{Active: debugActive}}
}

func (l *LinkedList) Append(val int64) {
	l.DebugLog("\nAPPEND: val = %v, linked list = %v\n", val, l)
	node := l.sentinel
	for node.next != nil {
		node = node.next
	}
	node.next = &Node{val: val, next: nil}
	l.DebugLog("result: linked list = %v\n", l)
}

func (l *LinkedList) AppendMany(values ...int64) {
	l.DebugLog("\nAPPEND: val = %v, linked list = %v\n", values, l)
	node := l.sentinel
	for node.next != nil {
		node = node.next
	}
	for _, val := range values {
		node.next = &Node{val: val, next: nil}
		node = node.next
	}
	l.DebugLog("result: linked list = %v\n", l)
}

func (l *LinkedList) Insert(val int64, index uint) {
	l.DebugLog("\nINSERT: val = %v, index = %v, linked list = %v\n", val, index, l)
	node := l.sentinel
	for i := uint(0); i < index; i++ {
		if node.next == nil {
			return
		}
		node = node.next
	}
	newNode := &Node{val: val, next: node.next}
	node.next = newNode
	l.DebugLog("result: linked list = %v\n", l)
}

func (l *LinkedList) Contains(val int64) bool {
	l.DebugLog("\nCONTAINS: val = %v, linked list = %v\n", val, l)
	node := l.sentinel
	for node != nil {
		if node.val == val {
			l.DebugLog("result: true\n")
			return true
		}
		node = node.next
	}
	l.DebugLog("result: false\n")
	return false
}

func (l *LinkedList) Remove(val int64) {
	l.DebugLog("\nREMOVE: val = %v, linked list = %v\n", val, l)
	node := l.sentinel
	for node.next != nil {
		if node.next.val == val {
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
		buf = append(buf, strconv.FormatInt(node.val, 10))
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

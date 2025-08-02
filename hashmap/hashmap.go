package hashmap

import (
	"fmt"
	"go-book/debug"
	"go-book/linkedlist"
)

type HashMap struct {
	table     []*linkedlist.LinkedList
	itemsHeld int
	hashFn    func(int64, int) int
	*debug.Debugger
}

func moduloHash(key int64, tableCap int) int {
	return int(key) % tableCap
}

func ConstructHashMap(debugActive bool) *HashMap {
	ll := linkedlist.ConstructLinkedList(false)
	table := []*linkedlist.LinkedList{ll}
	return &HashMap{table: table, itemsHeld: 0, hashFn: moduloHash, Debugger: &debug.Debugger{Active: debugActive}}
}

func (h *HashMap) Put(key, value int64) {
	h.DebugLog("\nPUT: key = %v, val = %v\n", key, value)
	if h.tableOverweight() {
		h.growTable()
	}
	bucket := h.hashFn(key, cap(h.table))
	if bucket >= len(h.table) {
		h.table = h.table[:bucket+1]
	}
	if h.table[bucket] == nil {
		h.table[bucket] = linkedlist.ConstructLinkedList(false)
	}
	h.table[bucket].Append(key, value)
	h.itemsHeld++
	h.DebugLog("%v\n", h)
	return
}

func (h *HashMap) Get(key int64) (int64, bool) {
	h.DebugLog("\nGET: key = %v\n", key)
	bucket := h.hashFn(key, cap(h.table))
	if bucket >= len(h.table) {
		h.DebugLog("result = 0, ok = false\n")
		return 0, false
	}
	result, ok := h.table[bucket].Get(key)
	h.DebugLog("result = %v, ok = %v\n", result, ok)
	return result, ok
}

func (h *HashMap) Remove(key int64) {
	return
}

func (h *HashMap) tableOverweight() bool {
	return h.itemsHeld >= cap(h.table)/2
}

func (h *HashMap) growTable() {
	h.DebugLog("\nGROWING TABLE: items held = %v, length = %v, capacity = %v\n", h.itemsHeld, len(h.table), cap(h.table))
	oldTable := h.table
	h.table = make([]*linkedlist.LinkedList, h.itemsHeld+1, 2*(cap(oldTable)+1))
	for _, ll := range oldTable {
		if ll != nil {
			keys, values := ll.Items()
			for i := range keys {
				bucket := h.hashFn(keys[i], cap(h.table))
				if bucket >= len(h.table) {
					h.table = h.table[:bucket+1]
				}
				if h.table[bucket] == nil {
					h.table[bucket] = linkedlist.ConstructLinkedList(false)
				}
				h.table[bucket].Append(keys[i], values[i])
			}
		}
	}
	oldTable = nil
	h.DebugLog("result: items held = %v, length = %v, capacity = %v\n", h.itemsHeld, len(h.table), cap(h.table))
}

func (h *HashMap) String() string {
	var buf []string
	buf = append(buf, "\n——————— Table ———————\n")
	for i := range h.table {
		buf = append(buf, fmt.Sprintf("%v | %v\n", i, h.table[i]))
	}
	return fmt.Sprintf("\n%v\n", buf)
}

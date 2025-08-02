package hashmap

import "testing"

const (
	debugLogs = true
)

func TestHashMapBasic(t *testing.T) {
	h := ConstructHashMap(debugLogs)
	h.Put(3, 3)
	h.Put(5, 5)
	h.Put(8, 8)
	h.Put(25, 1)
	h.Put(150, 3)

	h.Get(3)
	h.Get(8)
	h.Get(10)
	h.Get(100)
	h.Get(150)

	h.Put(235, 121)
}

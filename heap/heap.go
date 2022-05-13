package heap

const (
	MIN_HEAP = iota
	MAX_HEAP
)

// Element is the element of the heap.
// The heap compares the priority of elements by Priority().
// In min-heap, the smaller the Priority(), the higher the priority.
// In max-heap, the bigger the Priority(), the higher the priority.
type Element interface {
	Priority() int
}

// *Heap implements heap.Interface defined in container/heap.go
// and sort.Interface defined in sort/sort.go.
type Heap struct {
	flag     int // can be MIN_HEAP or MAX_HEAP
	size     int
	capacity int
	Array    []Element
}

// NewHeap returns a pointer to an instance of Heap.
// cap is the capacity of Heap.
// flag can be MIN_HEAP or MAX_HEAP.
func NewHeap(cap, flag int) *Heap {
	heap := &Heap{
		flag:     flag,
		size:     0,
		capacity: cap,
		Array:    make([]Element, 0, cap),
	}
	return heap
}

// Len returns the size of h.Array.
func (h *Heap) Len() int {
	return h.size
}

// Less compares the priority between two elements.
func (h *Heap) Less(i, j int) bool {
	if h.flag == 0 {
		return h.Array[i].Priority() < h.Array[j].Priority()
	}
	return h.Array[i].Priority() > h.Array[j].Priority()
}

// Swap
func (h *Heap) Swap(i, j int) {
	h.Array[i], h.Array[j] = h.Array[j], h.Array[i]
}

// Push adds an element to the heap.
func (h *Heap) Push(x interface{}) {
	h.Array = append(h.Array, x.(Element))
	h.shiftUp(h.size)
	h.size++
}

// Pop removes and returns the top element of th heap.
func (h *Heap) Pop() interface{} {
	h.Swap(0, h.size-1)
	v := h.Array[h.size-1]
	h.Array = h.Array[:h.size-1]
	h.size--
	h.shiftDown(0)
	return v
}

// Peek returns the top element of the heap.
func (h *Heap) Peek() Element {
	if h.size > 0 {
		return h.Array[0]
	}
	return nil
}

func (h *Heap) shiftUp(i int) {
	for {
		j := (i - 1) / 2
		if j == i || !h.Less(i, j) {
			break
		}
		h.Swap(i, j)
		i = j
	}
}

func (h *Heap) shiftDown(i0 int) {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= h.size || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < h.size && h.Less(j2, j1) {
			j = j2
		}

		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
}

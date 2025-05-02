package utils

// An IntHeap is a min-heap of ints.
type Priority []PriorityItem

type PriorityItem struct {
	Id       string // id of the RSS feed
	NextCall int64  // timestamp of the next fetch call
	LastCall int64  // timestamp of the last fetch call
}

func (h *Priority) Len() int {
	return len(*h)
}

func (h *Priority) Less(i, j int) bool {
	return (*h)[i].NextCall < (*h)[j].NextCall
}

func (h *Priority) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Priority) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(PriorityItem))
}

func (h *Priority) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

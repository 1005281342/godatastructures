package heap

type Heap interface {
	Len() int
	Empty() bool
	Top() (int, bool)
	Pop() (int, bool)
	Remove(index int) (int, bool)
	Push(v int)
}

package set

type Set interface {
	// Add 往集合中添加元素k
	Add(k interface{}) bool
	// Del 从集合中移除元素k
	Del(k interface{}) bool
	// Contain 集合中是否包含元素k
	Contain(k interface{}) bool
	// Len 集合的长度
	Len() int
	// Clear 清空集合
	Clear()
}

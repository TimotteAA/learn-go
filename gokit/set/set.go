package set

type Set[T comparable] interface {
	Add(val T)
	Delete(val T)
	Has(val T) bool
	Keys() []T
}

type MapSet[T comparable] struct {
	m map[T]interface{}
}

// 工厂函数
// set是动态的，不考虑容量
func NewMapSet[T comparable]() *MapSet[T] {
	return &MapSet[T]{
		m: make(map[T]interface{}),
	}
}

func (ms *MapSet[T]) Add(val T) {
	// 赋值匿名结构体
	ms.m[val] = struct{}{}
}	

func (ms *MapSet[T]) Delete(val T) {
	delete(ms.m, val)
}

func (ms *MapSet[T]) Has(val T) bool {
	_, exists := ms.m[val] 
	return exists
}

func (ms *MapSet[T]) Keys() []T {
	keys := make([]T, 0, len(ms.m))
	for val, _ := range ms.m {
		keys = append(keys, val)
	}
	return keys
}
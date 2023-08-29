package synx

import "sync"

// 支持泛型的sync.Map
type Map[K comparable, V any] struct {
	data *sync.Map
}

func (m *Map[K, V]) Load(key K) (V, bool) {
	var value V
	mapValue, ok := m.data.Load(key)
	// 类型断言
	if mapValue != nil {
		value = mapValue.(V)
	}
	return value, ok
}

func (m *Map[K, V]) LoadOrStore(key K, val V) (actual V, loaded bool) {
	var mapVal any
	mapVal, loaded = m.data.LoadOrStore(key, val)
	if mapVal != nil {
		actual = mapVal.(V)
	}
	return
}
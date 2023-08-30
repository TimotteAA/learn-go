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

func (m *Map[K, V]) Store(key K, val V) {
	m.data.Store(key, val)
}

func (m *Map[K, V]) Delete(key K) {
	m.data.Delete(key)
}

// 如果func返回false，将终止遍历
// 在原生的基础上做了一层类型断言
func (m *Map[K, V]) Range(f func(key K, val V) bool) {
	m.data.Range(func(key, value any) bool {
		var (
			k K
			v V
		)
		if key != nil {
			k = key.(K)
		}
		if value != nil {
			v = value.(V)
		}
		return f(k, v) 
	})
}

func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	var val any
	val, loaded = m.data.LoadAndDelete(key)
	if val != nil {
		value = val.(V)
	}
	return 
}

// 懒惰地加载或生成值，同时确保在多goroutine环境下的并发安全性。
// 如果多个goroutine尝试为同一个键加载或生成值，这个方法确保只有一个值会被存储，并且所有的goroutine都将获得相同的值。
func (m *Map[K, V]) LoadOrStoreFunc(key K, fn func() (V, error)) (actual V, loaded bool, err error) {
	val, ok := m.Load(key);
	if ok {
		return val, ok, nil
	}
	// 没有值，调用func创建
	val, err = fn()
	if err != nil {
		// 此处error已有值了
		return 
	}
	actual, loaded = m.LoadOrStore(key, val)
	return
}
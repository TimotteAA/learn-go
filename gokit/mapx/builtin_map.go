package mapx

type BuiltinMap[K comparable, V any] struct {
	data map[K]V
}

func NewBuiltInMap[K comparable, V any]() *BuiltinMap[K, V] {
	return &BuiltinMap[K, V]{
		data: make(map[K]V),
	}
}

func (m *BuiltinMap[K, V]) Add(key K, val V) error {
	m.data[key] = val
	return nil
}

func (m *BuiltinMap[K, V]) Has(key K) bool {
	_, exists := m.data[key]
	return exists
}

func (m *BuiltinMap[K, V]) Get(key K) (V, bool) {
	val, exists := m.data[key]
	return val, exists
}

func (m *BuiltinMap[K, V]) Delete(key K) (V, bool) {
	val, ok := m.data[key]
	delete(m.data, key)
	return val, ok
}

func (m *BuiltinMap[K, V]) Keys() []K {
	return Keys[K, V](m.data)
}

func (m *BuiltinMap[K, V]) Values() []V {
	return Values[K, V](m.data)
}

func (m *BuiltinMap[K, V]) Entries() ([]K, []V) {
	return Entries[K, V](m.data)
}
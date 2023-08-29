package mapx

// 返回map所有的keys数组
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
} 

func Values[K comparable, V any](m map[K]V) []V {
	vals := make([]V, 0, len(m))
	for _, val := range m {
		vals = append(vals, val)
	}
	return vals
}

func Entries[K comparable, V any](m map[K]V) ([]K, []V){
	keys := make([]K, 0, len(m))
	vals := make([]V, 0, len(m))
	for key, val := range m {
		keys = append(keys, key)
		vals = append(vals, val)
	}
	return keys, vals
}
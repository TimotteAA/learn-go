package slice

import "gokit/internal"

// 在idx添加元素ele
func Add[T any](s []T, idx int, ele T) []T {
	if idx < 0 || idx >= len(s) {
		internal.NewErrorIndexOutOfRange(idx, len(s))
	}

	// T类型的零值
	var zeroValue T
	s = append(s, zeroValue)
	// 从最后一个元素开始，把前面的往后移
	for i := len(s) - 1; i > idx; i-- {
		s[i] = s[i-1]
	}
	s[idx] = ele
	return s;
}


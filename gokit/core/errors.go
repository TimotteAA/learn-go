package core

import "fmt"

// 切片索引越界异常
func NewErrorIndexOutOfRange(idx int, len int) error {
	return fmt.Errorf("gokit: Index %d out of range, last idx: %d", idx, len - 1)
}
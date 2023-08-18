package slice

// 判断切片中是否存在某个可比较元素，调用ContainsFunc
func Contains[T comparable](s []T, dest T) bool {
	return ContainsFunc[T](s, func(src T) bool {
		return src == dest
	})
}

// 判断切片中是否有元素满足equa
func ContainsFunc[T any](s []T, equa func(src T) bool) bool {
	// 遍历比较是否有元素满足equaFunc
	for _, val := range s {
		if equa(val) {
			return true;
		}
	}
	return false;
}

// 判断src中是否存在dest中的任何一个元素
func ContainsAny[T comparable](src, dest []T) bool {

}

func ContainsAnyFunc[T any](src []T, equal func(src T) bool) bool {
	for _, val := range src {
		if equal(val) {
			return true;
		}
	}
	return false;
}
package main

func findIndex[T comparable](s []T, dest T) (int, error) {
	for idx, val := range s {
		if val == dest {
			return idx, nil
		}
	}
	return -1, nil
}

func main()  {
	s := make([]int, 3, 10)
	s = append(s, 5, 6, 78)
	idx, err := findIndex[int](s, 78)
	if (err == nil) {
		println("78的索引是 ", idx)
	}
}
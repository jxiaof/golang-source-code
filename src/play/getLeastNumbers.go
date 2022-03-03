package main

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	if k == len(arr) {
		return arr
	}
	if k > len(arr) {
		return arr
	}
	var r []int
	for i := 0; i < k; i++ {
		r = append(r, arr[i])
	}
	for i := k; i < len(arr); i++ {
		if arr[i] < r[0] {
			r[0] = arr[i]
			for j := 1; j < len(r); j++ {
				if r[j-1] > r[j] {
					r[j-1], r[j] = r[j], r[j-1]
				}
			}
		}
	}
	return r

}

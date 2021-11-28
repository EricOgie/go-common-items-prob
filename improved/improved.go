package main

// checkCommonItems returns true if the two arrays have at least one common item
// time complexity: O(n + m) = O(n)
func checkCommonItems(arr1 []int, arr2 []int) bool {
	var set1 = make(map[int]bool)
	for x := 0; x < len(arr1); x++ {
		set1[arr1[x]] = true
	}

	for _, v := range arr2 {
		if _, exists := set1[v]; exists {
			return true
		}
	}
	return false
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := []int{13, 19, 25, 31, 37, 43, 49, 55, 61}
	println(checkCommonItems(arr1, arr2))
}

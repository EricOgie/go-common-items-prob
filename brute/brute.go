package main

// Brute force solution
// Time complexity: O(n^2)
func checkCommonItems(arr1 []int, arr2 []int) bool {
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				return true
			}
		}
	}

	return false
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := []int{7, 13, 19, 25, 31, 37, 43, 49, 55, 61}
	println(checkCommonItems(arr1, arr2))
}

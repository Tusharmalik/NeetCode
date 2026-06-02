package main

//func main() {
//	arr := []int{1, 2, 4, 8, 13, 4, 15}
//	arr = []int{5, 6, 1, 2, 3, 4}
//	//fmt.Println(linearSearchFindAll(arr, 3, 0))
//
//	fmt.Println(rotatedBinarySearch(arr, 6, 0, len(arr)-1))
//}

func isSorted(arr []int, index int) bool {
	if index == len(arr)-1 {
		return true
	}

	return arr[index] < arr[index+1] && isSorted(arr, index+1)
}

func linearSearchFindIndex(arr []int, target int, index int) int {
	if index == len(arr)-1 {
		return -1
	}

	if arr[index] == target {
		return index
	}

	return linearSearchFindIndex(arr, target, index+1)
}

func linearSearchFindIndexLast(arr []int, target int, index int) int {
	if index == -1 {
		return -1
	}

	if arr[index] == target {
		return index
	}

	return linearSearchFindIndexLast(arr, target, index-1)
}

func linearSearchFindAll(arr []int, target int, index int) []int {
	var localArr []int

	if index == len(arr)-1 {
		return localArr
	}

	if arr[index] == target {
		localArr = append(localArr, index)
	}

	temp := linearSearchFindAll(arr, target, index+1)

	localArr = append(localArr, temp...)
	return localArr
}

func rotatedBinarySearch(arr []int, target int, start int, end int) int {
	if start > end {
		return -1
	}

	mid := start + (end-start)/2

	if arr[mid] == target {
		return mid
	}

	if arr[start] <= arr[mid] {
		if target >= arr[start] && target < arr[mid] {
			return rotatedBinarySearch(arr, target, start, mid-1)
		}
		return rotatedBinarySearch(arr, target, mid+1, end)
	}

	if target >= arr[mid] && target <= arr[end] {
		return rotatedBinarySearch(arr, target, mid+1, end)
	}
	return rotatedBinarySearch(arr, target, start, mid-1)
}

package main

import "fmt"

func main() {
	//triangleReverse(5, 0)
	//triangleNormal(0, 0, 5)
	//triangleNormal2(5, 0)
	arr := []int{5, 4, 3, 2, 1}

	//fmt.Println(bubbleSort(arr, len(arr)-1, 0))
	//fmt.Println(arr)

	//fmt.Println(selectionSort(arr, len(arr)-1, 0, 0))
	//fmt.Println(arr)

	//fmt.Println(mergeSort(arr))
	//fmt.Println(arr)

	//mergeSortInplace(arr, 0, len(arr)-1)
	//fmt.Println(arr)

	//fmt.Println(bubbleSortLoop(arr))

	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func bubbleSort(arr []int, r, c int) []int {
	if r == 0 {
		return arr
	}

	if r > c {
		if arr[c] > arr[c+1] {
			temp := arr[c]
			arr[c] = arr[c+1]
			arr[c+1] = temp
		}
		return bubbleSort(arr, r, c+1)
	}

	return bubbleSort(arr, r-1, 0)
}

func bubbleSortLoop(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		//temp := i
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				//temp = j
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		//arr[i], arr[temp] = arr[temp], arr[i]
	}
	return arr
}

func selectionSort(arr []int, r, c, max int) []int {
	if r == 0 {
		return arr
	}

	if r >= c {
		if arr[c] > arr[max] {
			max = c
		}
		return selectionSort(arr, r, c+1, max)
	}
	temp := arr[r]
	arr[r] = arr[max]
	arr[max] = temp

	return selectionSort(arr, r-1, 0, 0)
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	var merged []int
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	for i < len(left) {
		merged = append(merged, left[i])
		i++
	}
	for j < len(right) {
		merged = append(merged, right[j])
		j++
	}
	return merged
}

func mergeSortInplace(arr []int, left, right int) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2
	mergeSortInplace(arr, left, mid)
	mergeSortInplace(arr, mid+1, right)

	mergeInplace(arr, left, mid, right)
}

func mergeInplace(arr []int, left, mid, right int) {
	start2 := mid + 1
	start := left

	if arr[mid] <= arr[start2] {
		return
	}

	for start <= mid && start2 <= right {
		if arr[start] <= arr[start2] {
			start++
		} else {
			val := arr[start2]
			index := start2

			for start != index {
				arr[index] = arr[index-1]
				index--
			}
			arr[start] = val
			start2++
			start++
		}
	}
}

func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}

	start := low
	end := high
	mid := low + (high-low)/2
	pivot := arr[mid]

	for start <= end {
		for arr[start] < pivot {
			start++
		}

		for arr[end] > pivot {
			end--
		}

		if start <= end {
			arr[start], arr[end] = arr[end], arr[start]
			start++
			end--
		}
	}

	quickSort(arr, low, end)
	quickSort(arr, start, high)
}

func triangleReverse(r, c int) {
	if r == 0 {
		return
	}

	if r > c {
		fmt.Printf("*")
		triangleReverse(r, c+1)
	} else {
		fmt.Printf("\n")
		triangleReverse(r-1, 0)
	}
}

func triangleNormal2(r, c int) {
	if r == 0 {
		return
	}

	if r > c {
		triangleNormal2(r, c+1)
		fmt.Printf("*")
	} else {
		triangleNormal2(r-1, 0)
		fmt.Printf("\n")
	}

}

func triangleNormal(r, c, n int) {
	if r == n {
		return
	}

	if r >= c {
		fmt.Printf("*")
		triangleNormal(r, c+1, n)
	} else {
		fmt.Printf("\n")
		triangleNormal(r+1, 0, n)
	}

}

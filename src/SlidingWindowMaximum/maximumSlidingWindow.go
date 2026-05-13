package main

import (
	"fmt"
	"math"
	"math/bits"
)

func maxSlidingWindowBruteForce(nums []int, k int) []int {
	n := len(nums)
	res := make([]int, 0, n-k+1)

	for i := 0; i < n-k+1; i++ {
		localMax := math.MinInt32
		for j := i; j < i+k; j++ {
			if nums[j] > localMax {
				localMax = nums[j]
			}
		}
		res = append(res, localMax)
	}
	return res
}

func maxSlidingWindowDynamicProgramming(nums []int, k int) []int {
	n := len(nums)
	leftMax := make([]int, n)
	rightMax := make([]int, n)

	leftMax[0] = nums[0]
	rightMax[n-1] = nums[n-1]

	for i := 1; i < n; i++ {
		if i%k == 0 {
			leftMax[i] = nums[i]
		} else {
			leftMax[i] = max(leftMax[i-1], nums[i])
		}
		//fmt.Println("i:", i, "i%k:", i%k, " leftMax:", leftMax[i], " nums[i]:", nums[i], " leftMax[i-1]:", leftMax[i-1])
	}
	//fmt.Println(leftMax)

	for i := 1; i < n; i++ {
		if (n-1-i)%k == 0 {
			rightMax[n-1-i] = nums[n-1-i]
		} else {
			rightMax[n-1-i] = max(rightMax[n-i], nums[n-1-i])
		}
		//fmt.Println("i:", i, "(n-1-i)%k:", (n-1-i)%k, " rightMax:", rightMax[n-1-i], " nums[n-1-i]:", nums[n-1-i], " rightMax[n-i]:", rightMax[n-i])
	}
	//fmt.Println(rightMax)

	output := make([]int, n-k+1)

	for i := 0; i < n-k+1; i++ {
		//fmt.Println("i:", i, " leftMax[i+k-1]:", leftMax[i+k-1], " rightMax[i]:", rightMax[i])
		output[i] = max(leftMax[i+k-1], rightMax[i])
	}
	return output
}

type SegmentTree struct {
	n    int
	tree []int
}

func NewSegmentTree(N int, A []int) *SegmentTree {
	n := N
	for bits.OnesCount(uint(n)) != 1 {
		n++
	}

	st := &SegmentTree{
		n: n,
	}
	st.build(N, A)
	return st
}

func (st *SegmentTree) build(N int, A []int) {
	st.tree = make([]int, 2*st.n)
	for i := range st.tree {
		st.tree[i] = math.MinInt
	}
	for i := 0; i < N; i++ {
		st.tree[st.n+i] = A[i]
	}
	for i := st.n - 1; i > 0; i-- {
		st.tree[i] = max(st.tree[i<<1], st.tree[i<<1|1])
	}
}

func (st *SegmentTree) Query(l, r int) int {
	res := math.MinInt
	l += st.n
	r += st.n + 1

	for l < r {
		if l&1 == 1 {
			res = max(res, st.tree[l])
			l++
		}
		if r&1 == 1 {
			r--
			res = max(res, st.tree[r])
		}
		l >>= 1
		r >>= 1
	}
	return res
}

func maxSlidingWindowSegmentTree(nums []int, k int) []int {
	n := len(nums)
	segTree := NewSegmentTree(n, nums)
	output := make([]int, n-k+1)

	for i := 0; i <= n-k; i++ {
		output[i] = segTree.Query(i, i+k-1)
	}

	return output
}

func maxSlidingWindowHashSegmentTree(nums []int, k int) []int {
	hash := make(map[string]int)
	n := len(nums)

	for l := 0; l < n; l++ {
		r := l + 1
		strLeft := fmt.Sprintf("%d:%d", l, l)
		hash[strLeft] = nums[l]

		for r < n {
			strRight := fmt.Sprintf("%d:%d", r, r)
			str := fmt.Sprintf("%d:%d", l, r)
			strPrev := fmt.Sprintf("%d:%d", l, r-1)
			hash[strRight] = nums[r]
			hash[str] = max(hash[strPrev], hash[strRight])

			r++
		}
	}

	//fmt.Println(hash)

	res := make([]int, 0, n-k+1)

	for i := 0; i < n-k+1; i++ {
		//fmt.Println("i:", i, "i+k-1:", i+k-1, " hash[fmt.Sprintf(\"%d:%d\", i, i+k-1)]:", hash[fmt.Sprintf("%d:%d", i, i+k-1)])
		res = append(res, hash[fmt.Sprintf("%d:%d", i, i+k-1)])
	}

	return res
}

func main() {
	nums, k := []int{1, 2, 1, 0, 4, 2, 6}, 3
	fmt.Println(nums, k)
	//fmt.Println(maxSlidingWindowDynamicProgramming(nums, k))
	//fmt.Println(maxSlidingWindowBruteForce(nums, k))
	fmt.Println(maxSlidingWindowHashSegmentTree(nums, k))

	nums, k = []int{1, -1}, 1
	fmt.Println(maxSlidingWindowDynamicProgramming(nums, k))
	//fmt.Println(maxSlidingWindowBruteForce(nums, k))
}

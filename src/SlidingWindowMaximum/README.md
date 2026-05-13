# Sliding Window Maximum

You are given an array of integers `nums` and an integer `k`. There is a sliding window of size `k` that starts at the left edge of the array. The window slides one position to the right until it reaches the right edge of the array.

Return a list that contains the maximum element in the window at each step.

### Example 1:

```
Input: nums = [1,2,1,0,4,2,6], k = 3

Output: [2,2,4,4,6]

Explanation:
Window position                Max
---------------               -----
[1  2  1] 0  4  2  6            2
 1 [2  1  0] 4  2  6            2
 1  2 [1  0  4] 2  6            4
 1  2  1 [0  4  2] 6            4
 1  2  1  0 [4  2  6]           6
``` 

### Constraints:
- `1 <= nums.length <= 10^5`
- `-10^4 <= nums[i] <= 10^4`
- `1 <= k <= nums.length`
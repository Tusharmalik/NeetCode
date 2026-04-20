# Permutation in String

You are given two strings `s1` and `s2`.

Return `true` if `s2` contains a permutation of `s1`, or `false` otherwise. That means if a permutation of `s1` exists as a substring of `s2`, then return `true`.

Both strings only contain lowercase letters.

### Example 1:

```
Input: s1 = "abc", s2 = "lecabee"

Output: true
```
Explanation: s2 contains one permutation of s1 ("cba").

### Example 2:

```
Input: s1 = "abc", s2 = "lecaabee"

Output: false
```
Explanation: s2 does not contain any permutation of s1.

### Constraints:
- `1 <= s1.length, s2.length <= 10^4`
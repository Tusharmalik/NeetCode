# Longest Repeating Character Replacement

You are given a string `s` consisting of only uppercase english characters and an integer `k`. You can choose up to `k` characters of the string and replace them with any other uppercase English character.

After performing at most `k` replacements, return the length of the longest substring which contains only one distinct character.

### Example 1:

```
Input: s = "XYYX", k = 2

Output: 4
```
Explanation: Replace the two 'Y' with 'X' to form "XXXX".

### Example 2:

```
Input: s = "AAABABB", k = 1

Output: 5
```
Explanation: Replace the one 'B' in the middle with 'A' to form "AAAAABB".

### Constraints:
- `1 <= s.length <= 10^5`
- `s` consists of only uppercase English letters.
- `0 <= k <= s.length`
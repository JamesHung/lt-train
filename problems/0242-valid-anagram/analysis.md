## Original Prompt
Given two strings `s` and `t`, return `true` if `t` is an anagram of `s`, and `false` otherwise. Both strings contain lowercase English letters and have length up to `5 * 10^4`.

## English Explanation
An anagram has identical character counts. Allocate an array of size 26, iterate through both strings simultaneously, incrementing for `s[i]` and decrementing for `t[i]`. If every bucket returns to zero, both strings share the same multiset of characters; otherwise they differ.

### Follow-up (Unicode)
Use a `map[rune]int` or Go's `range` over runes to count arbitrary Unicode characters. Alternatively, sort the rune slices, but the counting map keeps the solution linear.

## 中文詳解
1. 先檢查長度，若不同則直接回傳 `false`。
2. 建立長度 26 的陣列作為頻率表，掃描字串時遇到 `s[i]` 就加一，遇到 `t[i]` 就減一。
3. 迴圈結束後如果所有頻率都回到 0，代表 `s` 與 `t` 的字元組成完全相同，即為 anagram；只要有任一 bucket 非 0，就代表出現差異。
4. 若需要支援 Unicode，只要改用 `map[rune]int` 或排序 `[]rune` 即可。

## Complexity
- Time: `O(n)` for n = len(s) = len(t) since the strings are traversed once.
- Space: `O(1)` extra for the 26-length array (or `O(k)` unique runes in the Unicode variant).

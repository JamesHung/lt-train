## Original Prompt
You are given `n` product versions numbered `1..n` and an API `isBadVersion(version)` that returns whether a version failed quality checks. Every version after a bad one is also bad. Find the very first bad version while minimizing API calls.

## English Explanation
Binary search the answer space of version indices. Maintain `lo` and `hi` as the smallest and largest candidates that could still be the first bad. Check the midpoint `mid`; if `isBadVersion(mid)` is true, the first bad version is at `mid` or earlier so shrink the right boundary to `mid`. Otherwise the first bad is strictly after `mid`, so move `lo` to `mid + 1`. The loop terminates when both boundaries converge on the first bad version.

## 中文詳解
1. 設定左邊界 `lo = 1`、右邊界 `hi = n`，代表目前可能是第一個壞版本的區間。
2. 每回合計算 `mid = lo + (hi - lo)/2`，避免溢位並鎖定中間版本。
3. 如果 `isBadVersion(mid)` 回傳 `true`，表示答案在 `mid` 或更左邊，將右邊界縮到 `mid`。
4. 否則答案必定在 `mid` 之後，將左邊界更新為 `mid + 1`。
5. 當 `lo == hi` 時即找到第一個壞版本並回傳。

## Complexity
- Time: `O(log n)` API calls because each iteration halves the search space.
- Space: `O(1)` extra memory.

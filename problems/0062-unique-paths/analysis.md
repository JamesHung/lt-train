## Original Prompt
A robot starts at the top-left of an `m x n` grid and wants to reach the bottom-right. Each move can only go right or down. Given `m` and `n`, return the number of unique paths. Constraints: `1 <= m, n <= 100`, answer ≤ `2 * 10^9`.

## English Explanation
Use dynamic programming with a single row. Every cell's path count equals the sum of paths from the top and from the left. Initialize the first row to 1 (only right moves). Then for each subsequent row, sweep columns left-to-right updating `row[c] += row[c-1]`, effectively folding in the top value that `row[c]` already holds. The last element becomes the total paths. `debugUniquePaths` can print the row per iteration if needed.

## 中文詳解
1. 定義 `row[c]` 為目前列、欄 `c` 的路徑數，初始化為 1（第一列只能一路往右）。
2. 從第二列開始，左到右更新 `row[c] += row[c-1]`，因為上方路徑數已存在於 `row[c]`，左邊路徑數是 `row[c-1]`。
3. 走完第 `m` 列後，`row[n-1]` 就是總走法。
4. 開啟 `debugUniquePaths` 可以看到每列的累積情況。

## Complexity
- Time: `O(m*n)`
- Space: `O(n)`


• DP Example (3x4 grid)

  - Start with row = [1 1 1 1] (first row only moves right).
  - Row 2 updates left→right:
      - c=1: row[1] += row[0] → 2
      - c=2: row[2] += row[1] → 3
      - c=3: row[3] += row[2] → 4
        Row after update: [1 2 3 4]
  - Row 3 updates:
      - c=1: 3
      - c=2: 6
      - c=3: 10
        Row after update: [1 3 6 10]
  - Answer is last cell row[3] = 10, matching all unique right/down paths.
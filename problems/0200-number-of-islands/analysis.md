## Original Prompt
Given an `m x n` grid of `'1'`s (land) and `'0'`s (water), count the number of islands, where islands are groups of horizontally or vertically connected land cells. All grid edges are surrounded by water. Example inputs include:
```
[["1","1","1","1","0"],
 ["1","1","0","1","0"],
 ["1","1","0","0","0"],
 ["0","0","0","0","0"]] -> 1
```
and
```
[["1","1","0","0","0"],
 ["1","1","0","0","0"],
 ["0","0","1","0","0"],
 ["0","0","0","1","1"]] -> 3
```
Constraints: `1 <= m, n <= 300`.

## English Explanation
We scan the grid cell by cell. When we encounter a `'1'`, that means we discovered a new island. We increment the count and run a DFS (or BFS) from that cell to mark all connected land as visited, e.g., by turning each `'1'` to `'0'`. The DFS explores four directions (up, down, left, right) and stops when it hits water or the grid edge. Because each cell is visited at most once, the entire algorithm is `O(mn)`.

## 中文詳解
逐格掃描，遇到 `'1'` 就表示找到新島。把答案加一，然後從該格開始 DFS，把所有連通的土地都「淹掉」成 `'0'`（四方向）。DFS 遇到邊界或水就停。每個格子只會被訪問一次，所以時間複雜度 `O(mn)`，空間則為遞迴深度最多 `O(mn)`（或 `O(min(m,n))` 視島形狀而定）。

## Complexity
- Time: `O(mn)` because each cell is processed at most once.
- Space: `O(mn)` worst-case recursion stack (all land), typically less.

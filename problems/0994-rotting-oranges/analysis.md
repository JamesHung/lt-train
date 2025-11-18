## Original Prompt
Given an `m x n` grid, each cell is `0` (empty), `1` (fresh orange), or `2` (rotten). Every minute, any fresh orange adjacent (4-directional) to a rotten orange becomes rotten. Return the minimum number of minutes needed for all oranges to rot, or `-1` if impossible. Example: `[[2,1,1],[1,1,0],[0,1,1]] -> 4`, `[[2,1,1],[0,1,1],[1,0,1]] -> -1`, `[[0,2]] -> 0`. Constraints: `1 <= m, n <= 10`.

## English Explanation
Treat all initially rotten oranges as sources in a BFS. Push them into a queue with minute `0`, and track how many fresh oranges exist. Each BFS step takes a cell and spreads rot to its neighbors, marking them rotten and decrementing the fresh count. The minute recorded for each newly rotten orange is its parent minute plus one. We track the last minute we process. If any fresh oranges remain after the BFS, return `-1`; otherwise return the last minute (or `0` if there were no fresh oranges initially).

## 中文詳解
把所有一開始的 2 當作多源 BFS 的起點，記錄 `minute=0` 推進 queue，同時計算 fresh 橙子數量。每次從 queue 取出一個腐爛橙子，往四個方向擴散，只要鄰居是 1 就改成 2，fresh--，並把它以 `minute+1` 推入 queue，表示下一分鐘才腐爛。BFS 結束時如果還有 fresh > 0 代表某些橙子無法腐爛，回傳 -1；否則回傳最後一次擴散的分鐘數（若原本沒有新鮮橙則直接 0）。

## Complexity
- Time: `O(mn)` since each cell enters the queue at most once.
- Space: `O(mn)` for the queue in the worst case.

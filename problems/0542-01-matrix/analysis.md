## Original Prompt
Given an `m x n` binary matrix `mat`, return a matrix where each cell contains the distance to the nearest `0`. Adjacent cells are those sharing an edge. At least one zero exists.

**Example 1**
- Input: `mat = [[0,0,0],[0,1,0],[0,0,0]]`
- Output: `[[0,0,0],[0,1,0],[0,0,0]]`

**Example 2**
- Input: `mat = [[0,0,0],[0,1,0],[1,1,1]]`
- Output: `[[0,0,0],[0,1,0],[1,2,1]]`

Constraints:
- `1 <= m, n <= 10^4`, `1 <= m * n <= 10^4`
- `mat[i][j] ∈ {0, 1}` and there is at least one zero.

## English Explanation
Treat every zero as a source in a single multi-source BFS. Initialize distances to `0` for zeros and `+∞` for ones, pushing each zero into the queue. When popping a cell, relax its neighbors: if the neighbor's recorded distance is larger than `dist[current] + 1`, update it and enqueue the neighbor. Because BFS expands in layers, the first time we reach a cell we have the shortest path to a zero. Each cell enqueues at most once after its distance stabilizes, so the algorithm runs in linear time.

## 中文詳解
1. 先建立 `dist` 陣列，零的位置距離為 0，其他位置設成無限大，並把所有零同時放進佇列當作 BFS 起點。
2. 不斷從佇列取出座標 `(r,c)`，檢查四個方向 `(nr,nc)`。若 `dist[nr][nc] > dist[r][c] + 1`，表示找到更短路徑，更新距離並把 `(nr,nc)` 推回佇列。
3. 因為 BFS 會按照距離逐層擴散，第一次更新某格即代表最短距離，再也不會被更大的值覆蓋。
4. 直到佇列清空，就完成所有格子的最近 0 距離。過程中每格最多入佇列一次，所以時間線性、額外空間也只儲存距離與佇列。

## Complexity
- Time: `O(m * n)` — each cell processed with at most four neighbor relaxations.
- Space: `O(m * n)` — distance matrix plus BFS queue.

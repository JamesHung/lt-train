## Original Prompt
Given a binary matrix `isWater`, assign a non-negative height to each cell such that every water cell has height `0` and the absolute difference between adjacent cells is at most `1`. Return any valid assignment whose maximum height is as large as possible.

**Example 1**
- Input: `isWater = [[0,1],[0,0]]`
- Output: `[[1,0],[2,1]]`

**Example 2**
- Input: `isWater = [[0,0,1],[1,0,0],[0,0,0]]`
- Output: `[[1,1,0],[0,1,1],[1,2,2]]`

Constraints:
- `1 <= m, n <= 1000`
- `isWater[i][j] ∈ {0,1}` and at least one water cell exists.

## English Explanation
Treat every water cell as a source in the same BFS. Start by pushing all water coordinates into the queue with height `0` and mark all land cells as unassigned. Each time we pop a cell we try to relax its four neighbors. If a neighbor has not been assigned a height yet, set it to the current height plus one and enqueue it. Because BFS explores in expanding layers, the first time a land cell is visited is via its closest water cell, so the value we set is the smallest feasible height. Since we only grow from lower heights outward, the farthest land automatically gets the largest possible number, satisfying the maximization requirement without extra work.

## 中文詳解
把所有水格子同時視為 BFS 的起點，距離都是 0。其他格子先標記成未造訪（例如 -1）。從佇列依序取出座標 `(r,c)` 後檢查四個方向 `(nr,nc)`，若還沒被指派高度，就把它設成 `height[r][c] + 1` 並推入佇列。BFS 會一層一層擴張，所以第一次造訪某塊陸地就是它距離水最短的路徑；因此指定的高度也就最小，同時其他尚未造訪的格子必須位在更遠的地方，自然會得到更大的高度，達到題目要的「最大化最高點」。

## Complexity
- Time: `O(m * n)` — each cell enters the queue at most once and we inspect at most four neighbors.
- Space: `O(m * n)` — for the output matrix plus the BFS queue.


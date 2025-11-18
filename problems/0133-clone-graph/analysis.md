## Original Prompt
Given a reference node of a connected undirected graph, return a deep copy of the graph. Each node has a value and a list of neighbors. Values are unique and match their node indices.

**Example 1**
- Input: `[[2,4],[1,3],[2,4],[1,3]]`
- Output: `[[2,4],[1,3],[2,4],[1,3]]`

**Example 2**
- Input: `[[]]`
- Output: `[[]]`

**Example 3**
- Input: `[]`
- Output: `[]`

Constraints:
- `0 <= number of nodes <= 100`
- `1 <= Node.val <= 100`

## English Explanation
Traverse the graph (DFS or BFS) and maintain a map from original node pointer to its cloned node. When visiting a node, create its copy, store it in the map, and recursively/iteratively clone its neighbors. When a neighbor has already been cloned, just reuse the pointer from the map. Because the graph is connected, starting from the given node covers every vertex; handling `nil` input returns `nil`.

## 中文詳解
利用 DFS 或 BFS 走訪圖，同步維護「原節點 → 複製節點」的對照表。當第一次遇到某個節點時，建立新節點、記錄到 map，並遞迴（或佇列）複製所有鄰居；若鄰居已經複製過，就直接從 map 取出即可。這樣可以避免重複建點也能正確處理環。若輸入為 `nil`，直接回傳 `nil` 代表空圖。

## Complexity
- Time: `O(V + E)` — each vertex and edge is visited once during cloning.
- Space: `O(V)` — the recursion stack (or queue) plus the map storing clones.


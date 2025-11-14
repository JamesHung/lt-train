## Original Prompt
Given the root of a binary tree, return the length (number of edges) of the diameter, i.e., the longest path between any two nodes. The path may or may not pass through the root.

## English Explanation
Use a depth-first search that returns the height of each subtree. At every node, the longest path crossing it equals `leftHeight + rightHeight`, because you can go down the deepest descendant on each side. Track the maximum of these sums globally, then return the greater height plus one to the parent. The global maximum is the diameter.

## 中文詳解
1. 以 DFS 後序方式計算每個節點的高度（節點到底部的最長邊數）。
2. 對於每個節點，會形成一條「左子樹最高 + 右子樹最高」的路徑，這條路徑的邊數即 `leftHeight + rightHeight`。
3. 用全域變數記錄所有節點中最大的 `leftHeight + rightHeight`，即為答案。
4. 每次回傳給父節點的高度為 `max(leftHeight, rightHeight) + 1`。

## Complexity
- Time: `O(n)` visiting each node once.
- Space: `O(h)` recursion stack, where `h` is tree height (worst-case `O(n)`).

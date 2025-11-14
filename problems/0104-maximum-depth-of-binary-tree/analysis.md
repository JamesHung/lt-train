## Original Prompt
Given the root of a binary tree, return its maximum depth (number of nodes along the longest path from root to leaf).

## English Explanation
Perform a DFS: the maximum depth of a node equals `1 + max(depth(left), depth(right))`. Recursively compute depths for both children, choosing 0 for `nil`. Apply this definition bottom-up so each node contributes one level beyond its deeper child.

## 中文詳解
1. 若節點為 `nil`，深度為 0。
2. 對每個節點遞迴求出左子樹深度 `leftDepth` 與右子樹深度 `rightDepth`。
3. 回傳 `max(leftDepth, rightDepth) + 1`，代表包含當前節點的層數。
4. 根節點返回的數值即為整棵樹的最大深度。

## Complexity
- Time: `O(n)` since每個節點恰訪問一次。
- Space: `O(h)` recursion stack, with `h` 為樹高 (最糟 `O(n)` for鏈狀樹)。

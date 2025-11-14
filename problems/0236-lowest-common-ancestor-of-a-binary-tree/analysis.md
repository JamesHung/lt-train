## Original Prompt
Given the root of a binary tree plus two nodes `p` and `q`, return their lowest common ancestor (a node that is the lowest in the tree while still having both as descendants; a node can be a descendant of itself).

## English Explanation
Recurse from the root. If the current node is `nil`, `p`, or `q`, immediately return it. Otherwise, search both subtrees. Each recursive call returns either `nil` (if neither node is in that subtree) or whichever of `p`/`q` it found—or the LCA if both live below that subtree. After both sides return:
- If both sides are non-`nil`, the current node is the LCA, so return it.
- Otherwise return the non-`nil` side so the ancestor chain keeps propagating the found node upward.

## 中文詳解
1. 從根節點做 DFS，若遇到 `nil`、`p`、`q`，直接回傳該指標。
2. 對左右子樹各自遞迴搜尋，取得 `left`、`right` 結果。
3. 若左右皆非 `nil`，表示 `p`、`q` 分別在左右子樹，當前節點就是 LCA。
4. 若只有一側非 `nil`，則回傳那一側，代表目前子樹內找到的節點往上傳。
5. 題目保證 `p`、`q` 都存在，因此最終一定會找到答案。

## Complexity
- Time: `O(n)` because every node may be visited once.
- Space: `O(h)` recursion stack, where `h` is the tree height.

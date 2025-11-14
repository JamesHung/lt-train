## Original Prompt
Given the root of a binary tree, return whether it is height-balanced (for every node, the height difference between left and right subtrees is at most 1).

## English Explanation
Use a post-order traversal that computes the height of each subtree while checking the balance constraint. The helper returns the subtree height but also flips a shared boolean to `false` if any node violates the balance rule (difference exceeds 1). Once imbalance is detected we can short-circuit by skipping further work in deeper calls.

## 中文詳解
1. 透過遞迴後序遍歷計算每個節點的左、右子樹高度。
2. 若某節點左右高度差大於 1，就標記為不平衡並可提早結束後續計算。
3. 回傳給上一層時，提供目前節點的高度 `max(leftHeight, rightHeight) + 1`。
4. 最終根節點完成後，依據標記判斷整棵樹是否平衡。

## Complexity
- Time: `O(n)` since each node is visited once.
- Space: `O(h)` recursion depth (`h` is tree height).

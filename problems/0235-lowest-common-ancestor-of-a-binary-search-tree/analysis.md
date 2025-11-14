## Original Prompt
Given a Binary Search Tree root and two nodes `p` and `q`, return their lowest common ancestor—the lowest node that has both as descendants (a node can be a descendant of itself).

## English Explanation
Because of the BST ordering, you can determine where the split between `p` and `q` occurs by comparing their values to the current node’s value. Starting at the root:
- If both values are smaller, the LCA must live in the left subtree, so move left.
- If both are larger, move right.
- Otherwise the current node separates the two (or equals one of them), making it the LCA.
This single pass stops as soon as the split is found.

## 中文詳解
1. 從根節點開始，觀察 `p.Val` 與 `q.Val` 相對於目前節點的大小。
2. 若兩者都小於目前節點值，代表 LCA 位於左子樹，往左走即可。
3. 若兩者都大於目前節點值，則往右子樹移動。
4. 一旦碰到節點位在兩者之間，或等於其中之一，該節點就是最低共同祖先；因為 BST 的排序保證在此節點上下分別含有 `p`、`q`。
5. 由於題目保證 `p`、`q` 皆存在，過程中一定會找到答案。

## Complexity
- Time: `O(h)` where `h` is the height of the BST (worst case `O(n)` if skewed).
- Space: `O(1)` iterative traversal.

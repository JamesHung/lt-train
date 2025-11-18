## Original Prompt
Given the root of a binary tree, determine if it is a valid Binary Search Tree (BST). A BST requires every node's left subtree to contain strictly smaller values and the right subtree to contain strictly larger values; both subtrees must also be BSTs. Example inputs: `[2,1,3] -> true`, `[5,1,4,null,null,3,6] -> false`. Constraints: 1 ≤ nodes ≤ 10^4, each value in [−2^31, 2^31 − 1].

## English Explanation
We validate recursively while carrying the valid range for each node. Initially a node may be between (−∞, +∞). When we go left, we limit the upper bound to the current node value; when we go right, we limit the lower bound. If any node falls outside its allowed interval we immediately return false. This enforces the strict inequality requirement across the entire tree, not just parent-child relationships. Using 64-bit bounds prevents overflow when values are ±2^31.

## 中文詳解
DFS 時順便傳遞「這個節點允許的最小/最大值」。根節點一開始允許 (−∞, +∞)。往左子樹遞迴時，把上限改成父節點值；往右子樹就把下限改成父節點值。只要任何節點不在 (low, high) 內就不是 BST，立即回傳 false。這樣可以抓到深層節點違反 BST 規則的情況，並且用 int64 來存上下界避免靠近 ±2^31 時溢位。

## Complexity
- Time: `O(n)` because we visit each node exactly once.
- Space: `O(h)` recursion stack, where `h` is tree height (worst case `O(n)` for skewed trees).

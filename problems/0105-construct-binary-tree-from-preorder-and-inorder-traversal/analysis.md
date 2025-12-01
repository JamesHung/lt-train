## Original Prompt
Given two integer arrays `preorder` and `inorder` that represent the preorder and inorder traversals of the same binary tree, construct and return that binary tree. Constraints: `1 <= preorder.length <= 3000`, `preorder.length == inorder.length`, values are unique and appear in both arrays.

## English Explanation
Preorder visits `root, left, right`. The first element of the current preorder slice is always the root of the current subtree. In inorder, elements left of that root index belong to the left subtree, and elements right of it belong to the right subtree. Build a map from value to inorder index for O(1) splits. Maintain a global `preIdx` pointing at the next root in preorder. Recursively:
1) Take `preorder[preIdx]` as root, increment `preIdx`.
2) Find its inorder index `mid`.
3) Recurse on `[l, mid-1]` for left, then `[mid+1, r]` for right.
This consumes each node once. A `debugBuildTree` flag can log the range per node.

## 中文詳解
1. Preorder 的第一個元素一定是目前子樹的根。根在 inorder 中的位置 `mid` 左邊就是左子樹，右邊是右子樹。
2. 先把 inorder 建成「值 → 位置」的 map，可 O(1) 找到 `mid`。
3. 用全域指標 `preIdx` 依序取 preorder[preIdx] 當根，`preIdx++` 後遞迴建左子樹（區間 `[l, mid-1]`），再建右子樹（`[mid+1, r]`）。
4. 每個節點訪問一次，`debugBuildTree` 可列印建樹時的 inorder 範圍，方便除錯。

## Complexity
- Time: `O(n)` for `n` nodes.
- Space: `O(n)` for the hashmap and recursion stack.

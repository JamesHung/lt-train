## Original Prompt
Given the root of a binary tree, invert it (swap left and right children recursively) and return the root. The tree has up to 100 nodes with values between `-100` and `100`. Return an empty tree when the input is empty.

## English Explanation
Perform a depth-first traversal. For each node, recursively invert its left and right subtrees, then swap them. Returning the current node lets the recursion rebuild the inverted tree. This can be implemented either recursively or iteratively with a queue/stack; recursion is concise and safe here because the tree depth is limited.

## 中文詳解
1. 若節點為 `nil` 直接回傳 `nil` 表示子樹已到底。
2. 遞迴處理右子樹與左子樹，並將結果分別指派回原本的左、右指標（順序調換）。
3. 每一層完成後就完成該節點的鏡像，最後回傳根節點即可得到整棵反轉後的樹。
4. 由於節點數最多 100，遞迴深度安全；若擔心堆疊，可改用 BFS/DFS 疊代方式。

## Complexity
- Time: `O(n)` because each node is visited once.
- Space: `O(h)` for recursion stack where `h` is tree height (`<= n` but typically smaller).

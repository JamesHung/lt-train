## Original Prompt
Given the root of a binary tree, imagine yourself standing on the right side of it. Return the values of the nodes you can see ordered from top to bottom.

## English Explanation
Do a breadth-first traversal and process one level at a time. Track `levelSize` so we know how many nodes belong to the current depth. While draining those nodes from the queue, keep updating `lastVal` with the most recently visited node. After the loop for that level finishes, `lastVal` is the rightmost node for that depth, so append it to the result. Enqueue children left-to-right so ordering is consistent; we only care about the final node of each level.

## 中文詳解
1. 以 BFS 分層遍歷，隊列初始放入根節點；若樹為空直接回傳 `nil`。
2. 取出當層節點數 `levelSize`，逐一彈出並記錄當前節點值到 `lastVal`。
3. 將每個節點的左、右子節點（若存在）依序加入隊列，維持層序順序。
4. 當前層處理完後，`lastVal` 就是從右側能看到的節點，加入答案切片。
5. 重複直到隊列為空。

## Complexity
- Time: `O(n)` because each node is visited once.
- Space: `O(n)` for the queue in the widest level.

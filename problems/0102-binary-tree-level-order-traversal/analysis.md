## Original Prompt
Given the root of a binary tree, return the level order traversal (left to right, level by level) of all node values.

**Example 1**
- Input: `[3,9,20,null,null,15,7]`
- Output: `[[3],[9,20],[15,7]]`

**Example 2**
- Input: `[1]`
- Output: `[[1]]`

**Example 3**
- Input: `[]`
- Output: `[]`

Constraints:
- `0 <= number of nodes <= 2000`
- `-1000 <= Node.val <= 1000`

## English Explanation
Perform a breadth-first search. Place the root into a queue and process nodes level by level: for each iteration record the current queue length, pop exactly that many nodes, append their values to the current level slice, and enqueue their children. Repeat until the queue is empty. Handling a `nil` root upfront lets us return an empty result. The queue guarantees nodes are visited in level order.

## 中文詳解
使用 BFS。先把根節點放進佇列，每次迴圈先記錄佇列長度 `levelSize`，表示這層有多少節點。接著取出這些節點，將值加進 `level` 陣列，並把左右子節點（若存在）加入佇列。當這層處理完就把 `level` append 到答案中，如此反覆直到佇列空。若根節點為空，直接回傳空陣列即可。

## Complexity
- Time: `O(n)` where `n` is the number of nodes; each node is enqueued and dequeued once.
- Space: `O(n)` in the worst case to hold the queue (or the last level).


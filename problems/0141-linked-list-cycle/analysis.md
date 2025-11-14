## Original Prompt
Given the head of a singly linked list, return `true` if the list contains a cycle; otherwise return `false`. A cycle exists when following `next` references eventually revisits a node. The index `pos` indicates where the tail connects but is only informational.

## English Explanation
Use Floyd’s tortoise and hare pointers. Advance `slow` one node at a time and `fast` two nodes. If the list has a cycle, the two pointers will eventually meet inside the loop. If `fast` or `fast.Next` becomes `nil`, we reached the end and there is no cycle. This algorithm uses constant extra memory.

## 中文詳解
1. 設定兩個指標：`slow` 每次走一步、`fast` 每次走兩步。
2. 在迴圈中同時推進；若其中任一指標遇到 `nil`，代表串列結束，沒有環。
3. 若某次迭代 `slow == fast`，表示兩者在環內相遇，串列確實存在環。
4. 全程僅用固定數量指標，因此額外空間為 `O(1)`。

## Complexity
- Time: `O(n)` because each pointer traverses at most the length of the list plus cycle.
- Space: `O(1)` extra memory.

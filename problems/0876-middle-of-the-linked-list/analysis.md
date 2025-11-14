## Original Prompt
Given the head of a singly linked list, return its middle node. If there are two middles, return the second one.

## English Explanation
Maintain two pointers: `slow` advances one step at a time while `fast` advances two. When `fast` reaches the end (or is about to pass it), `slow` sits at the middle. Because `fast` moves twice as quickly, `slow` will land on the second middle when the length is even, satisfying the requirement.

## 中文詳解
1. 設定 `slow = head`、`fast = head`。
2. 迴圈條件為 `fast != nil && fast.Next != nil`：每回合 `slow` 前進一格、`fast` 前進兩格。
3. 當 `fast` 抵達尾端或 `fast.Next` 為空時，`slow` 即停在中間節點；若節點數為偶數，它會落在第二個中間。
4. 回傳 `slow`。

## Complexity
- Time: `O(n)` single traversal.
- Space: `O(1)` extra memory.

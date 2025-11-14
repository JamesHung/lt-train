## Original Prompt
Given the head of a singly linked list, reverse the list and return the new head. Follow-up: implement both iterative and recursive approaches.

## English Explanation
Iterative: walk the list with two pointers, `prev` and `curr`. At each step preserve `next := curr.Next`, redirect `curr.Next` to `prev`, then advance both pointers. When `curr` becomes `nil`, `prev` holds the new head. Recursive: recurse to the end so that `reverseListRecursive(head.Next)` gives the reversed tail; then flip the next pointer via `head.Next.Next = head` and set `head.Next = nil` to avoid cycles.

## 中文詳解
1. **迭代法**：準備 `prev = nil`、`curr = head`。每次把 `curr.Next` 暫存後改指向 `prev`，再往前移動 `prev`、`curr`。當 `curr` 為 `nil` 時，`prev` 就是新頭節點。
2. **遞迴法**：先遞迴到尾端得到已反轉的子串，新尾端 `head.Next` 會變成當前節點的前一個。把 `head.Next.Next` 指回 `head`，最後把 `head.Next = nil`，避免形成環。

## Complexity
- Time: `O(n)` for both approaches.
- Space: Iterative uses `O(1)` extra memory; recursive needs `O(n)` call stack depth.

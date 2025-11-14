## Original Prompt
You are given the heads of two sorted linked lists `list1` and `list2`. Merge them into one sorted list by splicing the existing nodes together and return the head. Each list length is within `[0, 50]`, values lie between `-100` and `100`, and both lists are sorted in non-decreasing order.

## English Explanation
Use a dummy head and a pointer `cur` that always represents the tail of the merged list. Compare the current nodes from both lists, append the smaller one to `cur.Next`, then advance both `cur` and the list from which the node was taken. When either list runs out, append the remaining nodes from the other list. Return `dummy.Next` for the merged head.

## 中文詳解
1. 建立一個虛擬節點 `dummy` 以及指向它的指標 `cur`，方便統一處理頭節點。
2. 同步遍歷 `list1` 與 `list2`，每次比較目前節點的值，將較小者接在 `cur` 後方，並往前推進該串列。
3. 無論哪一條串列先耗盡，直接把另一條剩下的節點整段掛到 `cur.Next` 上。
4. 回傳 `dummy.Next` 即為合併後的排序串列；整個過程只遍歷各節點一次。

## Complexity
- Time: `O(m + n)` because every node of both lists is inspected exactly once.
- Space: `O(1)` extra since the merge is done in place using existing nodes.

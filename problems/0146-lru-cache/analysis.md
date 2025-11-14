## Original Prompt
Design an `LRUCache` supporting `get` and `put` in `O(1)` average time with capacity up to 3000 and up to 2e5 operations.

## English Explanation
Combine a hash map and a doubly linked list. The map stores key → node pointers for `O(1)` access. The list keeps usage order with the most recent at the front. `get` checks the map; if present, move that node to the front and return its value. `put` updates and moves when the key exists; otherwise remove the tail node (least recently used) when capacity is exceeded, insert a new node right after the head, and record it in the map.

## 中文詳解
1. 使用 map 保存 `key -> 雙向鏈結點`，讓查找或更新都能 `O(1)`。
2. 另建一個帶有虛擬 head/tail 的雙向鏈結串列，head 之後代表最新使用，tail 之前代表最久未使用。
3. `get`：若 key 存在，就把該節點移到 head 後面並回傳 value；若不存在回傳 `-1`。
4. `put`：若 key 已存在，直接更新 value 並移到前方；若不存在且容量已滿，先從 tail 後方移除節點並同時刪掉 map 中的項目，再插入新節點到 head 之後。

## Complexity
- `get` / `put`: `O(1)` 平均時間，因 map 和鏈結串列操作皆為常數。
- 空間：`O(capacity)`，存儲所有節點與 map。

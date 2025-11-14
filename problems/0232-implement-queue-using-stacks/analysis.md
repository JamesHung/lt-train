## Original Prompt
Implement a queue with standard queue operations (`push`, `pop`, `peek`, `empty`) using only stack operations. Achieve amortized `O(1)` per operation.

## English Explanation
Maintain two stacks: `in` for new pushes and `out` for serving pops/peeks. Push operations append to `in` directly. When `pop` or `peek` needs data but `out` is empty, pour all items from `in` into `out`, reversing order so the oldest element ends up on top of `out`. Subsequent pops/peeks use `out` until it empties again. Each element moves at most once from `in` to `out`, giving amortized `O(1)` time while keeping constant space aside from the stored items.

## 中文詳解
1. 設兩個堆疊：`in` 專門處理 `push`，`out` 提供 `pop/peek`。
2. `push(x)` 只需把 `x` 放進 `in`，成本為 `O(1)`。
3. `pop/peek` 時若 `out` 為空，就把 `in` 的元素逐一彈出並推入 `out`，完成順序反轉後即可取得隊首。
4. 每個元素最多被搬運一次（`in -> out`），因此長期平均成本仍是 `O(1)`；額外空間只來自錄入的元素。

## Complexity
- Amortized Time: `O(1)` per operation; worst-case `O(n)` occurs only during a transfer.
- Space: `O(n)` for storing queue elements across two stacks.

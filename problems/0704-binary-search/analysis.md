## Original Prompt
Given a sorted integer array `nums` and an integer `target`, return the index if `target` is found; otherwise return `-1`. You must design an algorithm with `O(log n)` runtime.

## English Explanation
Use binary search to continuously halve the search space. Track two pointers, `lo` and `hi`, as the inclusive bounds of the area still under consideration. Compute `mid`, compare `nums[mid]` to `target`, and drop the left or right half accordingly. When the value matches, return `mid`; if the pointers cross without a match, the value is absent so return `-1`.

## 中文詳解
1. 維護左右邊界 `lo` 與 `hi`，初始為陣列的首尾索引。
2. 每回合計算 `mid = lo + (hi - lo)/2`，避免溢位的同時鎖定中間位置。
3. 若 `nums[mid] == target` 直接回傳 `mid`。
4. 若 `nums[mid] < target`，代表答案只可能在右半邊，將 `lo` 移到 `mid + 1`；反之則把 `hi` 移到 `mid - 1`。
5. 若左右邊界交錯仍未找到，代表陣列中沒有 `target`，回傳 `-1`。

## Complexity
- Time: `O(log n)` because each step discards half of the remaining range.
- Space: `O(1)` auxiliary memory.

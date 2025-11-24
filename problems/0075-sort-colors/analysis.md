## Original Prompt
Given an array `nums` with `n` objects colored red, white, or blue (encoded as `0, 1, 2`), sort them in-place so same colors are adjacent in order red → white → blue. Do not use the library sort. `1 <= n <= 300`. Follow-up asks for a one-pass, constant-space algorithm.

## English Explanation
Use the Dutch National Flag approach with three pointers: `low` (next spot for 0), `mid` (current cursor), and `high` (next spot for 2 from the end). Walk while `mid <= high`: when `nums[mid]` is 0, swap it with `nums[low]` and advance both `low` and `mid`; when it is 1, just advance `mid`; when it is 2, swap with `nums[high]` and move `high` left (do not advance `mid` because the swapped-in value needs evaluation). Every element is touched at most once, giving a one-pass, constant-space sort.

## 中文詳解
使用荷蘭國旗演算法：維護三個指標 `low`（放 0 的下個位置）、`mid`（當前掃描位置）、`high`（放 2 的下個位置，自右往左縮）。迴圈條件 `mid <= high`：
1) `nums[mid] == 0` 時，與 `low` 位置交換，`low++`、`mid++`，因為換出的元素已經處理過。
2) `nums[mid] == 1` 時，已在正確區段，`mid++` 即可。
3) `nums[mid] == 2` 時，與 `high` 位置交換，`high--`，但 `mid` 不動，需檢查被換進來的值。
如此單趟掃描就能完成排序，且只用常數額外空間。

## Complexity
- Time: `O(n)` — 每個元素最多被交換與讀取一次。
- Space: `O(1)` — 僅使用固定數量指標。

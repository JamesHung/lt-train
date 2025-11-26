## Original Prompt
Given an `m x n` matrix, return all elements in spiral order. Constraints: `1 <= m, n <= 10`, values in `[-100, 100]`.

## English Explanation
Maintain four boundaries: `top`, `bottom`, `left`, `right`. Repeatedly traverse:
1) top row left→right, then `top++`
2) right column top→bottom, then `right--`
3) bottom row right→left (if `top <= bottom`), then `bottom--`
4) left column bottom→top (if `left <= right`), then `left++`
Stop when boundaries cross. This walks the matrix layer by layer in spiral order.

## 中文詳解
用四個邊界控制外圈：`top/bottom/left/right`。每輪依序：
1) 走上邊一排（左到右），上邊界下移
2) 走右邊一欄（上到下），右邊界左移
3) 若仍有未交錯，走下邊一排（右到左），下邊界上移
4) 若仍有未交錯，走左邊一欄（下到上），左邊界右移
當上下或左右邊界交錯就停止。這樣每次縮小一圈，能按照螺旋順序取完所有元素。

## Complexity
- Time: `O(m*n)` — each cell visited once.
- Space: `O(1)` extra besides the output slice.

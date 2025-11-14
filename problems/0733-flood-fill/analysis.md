## Original Prompt
Given an `m x n` grid `image`, repaint all pixels connected 4-directionally to `image[sr][sc]` that share the same original color so they become `color`. Return the modified image.

## English Explanation
Keep the original color and run a DFS/BFS from `(sr, sc)`. Whenever we pop a pixel that still has the original color, recolor it and inspect its four neighbors. Any neighbor that is inside the grid and still has the starting color is added to the worklist. If the target color equals the original color, nothing needs to change, so return early.

## 中文詳解
1. 先記錄起點 `(sr, sc)` 的原始顏色 `origin`，若 `origin == color` 則直接回傳。
2. 使用堆疊或佇列儲存待處理的座標，初始放入 `(sr, sc)`。
3. 每次取出一格，若顏色仍是 `origin`，就把它塗成 `color`，並檢查上下左右四個鄰居。
4. 只要鄰居在範圍內且顏色也是 `origin`，就加入堆疊持續處理，如此可以覆蓋所有連通區塊。
5. 當堆疊清空時代表塗色完成，回傳 `image`。

## Complexity
- Time: `O(mn)` in the worst case when the entire grid is recolored.
- Space: `O(mn)` in the worst case for the stack/queue storing boundary pixels (or `O(1)` extra beyond the recursion/stack if the region is small).

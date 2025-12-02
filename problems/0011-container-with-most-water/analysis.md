## Original Prompt
Given an array `height` where the `i`-th line has endpoints `(i, 0)` and `(i, height[i])`, find two lines that, with the x-axis, form a container holding the most water. Return that maximum area. Constraints: `2 <= n <= 1e5`, `0 <= height[i] <= 1e4`.

## English Explanation
Use two pointers at the ends. The area is limited by the shorter side: `min(height[l], height[r]) * (r - l)`. To potentially improve area, move the shorter side inward—moving the taller side cannot create a larger height than the current min while also reducing width. Repeat until pointers meet, tracking the best area. A `debugContainer` flag can log each step.

## 中文詳解
1. 夾兩端指標 `l=0`, `r=n-1`，每次容積 = `min(height[l], height[r]) * (r-l)`。
2. 想增加容積，只能嘗試提高較矮的那一側；因此如果 `height[l] < height[r]` 就 `l++`，否則 `r--`。
3. 迭代直到兩指標相遇，過程中記錄最大值。
4. `debugContainer` 可輸出每步的指標、當前高度與最大值，方便檢查。

## Complexity
- Time: `O(n)`
- Space: `O(1)`

# Median of Two Sorted Arrays

## Problem Statement
- Input: Two sorted arrays `nums1` (length `m`) and `nums2` (length `n`), where `0 <= m, n <= 1000`, `1 <= m + n <= 2000`, and each value is between `-10^6` and `10^6`.
- Output: The median of the combined elements, returned as a floating-point value.
- Requirement: Overall runtime must be `O(log(m + n))`.

## Approach
- Always binary search the shorter array so the search space stays minimal.
- For a tentative partition `partition1` in `nums1`, compute the matching partition `partition2` in `nums2` that ensures the left halves together hold `(m+n+1)/2` elements.
- Compare the bordering values (`maxLeft1`, `minRight1`, `maxLeft2`, `minRight2`). When both left maxima are ≤ the opposite right minima, the correct partition is found and the median is either the max of the left sides (odd total) or the average of max-left/min-right (even total).
- Adjust the binary search bounds depending on which side violates the ordering.

## Intuition (中文)
1. 把兩個陣列想像成被切成左半與右半的交錯分割；正確分割的特性是「左邊所有元素都 ≤ 右邊所有元素」。
2. 為了在 O(log(min(m,n))) 時間完成，只對較短的陣列做二分搜尋，同時計算另一個陣列的對應切點，確保左半部的元素數量正好等於 `(m+n+1)/2`。
3. 如果左邊的最大值仍然比右邊的最小值大，代表切過頭，要縮小或擴大搜尋區間直到符合條件；一旦符合，根據總長度是奇數或偶數決定中位數。

## 詳解 (中文)
- **目標**：維持一個分割位置，使得兩個陣列左半邊的元素個數加總等於 `(m+n+1)/2`，而且左邊的最大值不會大於右邊的最小值。
- **分割方式**：在較短陣列 `nums1` 上選一個 `partition1`，另一個陣列的切點就固定是 `partition2 = totalLeft - partition1`。這樣左右兩半的元素數量總是正確。
- **邊界處理**：若切點落在陣列頭或尾，就用 `math.MinInt`/`math.MaxInt` 當作不存在的左/右元素，簡化比較邏輯。
- **判斷條件**：當 `maxLeft1 <= minRight2` 且 `maxLeft2 <= minRight1` 時，代表兩邊分割都符合遞增順序，可以直接回傳中位數；否則調整二分搜尋上下界繼續找。
- **中位數計算**：總長度奇數時回傳左半邊最大值；偶數時回傳左半最大與右半最小的平均。

## Complexity
- Time: `O(log(min(m, n)))` due to binary search on the shorter array.
- Space: `O(1)` extra space beyond a few pointers/integers.

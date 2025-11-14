## Original Prompt
Given an array `prices` where `prices[i]` is the stock price on day `i`, buy once and sell once later to maximize profit. Return the maximum profit, or `0` if no profit is possible. Constraints: `1 <= len(prices) <= 10^5`, `0 <= prices[i] <= 10^4`.

## English Explanation
Track the lowest price seen so far while scanning left to right. At each day, compute the profit if you bought at the tracked minimum and sold today, and update the best profit. If the current price is lower than the minimum, refresh the minimum. This greedy scan guarantees the best buy-before-sell pair.

## 中文詳解
1. 由左至右掃描 `prices`，同時維護目前為止遇過的最低價格 `minPrice`。
2. 每當讀到新的價格 `price`，先計算 `price - minPrice` 作為當天賣出的潛在獲利，若高於目前最佳值就更新。
3. 若 `price < minPrice`，代表找到更好的買入點，立即更新 `minPrice`。
4. 掃描結束後的最佳獲利即為答案；若始終沒有正獲利，自然保持為 `0`。

## Complexity
- Time: `O(n)` because the array is scanned once.
- Space: `O(1)` extra.

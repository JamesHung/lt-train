## Original Prompt
You climb a staircase with `n` steps and can take either 1 or 2 steps at a time. Return how many distinct ways exist to reach the top.

## English Explanation
The recurrence is identical to Fibonacci: the ways to reach step `i` equal the sum of ways to reach `i-1` (take one more step) and `i-2` (take two steps). Iteratively accumulate this using two rolling variables to keep `O(1)` space. Handle the `n <= 2` bases directly.

## 中文詳解
1. 若 `n` 為 1 或 2，答案分別是 1、2，直接回傳。
2. 之後每一層的走法 = 前一層 + 前兩層。可用兩個變數 `prev1`、`prev2` 記錄即可。
因為從第 i 階往回看，只能從兩個位置跨上來：要嘛站在第 i-1 階再走 1 步，要嘛站在第 i-2 階一次跨 2 步。能抵達第 i-1 的走法共有 f(i-1) 種、能抵達第 i-2 的走法共有 f(i-2) 種，而且這兩批走法互斥（最後一步不同），把它們相加就是抵達第 i 階的總
  數 f(i) = f(i-1) + f(i-2)。所以用兩個滾動變數記錄前兩階的答案就能迭代算出後面的值
3. 從第 3 階一路迭代到 `n`，每回合更新 `curr = prev1 + prev2`，再把 `prev2 = prev1`、`prev1 = curr`。
4. 迴圈結束後 `prev1` 即為總走法數。

## Complexity
- Time: `O(n)` iterations.
- Space: `O(1)` extra storage.



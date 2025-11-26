## Original Prompt
Given an integer array `nums`, return `true` if it can be partitioned into two subsets whose sums are equal. Otherwise, return `false`. Constraints: `1 <= len(nums) <= 200`, `1 <= nums[i] <= 100`.

## English Explanation
This is a subset-sum problem: if total sum is odd, impossible. Otherwise target is `total/2`; we just need to know if any subset sums to target. Use 1D DP where `dp[s]=true` means a subset sums to `s`. Initialize `dp[0]=true`. For each `num`, iterate `s` descending from `target` to `num`, and set `dp[s] = dp[s] || dp[s-num]`. Descending order prevents reusing the same `num` multiple times. After processing all numbers, `dp[target]` reveals if a balanced partition exists.

## 中文詳解
把問題轉成「是否存在子集合和為 `total/2`」。若總和為奇數直接回傳 `false`。否則設定目標 `target = total/2`，用一維布林 DP：`dp[s]` 代表有沒有子集合可以組出和 `s`。初始 `dp[0]=true`。對每個 `num`，從 `target` 逆向迴圈到 `num`，若 `dp[s-num]` 為真則設 `dp[s]=true`，這樣每個元素只被用一次。處理完所有數字，查看 `dp[target]` 即可判斷是否能等分。

## Complexity
- Time: `O(n * target)` where `target = total/2` (worst-case `n*10000`).
- Space: `O(target)` for the DP array.

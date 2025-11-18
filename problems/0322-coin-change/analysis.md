## Original Prompt
Given coin denominations and a target amount, compute the minimum number of coins needed to make the amount, or return `-1` if it's impossible. Each coin can be used infinitely many times.

**Example 1**
- Input: `coins = [1,2,5], amount = 11`
- Output: `3`

**Example 2**
- Input: `coins = [2], amount = 3`
- Output: `-1`

**Example 3**
- Input: `coins = [1], amount = 0`
- Output: `0`

Constraints:
- `1 <= coins.length <= 12`
- `1 <= coins[i] <= 2^31 - 1`
- `0 <= amount <= 10^4`

## English Explanation
Use bottom-up dynamic programming. Let `dp[i]` be the fewest coins to make sum `i`. Initialize `dp[0] = 0` and fill the rest with infinity. For each `i` from `1` to `amount`, iterate over every coin; if the coin can fit (i.e., `coin <= i` and `dp[i-coin]` is reachable), update `dp[i] = min(dp[i], dp[i-coin]+1)`. After filling the table, if `dp[amount]` is still infinity, return `-1`, otherwise return it.

## 中文詳解
採用自底向上的 DP。定義 `dp[i]` 為湊成金額 `i` 所需的最少硬幣數。初始設 `dp[0]=0`，其他設成無限大。對每個金額 `i=1..amount`，檢查所有硬幣 `coin`：若 `coin <= i` 且 `dp[i-coin]` 可以到達，就更新 `dp[i] = min(dp[i], dp[i-coin]+1)`。表格填完後，若 `dp[amount]` 仍是無限大表示無法湊出，回傳 `-1`；否則回傳 `dp[amount]`。

###

假設我們已經算到 i = 7，此時 dp 的狀態大致為：

  i:      0 1 2 3 4 5 6 7
  dp[i]:  0 1 1 2 2 1 2 inf

  - 先看 coin = 1：coin <= i 成立。dp[7-1] = dp[6] = 2，不是無限大，表示 6 塊錢可以湊得出。dp[6] + 1 = 3，目前 dp[7] 是 inf，所以 3 < inf，我們可以更新 dp[7] = 3（代表 7 = 6 + 1，用三枚：1+1+5 或 1+1+1+... 等）。
  - 接著看 coin = 2：coin <= 7，dp[5] = 1（5 塊可由一枚 5 元湊出），dp[5] + 1 = 2，比目前的 dp[7] = 3 更小，所以更新 dp[7] = 2（表示 7 = 5 + 2，用兩枚硬幣）。
  - coin = 5：coin <= 7，dp[2] = 1，1 + 1 = 2，跟現有的 dp[7] 相同，不會更新。
  - coin = 11 不存在，所以略過。

  最後 dp[7] = 2，表示 7 塊最少用兩枚硬幣。這整段判斷就是確保：1) 這個硬幣能用；2) 剩下的金額曾經被算過；3) 用它可以讓 dp[i] 更小，才更新最優解。

## Complexity
- Time: `O(amount * len(coins))` since we process each coin for each amount up to the target.
- Space: `O(amount)` for the DP array.


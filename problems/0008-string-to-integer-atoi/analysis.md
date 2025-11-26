## Original Prompt
Implement `myAtoi(string s)` to convert a string into a 32-bit signed integer with these rules: skip leading whitespace; read optional sign; consume consecutive digits (ignoring leading zeros) until a non-digit; if no digits, return `0`; clamp overflow to `[-2^31, 2^31-1]`. Examples: `"42" -> 42`, `"   -042" -> -42`, `"1337c0d3" -> 1337`, `"0-1" -> 0`, `"words and 987" -> 0`. Constraints: `0 <= s.length <= 200`, characters may be letters, digits, space, `+`, `-`, `.`.

## English Explanation
Scan the string once. Skip leading spaces, read a sign, then accumulate digits. Track the current value as a non-negative `val` and a `sign` (`1` or `-1`). Before adding a new digit, check whether `val*10+digit` would overflow the 32-bit signed range. For positives, if `val > (INT_MAX-digit)/10`, clamp to `INT_MAX`; for negatives, compare to `INT_MAX+1` (absolute value of `INT_MIN`). Stop at the first non-digit and return `sign * val`, applying the clamp if triggered.

## 中文詳解
一次掃描字串：先跳過前導空白，再讀取正負號，之後連續讀數字。用 `val` 保存目前累積值、`sign` 表示正負。在加入新數字前做溢位判斷：正數若 `val > (INT_MAX-digit)/10`，直接回傳 `INT_MAX`；負數用 `INT_MAX+1` 作為可接受的絕對值上限，超過就回傳 `INT_MIN`。遇到第一個非數字就停止，回傳 `sign * val`（若已溢位則回傳邊界值）。

## Complexity
- Time: `O(n)` for scanning the string once.
- Space: `O(1)` extra.

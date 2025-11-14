## Original Prompt
Given binary strings `a` and `b`, return their sum as a binary string.

## English Explanation
Walk both strings from right to left, mimicking grade-school addition on base 2. Keep a `carry`, add the digits plus carry, push `sum % 2` to an output buffer, and update `carry = sum / 2`. Continue until both strings and the carry are consumed, then reverse the buffer to form the final answer.

## 中文詳解
1. 從字串尾端開始，分別取出 `a`、`b` 的當前位數。
2. 計算 `sum = carry + digitA + digitB`，把 `sum % 2` 轉成字元加入結果，`carry = sum / 2`。
3. 兩個字串都處理完後若仍有進位就再補上一位。
4. 因為結果是由最低位開始加入，最後記得把整個切片反轉再轉成字串。

## Complexity
- Time: `O(max(len(a), len(b)))`.
- Space: `O(max(len(a), len(b)))` for the output buffer.

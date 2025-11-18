## Original Prompt
Evaluate a Reverse Polish Notation expression described by `tokens`, where tokens are either numbers or one of `+`, `-`, `*`, `/`. Division truncates toward zero and the input is guaranteed valid.

**Example 1**
- Input: `["2","1","+","3","*"]`
- Output: `9`

**Example 2**
- Input: `["4","13","5","/","+"]`
- Output: `6`

**Example 3**
- Input: `["10","6","9","3","+","-11","*","/","*","17","+","5","+"]`
- Output: `22`

Constraints:
- `1 <= tokens.length <= 10^4`
- `tokens[i]` is either an operator or an integer in `[-200, 200]`.

## English Explanation
Use a stack. Iterate through `tokens` and push numbers as integers. When you see an operator, pop the top two numbers, apply the operation in order (`a op b`, where `b` is the last popped), and push the result back. Because division truncates toward zero by default in Go on integers, no special handling is needed. After processing every token, the stack contains one value which is the answer.

## 中文詳解
用堆疊來處理。遍歷 `tokens`，遇到數字就轉成整數後推入堆疊；遇到運算符號就從堆疊取出兩個數，依照順序計算（先彈出的是右操作數 `b`，再彈出的是左操作數 `a`），把結果再推回堆疊。Go 裡的整數除法預設就是向零截斷，因此不需要額外處理。所有 token 處理完後，堆疊上只會剩下最終結果。

## Complexity
- Time: `O(n)` — each token is processed once.
- Space: `O(n)` in the worst case for the stack (all tokens are numbers until the end).


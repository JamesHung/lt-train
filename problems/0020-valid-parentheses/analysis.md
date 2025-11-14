## Original Prompt
Given a string `s` consisting solely of `'('`, `')'`, `'{'`, `'}'`, `'['`, and `']'`, determine whether it is valid. A string is valid when every opening bracket is closed by the same type, closes in the correct order, and every closing bracket has a matching opener. Constraints: `1 <= s.length <= 10^4`.

## English Explanation
Walk through the string while using a stack that stores opening brackets. When an opener is seen, push it. When a closer is seen, the stack must be non-empty and its top must be the matching opener; otherwise the string is invalid. After processing the entire string, the stack must be empty to confirm all brackets were closed.

## 中文詳解
1. 準備一個堆疊，遇到左括號就推入（紀錄當下等待匹配的型別）。
2. 遇到右括號時，若堆疊為空或頂端不是對應的左括號，就代表順序或型別錯誤，直接回傳 `false`。
3. 若匹配成功就將堆疊頂端彈出，繼續往後掃描。
4. 迴圈結束後檢查堆疊是否為空；若仍有元素，表示尚未關閉的括號存在，故回傳 `false`，否則為 `true`。

## Complexity
- Time: `O(n)` because each character is pushed/popped at most once.
- Space: `O(n)` in the worst case when the string is all open brackets.

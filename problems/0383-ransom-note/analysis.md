## Original Prompt
Given strings `ransomNote` and `magazine`, determine if `ransomNote` can be constructed using the characters in `magazine`. Each letter in `magazine` can be used at most once.

## English Explanation
Because we only deal with lowercase letters, keep a frequency array of size 26. Count every character in `magazine`, then iterate through `ransomNote`, decrementing the corresponding count. If any count goes negative, the magazine does not provide enough of that letter, so return `false`. Otherwise all letters are available and the construction succeeds.

## 中文詳解
1. 由於只有小寫字母，建立長度 26 的整數陣列紀錄 magazine 中每個字母的次數。
2. 先走訪 `magazine` 把每個字母的計數加一。
3. 再走訪 `ransomNote`，對應的計數減一；若某一格變成負數，代表 magazine 中該字母不足，直接回傳 `false`。
4. 全部字母都能扣完就回傳 `true`。

## Complexity
- Time: `O(m + n)` where `m` and `n` are the lengths of `magazine` and `ransomNote`.
- Space: `O(1)` extra memory (fixed 26-length array).

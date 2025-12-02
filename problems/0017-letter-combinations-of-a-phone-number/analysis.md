## Original Prompt
Given digits `2-9`, return all possible letter combinations based on the phone keypad mapping. Digits length is 1 to 4.

## English Explanation
Backtrack over the digit string. Maintain a `path` byte slice whose length equals `digits`. For position `pos`, iterate the mapped letters, write each into `path[pos]`, and recurse to `pos+1`. When `pos` reaches `len(digits)`, append the string copy to results. Return empty slice for empty input. A `debugLetterCombos` flag can print combos as they’re produced.

## 中文詳解
1. 建立數字到字母的映射（2~9）。若輸入為空字串，直接回傳空 slice。
2. 用 DFS/回溯：`path` 的長度等於輸入長度。處理到第 `pos` 位時，列舉該數字對應的每個字母放到 `path[pos]`，遞迴到下一位。
3. 當 `pos == len(digits)` 時，將 `path` 複製成字串加入答案。
4. `debugLetterCombos` 可在生成組合時輸出，方便檢查。

## Complexity
- Time: `O(3^n * n)` (worst 4 letters per digit, n <= 4).
- Space: `O(n)` recursion + path.

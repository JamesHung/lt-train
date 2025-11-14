## Original Prompt
Given `s` and `numRows`, write the zigzag pattern row by row (e.g., `PAYPALISHIRING` with `numRows = 3` turns into `PAHNAPLSIIGYIR`). Return the final concatenated string. Constraints: `1 <= len(s) <= 1000`, `1 <= numRows <= 1000`.

## English Explanation
Treat each row as a buffer. Walk through the characters while moving down the rows, flip direction when hitting the top or bottom row, and append each character to the current row. Finally concatenate the rows.

## 中文詳解
1. 若 `numRows = 1` 或 `numRows >= len(s)`，Z 字形其實退化為原字串，直接 return。
2. 建立 `numRows` 個 `strings.Builder`，減少字串頻繁拷貝。
3. 用 `curRow` 表示目前所在列，`goingDown` 表示移動方向。每處理一個字元就寫入當前列；當 `curRow` 來到 0 或 `numRows-1` 時，就翻轉方向以模擬「折返」。
4. 迴圈跑完後依序串接每一個 builder 的內容就是答案。

## Complexity
- 時間：`O(n)`，每個字元只訪問一次。
- 空間：`O(n)`，額外記錄分列結果。

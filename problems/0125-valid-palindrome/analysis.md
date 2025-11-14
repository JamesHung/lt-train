## Original Prompt
Given a string `s`, determine whether it is a palindrome after converting uppercase to lowercase and removing all non-alphanumeric characters. Return `true` if the cleaned string reads the same forward and backward. Constraints: `1 <= len(s) <= 2 * 10^5`, `s` only contains printable ASCII.

## English Explanation
Use two pointers from the ends of the string. Move each pointer inward until it stops at an alphanumeric character. Compare the lowercase versions; if they differ, it's not a palindrome. Continue until the pointers cross. This avoids extra memory by not constructing the filtered string explicitly.

## 中文詳解
1. 設定左右指標 `left`, `right` 指向字串首尾。
2. 左指標往右跳過所有非英數字元，右指標往左跳過同類字元。
3. 對於停下來的兩個字元，轉成小寫後比較；若不同，立即回傳 `false`。
4. 若指標交會或交錯仍未出現不一致，就代表清理後字串左右對稱，回傳 `true`。

## Complexity
- Time: `O(n)` because each character is visited at most once by either pointer.
- Space: `O(1)` extra since we only use a few counters.

## Original Prompt
Implement a Trie (prefix tree) that supports insertion, full word search, and prefix search over lowercase strings.

**Example**
- `insert("apple")`
- `search("apple")` → `true`
- `search("app")` → `false`
- `startsWith("app")` → `true`
- `insert("app")`
- `search("app")` → `true`

Constraints:
- `1 <= word.length, prefix.length <= 2000`
- All strings contain lowercase letters.
- At most `3 * 10^4` operations are performed.

## English Explanation
Each node maintains a map from rune to child node plus a boolean flag indicating whether a word ends here. `insert` walks/creates nodes along the characters of the string and marks the final node as a word end. `search` and `startsWith` traverse using the characters; `search` additionally checks the end flag, whereas `startsWith` only ensures the path exists. Using maps keeps implementation simple; arrays of size 26 could also be used for slightly better performance.

## 中文詳解
Trie 的每個節點儲存一個由字元對應到子節點的 map，以及一個布林標記是否有單字在此結尾。`insert` 從根開始，逐字元往下，如果缺少子節點就新建，最後把終點設為 `end = true`。`search` 走完全字後還要確認當前節點的 `end` 為真才算找到；`startsWith` 則只需確認路徑存在。這裡用 map 讓程式簡潔（若需更快也可用固定 26 長度的陣列）。

## Complexity
- Time: `O(L)` per operation where `L` is the word/prefix length.
- Space: worst-case `O(TotalChars)` across all inserted words.


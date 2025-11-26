## Original Prompt
Given a string `s` and a dictionary `wordDict`, return `true` if `s` can be segmented into a space-separated sequence of one or more dictionary words. Words may be reused.

## English Explanation
Use dynamic programming on prefixes. Let `dp[i]` mean `s[:i]` can be segmented. Initialize `dp[0]=true` (empty string). For each `i` where `dp[i]` is true, try to extend using any dictionary word `w`: if `s[i:i+len(w)] == w`, mark `dp[i+len(w)] = true`. To prune work, precompute the maximum word length and only check up to that distance. At the end, `dp[n]` tells if the whole string segments.

## 中文詳解
用前綴 DP：`dp[i]` 表示 `s[0:i]` 可以被字典分割。設定 `dp[0]=true`。遍歷每個 `i`，如果 `dp[i]` 為真，嘗試用字典中的單字去延伸：當 `s[i : i+len(w)] == w` 時，設 `dp[i+len(w)] = true`。為了減少檢查次數，先算出字典中最長單字長度，`i` 之後最多只需要檢查這個長度。最後看 `dp[n]` 是否為真即可。

## Complexity
- Time: `O(n * L)` where `n` is `len(s)` and `L` 是字典單字的平均長度；最壞可視為 `O(n * maxWordLen)`，因為我們只向前看有限長度。
- Space: `O(n)` for the DP array.



  - for i := 0; i < n; i++ { … }：掃描所有前綴結尾位置 i。
  - if !dp[i] { continue }：只從「已知可切分」的前綴 (dp[i]=true) 繼續擴展；不可行的前綴直接跳過，省事。
  - limit := i + maxLen; if limit > n { limit = n }：最多只往後看字典最長單字長度 maxLen，避免檢查多餘長度。
  - for j := i + 1; j <= limit; j++ { … }：嘗試所有以 i 為起點、長度介於 1 到 maxLen 的候選片段 s[i:j]。
  - if _, ok := dict[s[i:j]]; ok { dp[j] = true … }：如果片段是字典詞，代表從開頭到 j 的前綴可切分，標記 dp[j]=true


› - i=1,2 不可行，跳過？


• 對，因為 dp[1]、dp[2] 沒被設成 true：從開頭出發能切出的字典詞只有 "cat"(到 3) 和 "cats"(到 4)，沒有任何詞能剛好停在索引 1 或 2，所以 dp[1]=dp[2]=false，外層迴圈遇到時就 continue，不再嘗試從那裡往後延伸。


› 但如果字典的第二個字有是 'ca' 跳過會怎樣？


• 不會漏掉。「跳過」只發生在 dp[i]==false 的位置，表示目前還沒有任何切分能到達那裡。如果字典有 "ca"，在 i=0 時內層會檢查長度 2 的片段 s[0:2]="ca"，命中後會把 dp[2]=true 設起來。等外層跑到 i=2 時，就不會被 continue，而是從那裡繼續往後試
  （例如檢查 s[2:3]="t", s[2:4]="ts" 等等）。所以只有「目前無法到達的前綴」才會被跳過；一旦某個位置因為字典詞變成可達，就會在之後的外層迴圈被處理。


  少了 if !dp[i] { continue } 會把「到不了的前綴」也拿來擴展，導致錯誤的 true。例子：s="ab", 字典 ["b"]。正確答案應該是 false，因為沒有詞能涵蓋開頭的 "a"。但如果你不檢查 dp[i]，在 i=1 時雖然 dp[1]=false，內層仍會看 s[1:2]="b"，發現它在字
  典裡就把 dp[2]=true，最後回傳 true，錯了。continue 能確保只從「已知可達」的切分點往後延伸，不會平空跳到一個原本 unreachable 的位置。
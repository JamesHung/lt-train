# Longest Palindromic Substring

## Problem Statement
- Input: A string `s` where `1 <= len(s) <= 1000`, containing digits and English letters.
- Output: The longest palindromic substring within `s`. If multiple answers exist, any one is acceptable.

## Approach
- Use the expand-around-center technique. For each index `i`, treat `i` as the center of an odd-length palindrome and `(i, i+1)` as the center of an even-length palindrome.
- Expand outward while characters match, track the length returned by each expansion, and update the global best substring boundaries.
- The algorithm runs in `O(n^2)` time in the worst case but uses only `O(1)` extra space.

## Intuition (中文)
1. 迴文一定是從某個中心往兩側對稱展開，因此以每個索引當作中心（同時考慮奇數與偶數長度）即可涵蓋所有可能。
2. 每次從中心向左右比較字元，直到不相等為止，就能得到該中心能形成的最長迴文長度。
3. 若找到的長度比目前紀錄的更長，就更新起迄索引，最後回傳該區間的子字串。

## Complexity
- Time: `O(n^2)` because each expansion can traverse the string, and we do it for each center.
- Space: `O(1)` extra space aside from a few indices.


## real use

Real-World Uses of “Longest Palindromic Substring”

(English → 中文在下方)

Where it shows up in practice

Bioinformatics: Detect hairpins/inverted repeats (palindromic motifs) in DNA/RNA; same scanning idea, but match reverse-complements.

Security/Anomaly Detection: Spot symmetric chunks in URLs/tokens as features for phishing/malware heuristics; long palindromes in “random” IDs hint at weak RNG.

Data Quality / ETL: Catch mirror/concatenation bugs that accidentally form prefix + body + reversed(prefix); route suspicious OCR/RTL artifacts to manual review.

NLP & UX: Real-time palindrome highlighting in editors/word games; extend to “near palindromes” for spellcheckers or symmetry-based mining.

Compression/Indexing: Palindromic tree (Eertree) / Manacher’s mindset informs string indexes and feature extraction even if compressors don’t target palindromes directly.

Your stack (CINNOX-style):

Guardrails: unit tests flag unusually long palindromes in IDs/room codes.

Threat intel: use longest-palindrome length as a URL risk feature with entropy, homoglyphs, TLD.

Nightly ETL check: scan new keys for abnormal symmetry to catch formatter regressions.

Why the LeetCode exercise matters

Builds two-pointer expansion / linear scanning instincts for streaming logs, sliding-window analytics, and per-request latency budgets.

Sharpens handling of odd vs even centers, boundaries, and O(n) vs O(n²) trade-offs.

Opens doors to streaming variants and domain tweaks (reverse-complement, approximate matches).

「最長回文子字串」在真實世界的應用

(中文)

會出現在哪些場景

生物資訊： 找 DNA/RNA 的髮夾結構／倒置重複（常見回文樣式）；概念同樣是中心擴張，只是比對的是互補鹼基序列。

資安／異常偵測： URL／Token 若出現長對稱片段可作為網路釣魚或惡意樣本的特徵；標榜隨機的 ID 若有長回文，可能代表 RNG 或編碼出問題。

資料品質／ETL： 拼接流程不小心做出 prefix + body + reversed(prefix) 的鏡像錯誤；OCR 或 RTL 方向處理出包可用回文長度做 QA 提醒。

NLP 與產品體驗： 在編輯器／字謎遊戲中即時高亮回文字串；延伸到「近似回文」（允許少量誤差）用於拼字建議或樣式探勘。

壓縮／索引思維： 雖然壓縮不會特別針對回文，但像 Eertree、Manacher 的思考有助於字串索引與特徵抽取。

你目前的後端場域：

風險控管：單元測試偵測 ID／房間碼是否出現不尋常長度的回文。

威脅情資：把「最長回文長度」與熵值、同形異義字、TLD 風險一起作為 URL 分數的特徵。

夜間 ETL 守門：掃描新鍵值的對稱性，提早抓到格式器回歸。Real-World Uses of “Longest Palindromic Substring”

(English → 中文在下方)

Where it shows up in practice

Bioinformatics: Detect hairpins/inverted repeats (palindromic motifs) in DNA/RNA; same scanning idea, but match reverse-complements.

Security/Anomaly Detection: Spot symmetric chunks in URLs/tokens as features for phishing/malware heuristics; long palindromes in “random” IDs hint at weak RNG.

Data Quality / ETL: Catch mirror/concatenation bugs that accidentally form prefix + body + reversed(prefix); route suspicious OCR/RTL artifacts to manual review.

NLP & UX: Real-time palindrome highlighting in editors/word games; extend to “near palindromes” for spellcheckers or symmetry-based mining.

Compression/Indexing: Palindromic tree (Eertree) / Manacher’s mindset informs string indexes and feature extraction even if compressors don’t target palindromes directly.

Your stack (CINNOX-style):

Guardrails: unit tests flag unusually long palindromes in IDs/room codes.

Threat intel: use longest-palindrome length as a URL risk feature with entropy, homoglyphs, TLD.

Nightly ETL check: scan new keys for abnormal symmetry to catch formatter regressions.

Why the LeetCode exercise matters

Builds two-pointer expansion / linear scanning instincts for streaming logs, sliding-window analytics, and per-request latency budgets.

Sharpens handling of odd vs even centers, boundaries, and O(n) vs O(n²) trade-offs.

Opens doors to streaming variants and domain tweaks (reverse-complement, approximate matches).

「最長回文子字串」在真實世界的應用

(中文)

會出現在哪些場景

生物資訊： 找 DNA/RNA 的髮夾結構／倒置重複（常見回文樣式）；概念同樣是中心擴張，只是比對的是互補鹼基序列。

資安／異常偵測： URL／Token 若出現長對稱片段可作為網路釣魚或惡意樣本的特徵；標榜隨機的 ID 若有長回文，可能代表 RNG 或編碼出問題。

資料品質／ETL： 拼接流程不小心做出 prefix + body + reversed(prefix) 的鏡像錯誤；OCR 或 RTL 方向處理出包可用回文長度做 QA 提醒。

NLP 與產品體驗： 在編輯器／字謎遊戲中即時高亮回文字串；延伸到「近似回文」（允許少量誤差）用於拼字建議或樣式探勘。

壓縮／索引思維： 雖然壓縮不會特別針對回文，但像 Eertree、Manacher 的思考有助於字串索引與特徵抽取。

你目前的後端場域：

風險控管：單元測試偵測 ID／房間碼是否出現不尋常長度的回文。

威脅情資：把「最長回文長度」與熵值、同形異義字、TLD 風險一起作為 URL 分數的特徵。

夜間 ETL 守門：掃描新鍵值的對稱性，提早抓到格式器回歸。
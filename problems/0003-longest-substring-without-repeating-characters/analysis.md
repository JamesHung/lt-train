# Longest Substring Without Repeating Characters

## Approach
- Maintain a sliding window `[start, i]` with no duplicate characters.
- Track the most recent index for each rune in a map. When a character repeats and its last index lies inside the current window, move `start` to one past that index to drop the duplicate.
- Update the maximum window length after each step. Since each character enters and leaves the window at most once, the algorithm runs in linear time.

### Intuition (中文)
1. 用兩個指標 `start`、`i` 當尺的左右端。`i` 每次往右走一個字元，`start` 只在發現重複時往右跳。
2. 用 map 記住每個字元上次出現在哪。如果字元 `ch` 又出現，而且上次位置還在尺子內，就把 `start` 移到 `lastSeen[ch]+1`，這樣尺內就沒有重複。
3. 每回合用 `i - start + 1` 算出當前尺長，如果比歷史最大值大就更新答案。

## Complexity
- Time: O(n) where n is the number of runes in the string.
- Space: O(k) for the map, where k is the size of the character set (at most n).

## 中文說明
- 使用滑動視窗維護 `[start, i]` 範圍內不含重複字元的子字串，並以 map 記錄每個字元最後出現的位置。
- 當遇到重複字元且其位置仍在視窗內時，把 `start` 移到該位置的下一個索引，視窗便移除舊的重複字元。
- 每次更新最大視窗長度即可得到答案，因此整體時間複雜度為 O(n)，額外空間為 O(k)。

## More challenges
159. Longest Substring with At Most Two Distinct Characters
340. Longest Substring with At Most K Distinct Characters
992. Subarrays with K Different Integers

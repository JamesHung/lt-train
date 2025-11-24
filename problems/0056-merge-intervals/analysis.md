## Original Prompt
Given an array of intervals where `intervals[i] = [start_i, end_i]`, merge all overlapping intervals and return the non-overlapping intervals that cover the entire input. Two intervals that touch (e.g., `[1,4]` and `[4,5]`) are considered overlapping.

**Example 1**
- Input: `intervals = [[1,3],[2,6],[8,10],[15,18]]`
- Output: `[[1,6],[8,10],[15,18]]`
- Explanation: `[1,3]` overlaps with `[2,6]`, so they combine into `[1,6]`.

**Example 2**
- Input: `intervals = [[1,4],[4,5]]`
- Output: `[[1,5]]`
- Explanation: Touching at `4` counts as overlap.

**Example 3**
- Input: `intervals = [[4,7],[1,4]]`
- Output: `[[1,7]]`
- Explanation: Sorting then merging yields a single range.

Constraints:
- `1 <= intervals.length <= 10^4`
- `intervals[i].length == 2`
- `0 <= start_i <= end_i <= 10^4`

## English Explanation
Sort the intervals by start (breaking ties by end). Keep a rolling "current" interval initialized to the first sorted entry. Scan the rest from left to right: if the next interval starts at or before `currentEnd`, they overlap, so expand `currentEnd = max(currentEnd, nextEnd)`. Otherwise, push the finished `[currentStart, currentEnd]` to the result and reset the current window to the new interval. Append the final window after the loop. Sorting lets us merge in one pass, touching endpoints merge naturally because we treat `start <= currentEnd` as overlap.

## 中文詳解
先依起點排序（若起點相同就依終點排序），這樣掃描時左到右的區間會依序排列。用第一個區間當作目前的合併範圍 `[curStart, curEnd]`，對每個後續區間：
1. 如果它的起點 `<= curEnd`，代表與當前區間相交或相連，更新 `curEnd = max(curEnd, intervalEnd)` 來延伸範圍。
2. 否則表示不重疊，先把當前範圍加入答案，再把 `[curStart, curEnd]` 重置成新區間。
最後把收尾的區間加入結果。排序花費 `O(n log n)`，合併只需一次線性掃描，額外空間為輸出所需。

## Complexity
- Time: `O(n log n)` due to sorting.
- Space: `O(n)` for the returned merged slice.

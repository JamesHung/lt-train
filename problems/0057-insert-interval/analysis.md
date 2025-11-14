## Original Prompt
You are given an array of non-overlapping `intervals` sorted by start value. Insert `newInterval` into the list while keeping it sorted and merging overlaps. Return the updated intervals.

**Example 1**
- Input: `intervals = [[1,3],[6,9]]`, `newInterval = [2,5]`
- Output: `[[1,5],[6,9]]`

**Example 2**
- Input: `intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]]`, `newInterval = [4,8]`
- Output: `[[1,2],[3,10],[12,16]]`
- Explanation: `[4,8]` overlaps with `[3,5],[6,7],[8,10]`.

Constraints:
- `0 <= intervals.length <= 10^4`
- `0 <= start_i <= end_i <= 10^5`
- intervals sorted by start, no overlaps initially
- `0 <= start <= end <= 10^5`

## English Explanation
We split the scan into three phases. First copy all intervals that end before the new interval starts—those never overlap. Next merge every interval whose start is <= the current `end` of the new interval by expanding `start`/`end` to cover them. After finishing the merging phase, append the combined interval once and then copy the remaining intervals that start after it. Each interval is processed exactly once, so the algorithm runs in linear time and uses only the output list for storage.

## 中文詳解
1. 建立結果陣列，容量預留 `len(intervals)+1`，並把 `start, end` 設成 `newInterval`。
2. 第一段：把所有「結束時間 < start」的區間直接加入結果，因為它們一定在新區間左邊且不重疊。
3. 第二段：只要目前區間的起點 `<= end` 就代表與新區間有重疊或相連，更新 `start = min(start, interval[0])` 與 `end = max(end, interval[1])` 來擴張合併範圍。
4. 合併完後，把這個 `[start, end]` 加入結果。
5. 第三段：把剩下所有區間（必然在右側）依序加入。整個流程單趟完成，僅使用常數額外空間。

## Complexity
- Time: `O(n)` — each interval inspected once.
- Space: `O(n)` — the returned slice stores the merged schedule.

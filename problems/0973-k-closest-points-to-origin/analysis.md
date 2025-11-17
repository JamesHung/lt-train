## Original Prompt
Given `points[i] = [xi, yi]` describing coordinates on the plane and an integer `k`, return the `k` points closest to the origin `(0,0)` by Euclidean distance. The answer can be in any order and is unique up to order.

**Example 1**
- Input: `points = [[1,3],[-2,2]], k = 1`
- Output: `[[-2,2]]`

**Example 2**
- Input: `points = [[3,3],[5,-1],[-2,4]], k = 2`
- Output: `[[3,3],[-2,4]]`

Constraints:
- `1 <= k <= points.length <= 10^4`
- `-10^4 <= xi, yi <= 10^4`

## English Explanation
We only need to order points by their distance to the origin, so sorting is the simplest approach. Instead of computing expensive square roots, we compare the squared distance `x^2 + y^2`, which preserves ordering. Sorting the array of points by this key and slicing off the first `k` items yields one valid answer. The constraints are small enough that `O(n log n)` time easily passes. If we needed to optimize further we could use quickselect or a max-heap, but here simplicity wins.

## 中文詳解
要找離原點最近的 `k` 個點，可以直接依照距離排序。距離的大小比較只要看平方即可，因此計算 `x^2 + y^2` 就能避免開根號。把整個 `points` 依此鍵值排序後，取出前 `k` 個就是答案。整體時間複雜度 `O(n log n)`，在 `n ≤ 10^4` 的範圍內非常足夠。若未來需要更佳效能，也可以改用 quickselect 或維護大小為 `k` 的最大堆，但目前不需要。

## Complexity
- Time: `O(n log n)` due to sorting all points.
- Space: `O(1)` extra aside from in-place sorting (ignoring the space used by the sort implementation).


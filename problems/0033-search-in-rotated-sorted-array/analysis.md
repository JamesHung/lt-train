## Original Prompt
You are given an ascending array `nums` with distinct values that may have been left-rotated at an unknown pivot. Given `nums` and an integer `target`, return the index of `target` if it exists or `-1` otherwise, using an `O(log n)` algorithm.

## English Explanation
Perform binary search, but at every step first decide which half is still sorted. Because a single rotation keeps two sorted pieces, for any window `[lo, hi]` either `[lo, mid]` or `[mid, hi]` is increasing.  
- Compute `mid`. If `nums[mid] == target`, return `mid`.  
- If the left half is sorted (`nums[lo] <= nums[mid]`), the target can only live there when `nums[lo] <= target < nums[mid]`; otherwise it must be in the right half.  
- If the right half is sorted (`nums[mid] <= nums[hi]`), the target can only live there when `nums[mid] < target <= nums[hi]`; otherwise search the left half.  
Updating `lo`/`hi` this way always discards half the window while keeping the target (if any) inside. Stop when `lo` surpasses `hi` and return `-1`.

_Example walk-through_: `nums = [4,5,6,7,0,1,2], target = 0`  
1. `lo=0, hi=6, mid=3 (7)` → left sorted, `0` not in `[4,7)` → search right (`lo=4`).  
2. `lo=4, hi=6, mid=5 (1)` → left sorted, `0` in `[0,1)` → search left (`hi=4`).  
3. `lo=4, hi=4, mid=4 (0)` → found at index 4.

## 中文詳解
核心想法是「先判斷哪一半仍然有序，再決定要丟掉哪一半」。  
1. 設 `lo`、`hi` 為兩端，重複計算 `mid`。若 `nums[mid] == target` 直接回傳。  
2. 若左半段排序 (`nums[lo] <= nums[mid]`)，只有 `nums[lo] <= target < nums[mid]` 時目標可能在左邊，否則必在右邊。  
3. 若右半段排序 (`nums[mid] <= nums[hi]`)，只有 `nums[mid] < target <= nums[hi]` 時目標可能在右邊，否則往左邊搜。  
4. 每次更新 `lo` 或 `hi` 只留下可能含目標的半段，因此視窗會以對半速度縮小。當指標交錯仍未找到，就回傳 `-1`。  

_範例_：`nums = [4,5,6,7,0,1,2], target = 0`  
1. `lo=0, hi=6, mid=3(7)`：左半有序，但 `0` 不在 `[4,7)`，改搜右半 (`lo=4`)。  
2. `lo=4, hi=6, mid=5(1)`：左半有序，`0` 落在 `[0,1)`，縮到左半 (`hi=4`)。  
3. `lo=4, hi=4, mid=4(0)`：找到索引 4。

## Complexity
- Time: `O(log n)` because the binary search discards half the range each iteration.
- Space: `O(1)` extra space.

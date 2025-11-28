## Original Prompt
Given an array of unique integers `nums`, return all possible subsets (the power set). The result must not contain duplicates. `1 <= len(nums) <= 10`, `-10 <= nums[i] <= 10`.

## English Explanation (DFS)
Use backtracking over the index. At position `idx`, choose to skip `nums[idx]` or take it, then recurse to `idx+1`. When `idx` reaches `len(nums)`, copy the current path into the answer. This enumerates every pick/not-pick combination once.

## 中文詳解（DFS）
用遞迴枚舉「選/不選」。在索引 `idx` 時，先不選 `nums[idx]` 往下走，再選 `nums[idx]` 進入下一層；到達 `len(nums)` 時複製當前路徑收集。這樣每個元素兩種狀態，總共 `2^n` 個子集都會被遍歷一次。

## Complexity
- Time: `O(n * 2^n)` — 每個子集的長度拷貝最多 `n`，共有 `2^n` 個。
- Space: `O(n)` 遞迴深度與當前路徑（輸出除外）。

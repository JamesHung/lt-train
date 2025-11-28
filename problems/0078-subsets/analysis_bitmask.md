## Original Prompt
Given an array of unique integers `nums`, return all possible subsets (the power set). The result must not contain duplicates. `1 <= len(nums) <= 10`, `-10 <= nums[i] <= 10`.

## English Explanation
Each element can be either in or out of a subset — a binary choice. With `n` items there are `2^n` subsets. Iterate masks from `0` to `(1<<n)-1`; for each bitmask, include `nums[i]` if bit `i` is set. Collect every subset.

## 中文詳解
每個元素只有「選/不選」兩種狀態，所以共有 `2^n` 個子集合。用位元遮罩表示：對 `mask` 從 `0` 到 `2^n-1`，若 `mask` 的第 `i` 位為 `1`，就把 `nums[i]` 放進子集合。全部列舉即可。

## Complexity
- Time: `O(n * 2^n)` — `2^n` 個遮罩，每個最多掃 `n` 位。
- Space: `O(n)` 臨時子集合（輸出除外）。


• 位元遮罩就是用每一個 bit 代表某元素「選/不選」。例：nums = [1, 2, 3]，長度 3，所以有 2^3 = 8 個 mask，從 0 到 7（二進位 000~111）：

  - 000 (0)：三個 bit 都是 0，子集 []
  - 001 (1)：只第 0 位是 1，取 nums[0]=1 → [1]
  - 010 (2)：只第 1 位是 1，取 nums[1]=2 → [2]
  - 011 (3)：第 0、1 位是 1，取 1,2 → [1,2]
  - 100 (4)：只第 2 位是 1，取 nums[2]=3 → [3]
  - 101 (5)：第 0、2 位是 1，取 1,3 → [1,3]
  - 110 (6)：第 1、2 位是 1，取 2,3 → [2,3]
  - 111 (7)：三位都是 1，取 1,2,3 → [1,2,3]
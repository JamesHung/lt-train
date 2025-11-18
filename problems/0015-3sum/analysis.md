## Original Prompt
Given an integer array `nums`, return all distinct triplets `[nums[i], nums[j], nums[k]]` such that `i`, `j`, and `k` are different indices and `nums[i] + nums[j] + nums[k] == 0`. The order of triplets does not matter and duplicate triplets must be removed.

**Example 1**
- Input: `nums = [-1,0,1,2,-1,-4]`
- Output: `[[-1,-1,2], [-1,0,1]]`

**Example 2**
- Input: `nums = [0,1,1]`
- Output: `[]`

**Example 3**
- Input: `nums = [0,0,0]`
- Output: `[[0,0,0]]`

Constraints:
- `3 <= nums.length <= 3000`
- `-10^5 <= nums[i] <= 10^5`

## English Explanation
The standard pattern is to sort the array, then fix the first element of the triplet and use a two-pointer sweep to find the other two numbers. Sorting allows us to use the classic `target = -nums[i]` approach: for each index `i`, move pointers `l = i+1` and `r = n-1`, shrinking the window depending on whether `nums[l] + nums[r]` is too small or too large. When we find a match we record the triplet and skip over duplicate values to avoid repeated answers. Sorting also makes it easy to skip repeated `nums[i]` values so we only process each distinct first element once.

## 中文詳解
先對陣列排序，這樣能使用「固定一個數、雙指標找另外兩個數」的套路。每次選定 `nums[i]` 後，設目標為 `-nums[i]`，然後用兩個指標 `l=i+1`、`r=n-1` 往中間夾。若 `nums[l]+nums[r]` 小於目標就把左指標往右移，大於目標就把右指標往左移；相等時代表找到一組解，記錄後把兩側相同的值跳過，避免重複答案。同理，外層迴圈也要跳過重複的 `nums[i]`，確保每種起始值只處理一次。

## Complexity
- Time: `O(n^2)` — sorting is `O(n log n)` and the two-pointer sweep for each index runs in total `O(n^2)`.
- Space: `O(1)` extra beyond the output (sorting is in-place).


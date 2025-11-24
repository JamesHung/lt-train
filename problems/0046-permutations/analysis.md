## Original Prompt
Given an array of distinct integers `nums`, return all possible permutations of `nums` in any order.

## English Explanation
Use backtracking to build permutations one position at a time. Maintain a `path` slice and a `used` boolean slice the same length as `nums`. At each depth, iterate all indices; when `used[i]` is false, mark it, append `nums[i]` to `path`, recurse, then undo the choice. When `path` length equals `len(nums)`, copy it into the result.

_Example (`[1,2,3]`)_  
Pick `1` → path `[1]`; pick `2` → `[1,2]`; pick `3` → `[1,2,3]` stored. Backtrack to try `[1,3,2]`, then start with `2` and `3`, producing all 6 permutations.

## 中文詳解
用回溯在每個位置挑一個尚未用過的數字：  
1. 準備 `path` (當前排列) 和 `used` 陣列記錄哪些 index 已選。  
2. 在每一層，遍歷所有 index；若 `used[i]` 為假，就選 `nums[i]`，標記 `used[i]=true`，加入 `path`，遞迴下一層。  
3. 當 `path` 長度等於 `nums` 長度時，代表一個完整排列，複製到答案。  
4. 回溯時把 `used[i]` 解除、`path` 彈出，繼續嘗試其它數字。  

## Complexity
- Time: `O(n * n!)` because there are `n!` permutations and each copies `O(n)` elements.  
- Space: `O(n)` for recursion depth and `used`/`path` (結果輸出除外)。


#
 解釋 path = path[:len(path)-1] 為什麼？
 例子：nums=[1,2,3]

  - 選 1，path=[1]，遞迴。
      - 選 2，path=[1,2]，遞迴。
          - 選 3，path=[1,2,3]，存入結果。回到上一層，此時應還原成 [1,2] 去嘗試「換 3」的分支。
            如果不做 path = path[:len(path)-1]，path 仍是 [1,2,3]，下一輪再嘗試「選 3」或「選其他」都從錯的狀態開始，結果會變成 [1,2,3,3] 之類的錯誤或重複排列。彈出最後一個元素才能乾淨地回到 [1,2]，再換成 [1,3]、或回到 [1] 改選 2、3，正確
            遍歷所有分支。

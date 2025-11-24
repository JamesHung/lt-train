## Original Prompt
Given an array of distinct integers `candidates` and an integer `target`, return all unique combinations of `candidates` where the chosen numbers sum to `target`. The same number may be chosen unlimited times. The number of unique valid combinations is fewer than 150 for the provided tests.

## English Explanation
Use backtracking on a sorted `candidates` array so we can prune early. Maintain a `path` and a remaining sum `rem`. Starting from index `start`, iterate each candidate `c`:
- If `c > rem`, stop because later values are larger (array is sorted).
- Otherwise push `c`, recurse with the same index (since reuse is allowed) and `rem - c`, then pop to try the next candidate.
When `rem` reaches zero, copy the current path into the result. Picking from `start` onward prevents generating permutations of the same combination.

_Example (`candidates=[2,3,6,7], target=7`)_  
Sorted array is the same. Explore `2`: path `[2]`, rem `5`; pick `2` again `[2,2]`, rem `3`; pick `3` gives rem `0`, record `[2,2,3]`. Backtrack and try `7`, rem `0`, record `[7]`. No other paths fit.

## 中文詳解
核心是用回溯列舉所有組合，並靠排序後的「過大就停止」來剪枝：  
1. 先排序 `candidates`，確保後面的數字只會更大。  
2. 定義遞迴 `dfs(start, rem)`，`start` 是可選數字的開頭索引，`rem` 是尚需的總和。  
3. 迴圈從 `start` 走訪每個候選值 `c`：  
   - 若 `c > rem`，後面都更大，直接 `break`。  
   - 否則把 `c` 放入路徑，呼叫 `dfs(i, rem-c)`（可以重複選同個 `c`，所以索引不前進），回來後再移除 `c`。  
4. 當 `rem == 0` 時，複製當前路徑加入答案。  
5. 因為每層只選擇「從 `start` 之後」的數字，不會產生相同數字不同順序的重複解。  

## Complexity
- Time: `O(S)` where `S` is the total size of all valid combinations; each candidate is tried in a depth-first manner with pruning on `c > rem`.
- Space: `O(T)` recursion and path depth, bounded by `target` (worst case用一堆 2 累加)。

## 
cur = append(cur, v)：把這個值放進路徑，遞迴往更深一層。
  2. 遞迴返回時，要把剛剛放進去的值拿掉，讓路徑回到「沒選這個值」的狀態，好嘗試下一個候選。Go 沒有內建 pop，所以用切片縮短一格：cur = cur[:len(cur)-1]。

  如果不做這步，cur 會越疊越長，下一輪會帶著舊的選擇，組合結果就錯亂或重複。


=== RUN   TestCombinationSum/example_1
push 2 -> temp [2] (rem 5)
push 2 -> temp [2 2] (rem 3)
push 2 -> temp [2 2 2] (rem 1)
return -> temp [2 2 2] (next pop 2)
pop 2 -> temp [2 2]
push 3 -> temp [2 2 3] (rem 0)
return -> temp [2 2 3] (next pop 3)
pop 3 -> temp [2 2]
return -> temp [2 2] (next pop 2)
pop 2 -> temp [2]
push 3 -> temp [2 3] (rem 2)
return -> temp [2 3] (next pop 3)
pop 3 -> temp [2]
return -> temp [2] (next pop 2)
pop 2 -> temp []
push 3 -> temp [3] (rem 4)
push 3 -> temp [3 3] (rem 1)
return -> temp [3 3] (next pop 3)
pop 3 -> temp [3]
return -> temp [3] (next pop 3)
pop 3 -> temp []
push 6 -> temp [6] (rem 1)
return -> temp [6] (next pop 6)
pop 6 -> temp []
push 7 -> temp [7] (rem 0)
return -> temp [7] (next pop 7)
pop 7 -> temp []
## Original Prompt
Given an array `nums`, return the element that appears more than ⌊n / 2⌋ times. The majority element always exists. Follow-up asks for an `O(n)` time, `O(1)` space approach.

## English Explanation
Use the Boyer-Moore voting algorithm. Maintain a candidate and a counter; when the counter hits zero, set the current number as candidate. Increment the counter if the current element equals the candidate, otherwise decrement. Since the majority element dominates by more than half, it is guaranteed to be the final candidate.

## 中文詳解
1. 初始化 `candidate = 0`、`count = 0`。
2. 逐一走訪陣列：
   - 若 `count == 0`，把當前數字設為新的 `candidate`。
   - 如果數字等於 `candidate`，`count++`，否則 `count--`。
3. 因為多數元素佔比超過一半，最終 `candidate` 必然是答案。


## Complexity
- Time: `O(n)` single pass.
- Space: `O(1)` extra storage.


### 小記
用hashmap 也可以 但Time Complexity是O(n), Space O(1)
或是可以先sort 直接取中間值, 因為會經過中間點

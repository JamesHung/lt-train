## Original Prompt
Given an integer array `nums`, return `true` if any value appears at least twice in the array, and return `false` if every element is distinct. Constraints: `1 <= nums.length <= 1e5` and each value lies in `[-1e9, 1e9]`.

## English Explanation
The fastest general-purpose check uses a hash set. Iterate through `nums`, and for each value check whether it already exists in the set. If it does, we have found a duplicate and can return `true`. Otherwise insert the value and continue. If the loop finishes, all elements were unique so return `false`.

## 中文詳解
1. 建立一個 `set`（在 Go 中用 map[ int]struct{} 表示）。
2. 逐一走訪 `nums`，如果目前元素已經存在於 set，代表出現第二次，直接回傳 `true`。
3. 若不存在，就把它加入 set，繼續往下處理。
4. 直到迴圈結束都沒有撞到已存在的元素，就說明所有值都獨一無二，回傳 `false`。

## Complexity
- Time: `O(n)` because each element is inserted and looked up once on average.
- Space: `O(n)` in the worst case when every number is distinct and must stay in the set.

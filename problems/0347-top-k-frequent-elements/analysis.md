## Original Prompt
Return the `k` most frequent numbers in `nums`. `1 <= len(nums) <= 1e5`, `-1e4 <= nums[i] <= 1e4`, and the answer is unique. Follow-up asks for better than `O(n log n)` time.

## English Explanation
Count each value with a hash map, then bucket the numbers by frequency. Because a value can appear at most `n` times, we can create `n + 1` buckets where index `f` stores all numbers appearing `f` times. Scan buckets from highest frequency downward and collect numbers until `k` slots are filled. Bucket construction and the final scan are both linear.

## 中文詳解
1. 先用 map 統計每個數字出現次數。
2. 建立 `len(nums) + 1` 個 bucket，bucket `f` 裡放所有出現 `f` 次的數字。
3. 從最大頻率開始往下掃，逐一把 bucket 內的數字加入答案直到湊滿 `k` 個。
4. 由於每個數字只進出 bucket 一次，整體時間僅 `O(n)`。

## Complexity
- 時間：`O(n)`，計數和 bucket 掃描皆線性。
- 空間：`O(n)`，map 與 bucket 存放所有元素。

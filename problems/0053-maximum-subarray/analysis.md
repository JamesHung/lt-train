## Original Prompt
Given an integer array `nums`, find the subarray with the largest sum, and return its sum.

**Example 1**
- Input: `nums = [-2,1,-3,4,-1,2,1,-5,4]`
- Output: `6`
- Explanation: The subarray `[4,-1,2,1]` has the largest sum `6`.

**Example 2**
- Input: `nums = [1]`
- Output: `1`
- Explanation: The subarray `[1]` has the largest sum `1`.

**Example 3**
- Input: `nums = [5,4,-1,7,8]`
- Output: `23`
- Explanation: The subarray `[5,4,-1,7,8]` has the largest sum `23`.

Constraints
- `1 <= nums.length <= 10^5`
- `-10^4 <= nums[i] <= 10^4`

Follow up: If you have figured out the `O(n)` solution, try coding another solution using the divide and conquer approach.

## English Explanation
Kadane's algorithm keeps a rolling sum for the best subarray that ends at the current index. When extending the previous subarray hurts more than starting fresh, we reset the rolling sum to the current value; otherwise we extend it. Track the maximum of all rolling sums as the answer. This single pass covers all candidates in `O(n)` time and `O(1)` extra memory.

## Follow-up 0053-1: Divide & Conquer
Split the array at the midpoint and recursively solve the left and right halves. The maximum subarray for the whole interval is the maximum of:
- the best entirely inside the left half,
- the best entirely inside the right half,
- or a subarray that crosses the midpoint.

To evaluate the crossing candidate in linear time per level, walk left from `mid` and keep the largest suffix sum, then walk right from `mid+1` and keep the largest prefix sum; their sum is the best crossing segment. Recursing until single elements gives the base value. The recurrence performs `O(n)` work per depth with `O(log n)` levels, so the runtime is `O(n log n)` and the recursion stack costs `O(log n)` space.

## 中文詳解
1. 維護 `current` 代表「以當前元素為結尾」的最佳子陣列和。
2. 每次處理新元素 `num` 時，比較 `current + num` 與 `num`，決定是延續舊區間還是重新開始。如果舊區間變成負擔，就直接從 `num` 重新累積。
3. 另外維護 `best`，隨時更新遇過的最大值。這樣所有可能的區間都會在掃描過程中被計算到，避免重複計算。
4. 全程只需單趟掃描與常數額外空間。
5. 注意 i= 1 開始, 先取預設值index 0

## Complexity
- Time: `O(n)` — single pass across the input.
- Space: `O(1)` — only a few integers are stored regardless of input size.

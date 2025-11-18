## Original Prompt
Given `nums`, return an array where each index `i` contains the product of all numbers except `nums[i]`. The answer fits in 32-bit signed integer, division is not allowed, and runtime must be `O(n)`.

**Example 1**
- Input: `nums = [1,2,3,4]`
- Output: `[24,12,8,6]`

**Example 2**
- Input: `nums = [-1,1,0,-3,3]`
- Output: `[0,0,9,0,0]`

Constraints: 
- `2 <= nums.length <= 10^5`, `-30 <= nums[i] <= 30`

## English Explanation
Compute prefix products and suffix products without extra arrays. First pass left-to-right: store in `result[i]` the product of all numbers to the left (maintain a running `prefix`). Second pass right-to-left: keep a `suffix` product and multiply `result[i]` by it before updating `suffix *= nums[i]`. This way each cell gets product of numbers before and after it. This uses O(1) extra space besides the output.

## 中文詳解
不能用除法就用前綴 / 後綴乘積。第一次從左到右，維持 `prefix`（左邊所有數的乘積），把它存進 `result[i]`，再讓 `prefix *= nums[i]`。第二次從右到左，維持 `suffix`（右邊乘積），先讓 `result[i] *= suffix`，再 `suffix *= nums[i]`。最後 `result[i]` 就是左側乘積 × 右側乘積，正好等於「除了 `i` 之外所有元素的乘積」。

## Complexity
- Time: `O(n)` — two passes through the array.
- Space: `O(1)` extra (output array excluded).


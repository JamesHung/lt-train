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

### 例子：`nums = [1, 2, 3, 4]`

1. **Prefix pass（左到右）**  
   - `prefix` 初始為 1，表示「在陣列最左邊還沒有數字」時的乘積。  
   - i=0：result[0] 先寫入 `prefix = 1`，因為左邊沒有數；寫完後 `prefix *= nums[0] = 1`。  
   - i=1：result[1] 寫入目前 `prefix = 1`（左邊只有 1）；再更新 `prefix *= nums[1] = 2`。  
   - i=2：result[2] 寫入 `prefix = 2`（1×2）；更新後 `prefix = 6`。  
   - i=3：result[3] 寫入 `prefix = 6`（1×2×3）；更新後 `prefix = 24`。  
   - 此時 result 變成 `[1, 1, 2, 6]`，正好是每個位置左邊所有數字的乘積（prefix array）。

2. **Suffix pass（右到左）**  
   - `suffix` 也從 1 開始，代表「在陣列最右邊還沒有數」。  
   - i=3：result[3]=6 乘上 `suffix = 1`，保持 6，然後 `suffix *= nums[3] = 4`。  
   - i=2：result[2]=2 乘上 `suffix = 4` 得到 8，再把 `suffix *= nums[2] = 12`。  
   - i=1：result[1]=1 乘上 12 得到 12，更新 `suffix *= nums[1] = 24`。  
   - i=0：result[0]=1 乘上 24 得到 24。  
   - 這一趟等於把右邊所有元素的乘積（postfix array `[24, 12, 4, 1]`）逐一乘進 result。

3. **最終結果**  
   - result 變成 `[24, 12, 8, 6]`，對於每個 i 來說就是「左邊乘積 × 右邊乘積」，也就是除掉自己後剩下所有元素的乘積。

> 直覺想法：prefix pass 幫每個位置先記下「左邊乘積」，suffix pass 再補上「右邊乘積」，兩者相乘就完成題目要求。

## Complexity
- Time: `O(n)` — two passes through the array.
- Space: `O(1)` extra (output array excluded).

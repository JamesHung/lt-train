## Original Prompt
Design a stack supporting `push`, `pop`, `top`, and retrieving the minimum element all in `O(1)` time. Implement `MinStack` with:
- `MinStack()` constructor
- `Push(val)`
- `Pop()`
- `Top()`
- `GetMin()`

Example:
```
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]
```
Output: `[null,null,null,null,-3,null,0,-2]`
Constraints: values are in [`-2^31`, `2^31-1`], calls up to `3 * 10^4`, and `pop`, `top`, `getMin` are only called on non-empty stacks.

## English Explanation
Keep two stacks in one structure: the normal value stack, and a parallel stack that remembers the minimum after each push. Whenever we push, we compare the incoming value with the previous minimum (top of the min stack) and store the smaller of the two. When we pop, both stacks shrink together so the remaining top of the `mins` stack is the current minimum. `Top` simply returns the last pushed value, while `GetMin` looks at the top of the `mins` stack. Every operation only touches the slice tail so each one runs in constant time.

## 中文詳解
把資料堆疊和「目前最小值」堆疊一起維護。Push 時除了把值放進主堆疊，也把 `min(當前值, 之前的最小值)` 推進另一個堆疊，表示「推入這個元素之後堆疊的最小值」。Pop 兩個堆疊都一起彈出，這樣剩下的最小值就會停在 `mins` 的頂端。Top 直接讀主堆疊頂端，GetMin 讀 `mins` 頂端，全部都是 slice 末端操作，複雜度都是常數。

## Complexity
- Time: `O(1)` for every operation.
- Space: `O(n)` to store the stack elements plus the mirrored min stack.

# lt-train

Repository for collecting LeetCode problems, their Go solutions, and accompanying Markdown analyses.

### Workflow
- When a new LeetCode prompt is provided, create a subdirectory or file entry in this repo for that problem.
- Create `handmade.gobk` for the function with empty implmentation 
- Implement the solution in Go (`solution.go` file) following idiomatic Go patterns.
- Add a Markdown file that explains the approach, complexity, and any relevant insights.
- Keep each problem’s artifacts together (e.g., same folder or shared prefix) for easy navigation.

### Additional Rules
- 每題的 `analysis.md` 必須包含原始題目, 英文解釋, 中文詳解，方便快速回顧。
- 預設在解法中加入可控制的 debug log（例如用常數開關 `debugLogs` 或 `debugMedianLogs`），並在除錯時開啟輸出。

### Problems
- [0001 Two Sum](problems/0001-two-sum)
- [0002 Add Two Numbers](problems/0002-add-two-numbers)
- [0003 Longest Substring Without Repeating Characters](problems/0003-longest-substring-without-repeating-characters)
- [0004 Median of Two Sorted Arrays](problems/0004-median-of-two-sorted-arrays)
- [0005 Longest Palindromic Substring](problems/0005-longest-palindromic-substring)
- [0006 Zigzag Conversion](problems/0006-zigzag-conversion)
- [0015 3Sum](problems/0015-3sum)
- [0033 Search in Rotated Sorted Array](problems/0033-search-in-rotated-sorted-array)
- [0039 Combination Sum](problems/0039-combination-sum)
- [0046 Permutations](problems/0046-permutations)
- [0102 Binary Tree Level Order Traversal](problems/0102-binary-tree-level-order-traversal)
- [0133 Clone Graph](problems/0133-clone-graph)
- [0150 Evaluate Reverse Polish Notation](problems/0150-evaluate-reverse-polish-notation)
- [0155 Min Stack](problems/0155-min-stack)
- [0207 Course Schedule](problems/0207-course-schedule)
- [0200 Number of Islands](problems/0200-number-of-islands)
- [0208 Implement Trie (Prefix Tree)](problems/0208-implement-trie-prefix-tree)
- [0322 Coin Change](problems/0322-coin-change)
- [0238 Product of Array Except Self](problems/0238-product-of-array-except-self)
- [0020 Valid Parentheses](problems/0020-valid-parentheses)
- [0021 Merge Two Sorted Lists](problems/0021-merge-two-sorted-lists)
- [0053 Maximum Subarray](problems/0053-maximum-subarray)
- [0056 Merge Intervals](problems/0056-merge-intervals)
- [0057 Insert Interval](problems/0057-insert-interval)
- [0067 Add Binary](problems/0067-add-binary)
- [0070 Climbing Stairs](problems/0070-climbing-stairs)
- [0075 Sort Colors](problems/0075-sort-colors)
- [0098 Validate Binary Search Tree](problems/0098-validate-binary-search-tree)
- [0104 Maximum Depth of Binary Tree](problems/0104-maximum-depth-of-binary-tree)
- [0110 Balanced Binary Tree](problems/0110-balanced-binary-tree)
- [0121 Best Time to Buy and Sell Stock](problems/0121-best-time-to-buy-and-sell-stock)
- [0125 Valid Palindrome](problems/0125-valid-palindrome)
- [0141 Linked List Cycle](problems/0141-linked-list-cycle)
- [0146 LRU Cache](problems/0146-lru-cache)
- [0169 Majority Element](problems/0169-majority-element)
- [0206 Reverse Linked List](problems/0206-reverse-linked-list)
- [0217 Contains Duplicate](problems/0217-contains-duplicate)
- [0226 Invert Binary Tree](problems/0226-invert-binary-tree)
- [0232 Implement Queue using Stacks](problems/0232-implement-queue-using-stacks)
- [0235 Lowest Common Ancestor of a Binary Search Tree](problems/0235-lowest-common-ancestor-of-a-binary-search-tree)
- [0236 Lowest Common Ancestor of a Binary Tree](problems/0236-lowest-common-ancestor-of-a-binary-tree)
- [0242 Valid Anagram](problems/0242-valid-anagram)
- [0278 First Bad Version](problems/0278-first-bad-version)
- [0347 Top K Frequent Elements](problems/0347-top-k-frequent-elements)
- [0383 Ransom Note](problems/0383-ransom-note)
- [0409 Longest Palindrome](problems/0409-longest-palindrome)
- [0542 01 Matrix](problems/0542-01-matrix)
- [0543 Diameter of Binary Tree](problems/0543-diameter-of-binary-tree)
- [0876 Middle of the Linked List](problems/0876-middle-of-the-linked-list)
- [0704 Binary Search](problems/0704-binary-search)
- [0721 Accounts Merge](problems/0721-accounts-merge)
- [0733 Flood Fill](problems/0733-flood-fill)
- [0994 Rotting Oranges](problems/0994-rotting-oranges)
- [1765 Map of Highest Peak](problems/1765-map-of-highest-peak)
- [0973 K Closest Points to Origin](problems/0973-k-closest-points-to-origin)
- [0981 Time Based Key-Value Store](problems/0981-time-based-key-value-store)

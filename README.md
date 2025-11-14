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
- [0020 Valid Parentheses](problems/0020-valid-parentheses)
- [0021 Merge Two Sorted Lists](problems/0021-merge-two-sorted-lists)
- [0067 Add Binary](problems/0067-add-binary)
- [0070 Climbing Stairs](problems/0070-climbing-stairs)
- [0104 Maximum Depth of Binary Tree](problems/0104-maximum-depth-of-binary-tree)
- [0110 Balanced Binary Tree](problems/0110-balanced-binary-tree)
- [0121 Best Time to Buy and Sell Stock](problems/0121-best-time-to-buy-and-sell-stock)
- [0125 Valid Palindrome](problems/0125-valid-palindrome)
- [0141 Linked List Cycle](problems/0141-linked-list-cycle)
- [0146 LRU Cache](problems/0146-lru-cache)
- [0169 Majority Element](problems/0169-majority-element)
- [0206 Reverse Linked List](problems/0206-reverse-linked-list)
- [0226 Invert Binary Tree](problems/0226-invert-binary-tree)
- [0232 Implement Queue using Stacks](problems/0232-implement-queue-using-stacks)
- [0235 Lowest Common Ancestor of a Binary Search Tree](problems/0235-lowest-common-ancestor-of-a-binary-search-tree)
- [0236 Lowest Common Ancestor of a Binary Tree](problems/0236-lowest-common-ancestor-of-a-binary-tree)
- [0242 Valid Anagram](problems/0242-valid-anagram)
- [0278 First Bad Version](problems/0278-first-bad-version)
- [0347 Top K Frequent Elements](problems/0347-top-k-frequent-elements)
- [0383 Ransom Note](problems/0383-ransom-note)
- [0409 Longest Palindrome](problems/0409-longest-palindrome)
- [0543 Diameter of Binary Tree](problems/0543-diameter-of-binary-tree)
- [0876 Middle of the Linked List](problems/0876-middle-of-the-linked-list)
- [0704 Binary Search](problems/0704-binary-search)
- [0733 Flood Fill](problems/0733-flood-fill)

# Add Two Numbers

## Approach
- Walk both input lists simultaneously while tracking a carry.
- At each step sum the carry plus the current node values (treating missing nodes as zero), create a new node with `sum % 10`, and keep `sum / 10` as the next carry.
- Continue iterating until both lists and the carry are exhausted, guaranteeing correct handling of uneven lengths and a trailing carry.
- Use a dummy head to simplify list construction.

## Complexity
- Time: O(max(m, n)) where m and n are the lengths of the lists, because each node is visited once.
- Space: O(1) auxiliary space (besides the output list) since we only store pointers and integers.

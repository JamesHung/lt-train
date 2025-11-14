# Two Sum

## Approach
- Iterate through the array once while maintaining a hash map from value to index.
- For each element `val`, check if `target - val` already exists in the map. If it does, the stored index together with the current index form the unique answer.
- Otherwise store the current value/index pair and continue. Since only one answer exists, the loop returns immediately when the pair is found.

## Complexity
- Time: O(n) because each element is processed at most once.
- Space: O(n) in the worst case for the hash map when no pair is found until the end of the loop.

## Follow-up
- The hash map approach is strictly better than the naive O(n²) double loop and satisfies the sub-quadratic requirement.

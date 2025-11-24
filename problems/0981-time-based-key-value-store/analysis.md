## Original Prompt
Design a time-based key-value store that supports storing multiple values for the same key at different timestamps and retrieving the value at a target time. Implement a class with:
- `TimeMap()` — initialize the structure.
- `set(key, value, timestamp)` — store `value` for `key` at `timestamp`.
- `get(key, timestamp)` — return the value with the largest timestamp `<=` the given one, or empty string if none exists.

Timestamps in `set` are strictly increasing for each key.

**Example**
```
Input:
["TimeMap", "set", "get", "get", "set", "get", "get"]
[[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
Output:
[null, null, "bar", "bar", null, "bar2", "bar2"]
```

Constraints:
- `1 <= key.length, value.length <= 100`
- keys/values are lowercase letters or digits
- `1 <= timestamp <= 10^7`
- at most `2 * 10^5` operations

## English Explanation
Use a map from key to a slice of `(timestamp, value)` pairs. Because timestamps for a given key are added in strictly increasing order, each slice stays sorted. `set` simply appends to the slice. For `get`, binary search the slice to find the first element whose timestamp is greater than the query; the answer is the element just before it (largest timestamp `<= target`). If no such element exists, return an empty string. Each operation is logarithmic in the number of entries per key.

## 中文詳解
用一個 `map[string][]pair` 對應每個 key 的歷史 `(timestamp, value)`。題目保證同一 key 的 `set` 時間戳嚴格遞增，因此不用再排序。`set` 直接把新的 pair 追加到 slice。`get` 對該 key 的 slice 做二分搜尋：找到第一個時間戳「大於」查詢時間的位置，答案就是它前一個元素；如果二分結果為 0，代表沒有時間戳小於等於查詢值，回傳空字串。這樣保持 `O(log n)` 查詢，追加仍是 `O(1)`。

## Complexity
- Time: `O(log m)` per `get` (m = versions for the key), `O(1)` per `set`.
- Space: `O(n)` to store all entries across keys.

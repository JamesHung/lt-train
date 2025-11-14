## Original Prompt
Given a string `s` containing lowercase and uppercase letters, return the length of the longest palindrome that can be built using those letters. Letters are case sensitive.

## English Explanation
The palindrome can use every even-count letter in full, and odd-count letters contribute all but one of their characters to mirrored halves. If any odd-count letter exists, exactly one single character may sit in the center. Count letter frequencies, add `(count / 2) * 2` for each, and keep a flag when an odd count appears to add one more character at the end.

## 中文詳解
1. 先統計字串中每個字元的出現次數（大小寫視為不同字元）。
2. 對於每個字元，只要取 `(count / 2) * 2` 這個偶數部分加入答案，代表可以成對放到左右兩邊。
回文的左右兩半必須對稱，某個字元若要被放到兩邊，就必須拿偶數個（2、4、6…）分成一半放左、一半放右。因此對於次數為 count 的字元，我們只能使用「最大且不超過它的偶數」，也就是 (count / 2) * 2。這樣確保每個字元用量都能分成兩份，左右對稱排進
  回文。剩下來的 1（如果 count 為奇數）只能留給回文中心
  
3. 若某個字元出現奇數次，記錄有「剩下 1 個」可放在回文中心。
中心最多只能放一個字元。雖然有兩個字元是奇數次 (1 和 3)，但 (count/2)*2 已經把它們各自的偶數部分放進左右兩邊，剩下的兩個「落單」字元都想佔中心，只能選其中一個；另一個必須捨棄。程式邏輯就是先累積所有偶數部分，再用 hasOdd 記錄「是否存在任
  何奇數次字元」，最後若 hasOdd 為真就只額外加 1，不會因為多個奇數而重複加。
4. 迭代完成後，只要曾經有奇數次出現，就把答案再加 1，表示使用一個中心字元。

## Complexity
- Time: `O(n)` to traverse the string once and the frequency map.
- Space: `O(1)` relative to charset size (bounded by 52 letters).

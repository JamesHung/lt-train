## Original Prompt
Given a list `accounts` where `accounts[i] = [name, email1, email2, ...]`, merge accounts that share at least one common email. All accounts of the same person share the same name. Return merged accounts where the first element is the name and the remaining elements are the sorted unique emails; the order of accounts in the output does not matter. Constraints: up to `1000` accounts, each with `2-10` strings, email length up to `30`.

## English Explanation
Treat each email as a node and union any emails that appear in the same account, because they must belong to the same person. A disjoint-set union (DSU) maps every email to a parent; for each account we union the first email with all others. We also remember `email -> name` (name is consistent for all emails of a person). After processing all accounts, we iterate over every email, find its root, and bucket emails by root. Finally, sort each bucket’s emails and prepend the owner’s name from the map.

## 中文詳解
把每個 email 當作圖上的節點，同一個帳號內的 email 必定屬於同一個人，因此把它們 union 起來。用並查集（DSU）維護 `parent` 與 `rank`，先確保帳號中的第一個 email 存在，再將它與後續每個 email union。額外用 `email -> name` 記錄擁有者姓名（題目保證同一人的所有帳號名字一致）。所有帳號處理完後，逐一檢查每個 email，找到它的代表元並將其收集到 `root -> emails` 的桶子中。最後對每個桶排序 email，並用桶內任何一個 email 取得對應的姓名，組合成 `[name, sortedEmails...]`。

## Complexity
- Time: `O(N log N)` where `N` is total emails; unions are near-`O(1)` and we sort emails per component.
- Space: `O(N)` for DSU maps and grouped emails.

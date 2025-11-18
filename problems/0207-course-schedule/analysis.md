## Original Prompt
Given `numCourses` labeled `0..numCourses-1` and prerequisite pairs `[a, b]` meaning you must take `b` before `a`, determine if you can finish all courses.

**Example 1**
- Input: `numCourses = 2, prerequisites = [[1,0]]`
- Output: `true`

**Example 2**
- Input: `numCourses = 2, prerequisites = [[1,0],[0,1]]`
- Output: `false`

Constraints:
- `1 <= numCourses <= 2000`
- `0 <= prerequisites.length <= 5000`

## English Explanation
Treat courses as nodes in a directed graph. The question asks if the graph has a cycle. Use Kahn’s algorithm (BFS topological sort): build adjacency lists and indegree counts, enqueue every course with indegree zero, and repeatedly pop a course, “take” it, and decrement indegrees of its dependents. If at the end we processed all courses, the graph is acyclic and every course is reachable; otherwise a cycle exists and finishing is impossible.

## 中文詳解
把課程當作有向圖的節點，`[a,b]` 代表 `b -> a`。問題等同於圖中是否存在環。使用 Kahn 拓樸排序：先計算每個節點的入度，將所有入度為 0 的課程加入佇列。每次從佇列取出一門課，視為已修完，並把它指向的課程入度減一，如果某門課入度降到 0 就也入佇列。若最終能處理完 `numCourses` 門課，表示沒有環；否則存在互相依賴的循環，無法全部修完。

###

• - graph := make([][]int, numCourses) 建一個長度為 numCourses 的鄰接表，每門課對應一個 slice，儲存它後續可以解鎖的課程。
- indegree := make([]int, numCourses) 建立入度陣列，indegree[i] 表示課程 i 被多少先修課指向。
- for _, edge := range prerequisites { ... } 逐筆處理先修關係 edge = [a,b]：graph[b] = append(graph[b], a) 把 b → a 連邊；indegree[a]++ 代表 a 多了一個必修先決。
- queue := make([]int, 0, numCourses) 建 BFS 佇列，for course, deg := range indegree 把所有入度為 0 的課程 enqueue，這些課不需要先修就能修。
- 若 debugCanFinish 為 true，印出初始 indegree 和 queue 供偵錯。
- visited := 0 計數已「修完」的課，for head := 0; head < len(queue); head++ 以 head 指標遍歷 queue，模擬不斷取出佇列前端。
- cur := queue[head] 代表現在修的課，visited++ 增加完成數；若 debug，印出資訊。
- for _, next := range graph[cur] 遍歷 cur 解鎖的後續課程，對每個 next 做 indegree[next]--，並視情況印 debug。
- 當某課的入度變為 0，表示所有先修都完成，queue = append(queue, next) 把它加入佇列，必要時印出追蹤。
- 最後若有開啟 debug，再印 visited 與 numCourses；函式回傳 visited == numCourses，代表是否能修完全部課程（即圖是否無環）。

## Complexity
- Time: `O(numCourses + prerequisites.length)` because each edge and node is processed once.
- Space: `O(numCourses + prerequisites.length)` for the adjacency list, indegrees, and queue.


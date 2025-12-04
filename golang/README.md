# Go practice plan

Use this repo to keep small Go experiments under `examples/` and to track the study path below. Current sample: `go run ./examples/simplejson`.

## 1. Go language fundamentals

- Syntax & types: ints/uints, rune/byte, string, bool, float32/64, complex.
- Composite types: arrays vs slices (backing array, len vs cap, append, re-slicing, copy); maps (zero value nil, creation, lookup, delete, random iteration order); structs (embedding vs inheritance, tags for JSON/DB).
- Pointers: `&` and `*`, zero value nil, when to pick pointer vs value receivers.
- Functions: multiple return values (including `(T, error)`), named returns and why naked returns are avoided, variadic functions and spreading slices.
- Control flow & error handling: if, for, switch (including type switch); idiomatic error handling with `%w`; panic vs error vs `log.Fatal`.
- Interfaces: implicit implementation, small behavior-focused designs, compile-time assertions, interface value holding a nil concrete value.
- Nil interface check: an interface is `nil` only when both type and value are nil. Factories should return `nil` (type=nil, data=nil) on invalid input; avoid returning a nil pointer wrapped in an interface. See `examples/nil_value_check/main.go` for factory vs bad factory, wrapped-nil detection, and a repository pattern that returns `(*T, error)` instead of an interface.
- Methods & receivers: value vs pointer receiver, mutation vs copy cost, method sets and receiver choice.

## 2. Memory model: slices, maps, escape analysis

- Slices: internals (ptr/len/cap), append in-place vs new allocation, sharing slices across goroutines, `copy(dst, src)` semantics.
- Slice behavior quick tests:
  - Test 1 — basic slicing:

    ```go
    s := []int{1, 2, 3, 4}
    t := s[1:3]
    // s=[1 2 3 4], t=[2 3]
    ```

  - Test 2 — len/cap after slicing:

    ```go
    t := s[1:3]
    fmt.Println(len(t), cap(t)) // 2 3 (cap counts from start index)
    ```

  - Test 3 — shared backing array:

    ```go
    s := []int{1, 2, 3}
    t := s[:2]
    t[0] = 999
    // s=[999 2 3], t=[999 2]
    ```

  - Test 4 — append reuses backing array (enough cap):

    ```go
    s := []int{1, 2, 3}
    t := s[:2]
    t = append(t, 100)
    // s=[1 2 100], t=[1 2 100]
    ```

  - Test 5 — append reallocates (cap insufficient):

    ```go
    s := []int{1, 2, 3}
    t := s[:2]
    t = append(t, 100, 200)
    // s=[1 2 3], t=[1 2 100 200]
    ```

  - Test 6 — full slice expression to cap-limit:

    ```go
    s := []int{1, 2, 3, 4}
    t := s[1:2:2]
    fmt.Println(len(t), cap(t)) // 1 1 (forces next append to allocate)
    ```

  - Test 7 — zero-length slice (shares cap):

    ```go
    s := []int{1, 2, 3}
    t := s[:0]
    fmt.Println(t, len(t), cap(t)) // [] 0 3
    ```

  - Test 8 — append after zero-length slice uses backing array:

    ```go
    s := []int{1, 2, 3}
    t := s[:0]
    t = append(t, 999)
    // s=[999 2 3], t=[999]
    ```

  - Test 9 — copy creates new memory:

    ```go
    s := []int{1, 2, 3}
    t := make([]int, len(s))
    copy(t, s)
    t[0] = 999
    // s=[1 2 3], t=[999 2 3]
    ```

  - Test 10 — append to nil slice works:

    ```go
    var s []int
    s = append(s, 123)
    // s=[123], len=1, cap=1
    ```

  - Core concepts: slice is a view (ptr/len/cap), append reuses backing array until cap is exceeded then allocates, slicing itself does not copy unless you `copy` or restrict cap with `[low:high:max]`.
- Append and shared backing array:

  ```go
  s := []int{1, 2, 3} // len=3 cap=3
  t := s[:2]          // len=2 cap=3 (same backing array)
  t = append(t, 999)
  // cap(t) had room, so append wrote in place:
  // s => [1 2 999], t => [1 2 999]

  t = s[:2]
  t = append(t, 999, 1000) // exceeds cap -> allocates new array
  // s => [1 2 3], t => [1 2 999 1000]
  ```

  Why: slices carry len/cap pointers into the same backing array; append reuses it until cap is exceeded, then allocates a new array.
- Sharing slices across goroutines: writing from multiple goroutines causes data races. Example:

```go
s := []int{1, 2, 3}

go func() { s[0] = 100 }()        // write
go func() { fmt.Println(s) }()    // read; data race
```

Solutions: guard with `sync.Mutex`, copy before sharing, serialize through channels, or treat the slice as immutable.

See `examples/slice_safety/main.go` for runnable samples of each approach.
- copy(dst, src) semantics:

```go
dst := make([]int, 3)
src := []int{100, 200, 300, 400}

n := copy(dst, src)
fmt.Println(dst, n) // [100 200 300], n = 3 (min(len(dst), len(src)))
```

Rules: copies `min(len(dst), len(src))`, does not grow slices. Useful for cloning: `clone := append([]T(nil), original...)` (or `append([]T{}, original...)`).
- Maps:
  - Zero value nil:

    ```go
    var m map[string]int
    fmt.Println(len(m)) // 0 works
    m["a"] = 1          // panic: assignment to entry in nil map
    ```

  - Initialize with make:

    ```go
    m = make(map[string]int)
    m["a"] = 1
    ```

  - Hint capacity (optional): `m := make(map[string]int, 100)`; runtime may allocate larger buckets as needed.
  - Not safe for concurrent writes without locks or `sync.Map`:

    ```go
    m := make(map[int]int)

    go func() { m[1] = 10 }()    // write
    go func() { fmt.Println(m) }() // read
    // can panic: concurrent map writes (undefined behavior on older Go)
    ```

    Safer options: guard with `sync.Mutex`, prefer `sync.RWMutex` for read-heavy access, or use `sync.Map` for many goroutines and rare writes.
    See `examples/map_safety/main.go` for runnable samples.
- Escape analysis: when values move to heap, patterns that cause escape (returning locals, interface conversions), rough understanding of why.
  - Q3: does this escape?

    ```go
    func foo() []int {
        s := []int{1, 2, 3}
        return s
    }
    ```

    Yes. The slice header is returned and its backing array must outlive the function, so the array is allocated on the heap.

  - Q4: does this escape?

    ```go
    func bar() string {
        s := "hello"
        return s
    }
    ```

    No escape. Strings are immutable and small; the compiler can keep it on the stack or in read-only data.

### Summary cheat sheet (interview quick recall)

- Slices: ptr/len/cap; append in-place until cap exceeded then new array; slicing shares backing array; avoid sharing writable slices across goroutines; `copy(dst, src)` copies `min(len)` of each.
- Maps: zero value nil (reads len=0, writes panic), use `make(map[K]V, capHint)`, not safe for concurrent writes; use mutex or `sync.Map`.
- Escape analysis: returning pointers/slices of locals escapes; interface conversions can force escape; closures capturing locals escape; more heap allocations mean more GC.

## 3. Concurrency

- Goroutines: `go` keyword, when to use, lifetime/cleanup and leak risks.
- Channels:
  - Unbuffered `make(chan T)` → send blocks until receive (sync handoff).
  - Buffered `make(chan T, n)` → send until buffer full; useful for decoupling/queueing.
  - `close(ch)`: receive returns zero value with `ok=false`; send to closed panics.
- select patterns:
  - Multiplex send/receive across channels.
  - `default` branch to avoid blocking.
  - Timeouts with `time.After`.
- sync package: `sync.WaitGroup`, `sync.Mutex` / `sync.RWMutex`, `sync.Once`.
- Context: `context.Background`/`TODO`, `WithCancel`/`WithTimeout`/`WithDeadline`, passing context first and honoring cancellation; loops should select on `ctx.Done()` to avoid leaks.
- Common follow-up answers:
  1) Prefer ticker + select over `time.Sleep` in loops so cancellation via `ctx.Done()` can break immediately.
  2) Context should be passed in from the caller so lifecycle is controlled externally; otherwise goroutines can leak.
  3) `ctx.Done()` is closed on cancel, so the loop will exit immediately and not hang.
  - Context patterns:

    ```go
    // WithTimeout (common)
    ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
    defer cancel()
    res, err := client.Do(ctx)
    if err != nil {
        return fmt.Errorf("api error: %w", err)
    }

    // WithDeadline (fixed cutoff)
    deadline := time.Now().Add(200 * time.Millisecond)
    ctx, cancel = context.WithDeadline(context.Background(), deadline)
    defer cancel()

    // WithCancel (manage goroutine lifetime)
    ctx, cancel = context.WithCancel(context.Background())
    go func() {
        defer wg.Done()
        for {
            select {
            case <-ctx.Done():
                return
            default:
                work()
            }
        }
    }()
    // later:
    cancel()
    ```
- Q5-1 Unbuffered vs buffered channel:
  - Unbuffered `make(chan T)`: send/receive must rendezvous; both block until the other side is ready. Use for sync hand-offs or signals.
  - Buffered `make(chan T, n)`: send does not block until buffer is full; receive drains buffer first. Use for producer/consumer queues, decoupling speeds.
  - One-liner: unbuffered = synchronous hand-off; buffered = asynchronous queue.
- Q5-2 RWMutex:
  - Readers share the read lock concurrently; writers take exclusive lock.
  - Best when reads dominate and writes are infrequent (caches/config/shared maps).
  - One-liner: `Mutex` → all access serialized; `RWMutex` → concurrent reads, writes still exclusive.
  - Upgrade? You cannot upgrade `RLock` to `Lock`; doing so deadlocks. Drop the read lock first or redesign (plain `Mutex`, copy-on-write, etc.).
- Benchmarks (see `examples/rwmutex_cache` and run `go test -bench=. -run ^$ ./examples/rwmutex_cache`):
  - Read-heavy: `sync.RWMutex` usually outperforms `sync.Mutex` because reads can run concurrently.
  - Write-heavy: `sync.RWMutex` may be no faster or even slower; extra complexity and write locks remove the read advantage.
  - Interview soundbite: “In read-heavy cases RWMutex helps because readers run concurrently. In write-heavy cases Mutex can be simpler and faster. Choose based on read/write ratio, not ‘RWMutex is always faster’.”
- See `examples/concurrency/main.go` for goroutine + WaitGroup, unbuffered/buffered channels, channel close, select multiplex/default/timeout, mutex, context cancel/timeout with cancellation-safe loops, and a producer/consumer with timeout. See `examples/rwmutex_cache` for an RWMutex-protected cache demo and mutex vs RWMutex benchmarks. See `examples/cache_worker` for a background-refresh cache that avoids starvation by keeping read locks short and doing loads outside the lock.

## 4. Standard library to know

- fmt, errors, log; time (Duration, Sleep, Ticker, Timer).
- net/http: `http.HandleFunc`, `ListenAndServe`, handlers and middleware with `http.Handler`.
- encoding/json: tags (`json:"user_id,omitempty"`, `-`), custom marshal/unmarshal methods.
- io (Reader, Writer, Copy); os (files, env vars); for backend work, `database/sql` as needed.

## 5. Tooling & ecosystem

- go mod: `go mod init`, `go get`, replace directives, versioning.
- go build, go test, go vet.
- Benchmarks with `testing.B`.
- Profiling basics: pprof at a high level.
- Project structure: `cmd/`, `pkg/`, `internal/`; avoid circular deps.

## 6. Interview-style coding exercises

- Classic coding: stack/queue via slices; worker pool with goroutines + channels; rate limiter (token/leaky bucket with `time.Ticker`).
- String/slice tasks: reverse string/slice (UTF-8 caveat), remove duplicates.
- Map/struct tasks: grouping data, counting frequencies.
- Concurrency exercises: producer–consumer (single vs multiple producers/consumers), fan-out/fan-in, graceful shutdown via context or channel close.
- Design exercises: job processing service design; request timeout/cancellation end-to-end with contexts across handlers and downstream calls.

## 7. Go-specific conceptual questions

- Generics history (pre-1.18) and basic syntax now.
- new vs make; nil slice vs empty slice; panic vs `os.Exit`.
- Common closure bug when taking the address of range loop variable.
- Data race: what/why, detecting with `-race`; rough scheduler understanding (goroutines to OS threads).

## 8. System design with Go

- RESTful API service design: routing, handlers, middleware, error handling, logging, metrics.
- Observability: structured logging, metrics, tracing.
- Config/environment: env vars, config files, flags or config library.
- Testing strategy: unit tests (testing package, mocking interfaces), integration tests (test DB or docker-compose), table-driven tests.

## 9. Study schedule (2–3 weeks)

- Week 1 – language & concurrency: Day 1–2 types/slices/maps/interfaces + small JSON examples; Day 3–4 concurrency/context (N workers + cancel); Day 5–7 stdlib/tooling, tiny net/http REST API with in-memory store.
- Week 2 – coding & patterns: daily 1–2 problems; implement worker pool, fan-in/fan-out, rate limiter, graceful shutdown server; add tests/benchmarks for at least one.
- Week 3 – mock interviews & system design: do 2–3 mocks; prep 2–3 story projects (service built in Go, perf issue, concurrency bug) with context/problem/options/solution/results.

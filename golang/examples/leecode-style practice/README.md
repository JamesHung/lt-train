# 6. Common Interview-Style Coding Questions in Go

Interview questions may not be LeetCode-hard, but they will test your Go fluency.

## Classic Coding Exercises
- Stack/queue using slices
- Worker pool using goroutines + channels
- Rate limiter (token bucket or leaky bucket) using `time.Ticker` + channels

## String and Slice Manipulation
- Reverse string/slice (be careful with UTF-8 when reversing strings)
- Remove duplicates from a slice

## Map and Struct Tasks
- Grouping data
- Counting frequencies

## Concurrency Exercises

### Producer–Consumer Patterns
- Single producer, multiple consumers
- Multiple producers, single consumer

### Fan-Out / Fan-In
- Spread work across multiple goroutines (fan-out)
- Aggregate results back (fan-in)

### Graceful Shutdown
- Stop goroutines on context cancellation
- Stop goroutines on channel close

## Design Exercises in Go (Senior Level)

### Design a Job Processing Service
- API layer
- Queue
- Workers
- Retries
- Idempotency

### Request Timeout and Cancellation (End-to-End)
- HTTP handler receives context
- Pass the same context to DB calls
- Pass the same context to external API calls

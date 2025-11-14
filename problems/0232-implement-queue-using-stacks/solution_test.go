package implementqueueusingstacks

import "testing"

func TestMyQueue(t *testing.T) {
	q := Constructor()

	q.Push(1)
	q.Push(2)

	if peek := q.Peek(); peek != 1 {
		t.Fatalf("Peek() = %d, want 1", peek)
	}

	if pop := q.Pop(); pop != 1 {
		t.Fatalf("Pop() = %d, want 1", pop)
	}

	if q.Empty() {
		t.Fatalf("Empty() = true, want false")
	}

	if pop := q.Pop(); pop != 2 {
		t.Fatalf("Pop() = %d, want 2", pop)
	}

	if !q.Empty() {
		t.Fatalf("Empty() = false, want true")
	}
}

func TestInterleavedOperations(t *testing.T) {
	q := Constructor()
	for i := 1; i <= 5; i++ {
		q.Push(i)
	}

	for i := 1; i <= 3; i++ {
		if pop := q.Pop(); pop != i {
			t.Fatalf("Pop() = %d, want %d", pop, i)
		}
	}

	q.Push(6)
	q.Push(7)

	expected := []int{4, 5, 6, 7}
	for idx, want := range expected {
		if pop := q.Pop(); pop != want {
			t.Fatalf("Pop #%d = %d, want %d", idx, pop, want)
		}
	}

	if !q.Empty() {
		t.Fatalf("queue should be empty")
	}
}

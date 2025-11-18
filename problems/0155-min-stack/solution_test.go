package minstack

import "testing"

func TestMinStackExample(t *testing.T) {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)

	if got := stack.GetMin(); got != -3 {
		t.Fatalf("GetMin() = %d, want -3", got)
	}

	stack.Pop()
	if got := stack.Top(); got != 0 {
		t.Fatalf("Top() = %d, want 0", got)
	}

	if got := stack.GetMin(); got != -2 {
		t.Fatalf("GetMin() = %d, want -2", got)
	}
}

func TestMinStackMixedValues(t *testing.T) {
	stack := Constructor()
	inputs := []int{2, 0, 3, 0, -1, 4}
	mins := []int{2, 0, 0, 0, -1, -1}

	for i, v := range inputs {
		stack.Push(v)
		if got := stack.GetMin(); got != mins[i] {
			t.Fatalf("step %d: GetMin() = %d, want %d", i, got, mins[i])
		}
	}

	expectedPops := []struct {
		top int
		min int
	}{
		{4, -1},
		{-1, -1},
		{0, 0},
		{3, 0},
		{0, 0},
		{2, 2},
	}

	for i, expect := range expectedPops {
		if got := stack.Top(); got != expect.top {
			t.Fatalf("pop step %d: Top() = %d, want %d", i, got, expect.top)
		}
		if got := stack.GetMin(); got != expect.min {
			t.Fatalf("pop step %d: GetMin() = %d, want %d", i, got, expect.min)
		}
		stack.Pop()
	}
}

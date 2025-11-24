package timebasedkeyvaluestore

import "testing"

func TestTimeMapExample(t *testing.T) {
	tm := Constructor()
	tm.Set("foo", "bar", 1)

	if got := tm.Get("foo", 1); got != "bar" {
		t.Fatalf("Get(foo,1)=%q want bar", got)
	}
	if got := tm.Get("foo", 3); got != "bar" {
		t.Fatalf("Get(foo,3)=%q want bar", got)
	}

	tm.Set("foo", "bar2", 4)
	if got := tm.Get("foo", 4); got != "bar2" {
		t.Fatalf("Get(foo,4)=%q want bar2", got)
	}
	if got := tm.Get("foo", 5); got != "bar2" {
		t.Fatalf("Get(foo,5)=%q want bar2", got)
	}
}

func TestTimeMapEdgeCases(t *testing.T) {
	tm := Constructor()

	if got := tm.Get("missing", 10); got != "" {
		t.Fatalf("expected empty for missing key, got %q", got)
	}

	tm.Set("a", "first", 5)
	if got := tm.Get("a", 4); got != "" {
		t.Fatalf("expected empty before first timestamp, got %q", got)
	}
	if got := tm.Get("a", 5); got != "first" {
		t.Fatalf("expected first at exact timestamp, got %q", got)
	}

	tm.Set("a", "second", 10)
	tm.Set("a", "third", 15)
	if got := tm.Get("a", 12); got != "second" {
		t.Fatalf("expected second at ts=12, got %q", got)
	}
	if got := tm.Get("a", 20); got != "third" {
		t.Fatalf("expected third at ts=20, got %q", got)
	}
}

func TestTimeMapIndependentKeys(t *testing.T) {
	tm := Constructor()
	tm.Set("x", "one", 1)
	tm.Set("y", "two", 2)
	tm.Set("x", "three", 3)

	if got := tm.Get("x", 2); got != "one" {
		t.Fatalf("expected one for x@2, got %q", got)
	}
	if got := tm.Get("y", 3); got != "two" {
		t.Fatalf("expected two for y@3, got %q", got)
	}
}

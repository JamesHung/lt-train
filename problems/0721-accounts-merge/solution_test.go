package accountsmerge

import (
	"reflect"
	"sort"
	"testing"
)

func normalizeAccounts(in [][]string) [][]string {
	cpy := make([][]string, len(in))
	for i, acc := range in {
		accCopy := make([]string, len(acc))
		copy(accCopy, acc)
		if len(accCopy) > 1 {
			sort.Strings(accCopy[1:])
		}
		cpy[i] = accCopy
	}
	sort.Slice(cpy, func(i, j int) bool {
		a, b := cpy[i], cpy[j]
		if a[0] != b[0] {
			return a[0] < b[0]
		}
		limit := len(a)
		if len(b) < limit {
			limit = len(b)
		}
		for k := 1; k < limit; k++ {
			if a[k] != b[k] {
				return a[k] < b[k]
			}
		}
		return len(a) < len(b)
	})
	return cpy
}

func TestAccountsMerge(t *testing.T) {
	tests := []struct {
		name     string
		accounts [][]string
		want     [][]string
	}{
		{
			name: "example1",
			accounts: [][]string{
				{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
				{"John", "johnsmith@mail.com", "john00@mail.com"},
				{"Mary", "mary@mail.com"},
				{"John", "johnnybravo@mail.com"},
			},
			want: [][]string{
				{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
				{"Mary", "mary@mail.com"},
				{"John", "johnnybravo@mail.com"},
			},
		},
		{
			name: "example2",
			accounts: [][]string{
				{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"},
				{"Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"},
				{"Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"},
				{"Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"},
				{"Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"},
			},
			want: [][]string{
				{"Ethan", "Ethan0@m.co", "Ethan4@m.co", "Ethan5@m.co"},
				{"Gabe", "Gabe0@m.co", "Gabe1@m.co", "Gabe3@m.co"},
				{"Hanzo", "Hanzo0@m.co", "Hanzo1@m.co", "Hanzo3@m.co"},
				{"Kevin", "Kevin0@m.co", "Kevin3@m.co", "Kevin5@m.co"},
				{"Fern", "Fern0@m.co", "Fern1@m.co", "Fern5@m.co"},
			},
		},
		{
			name: "sameNameNoOverlap",
			accounts: [][]string{
				{"Alex", "a@mail.com"},
				{"Alex", "b@mail.com"},
			},
			want: [][]string{
				{"Alex", "a@mail.com"},
				{"Alex", "b@mail.com"},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := normalizeAccounts(accountsMerge(tc.accounts))
			want := normalizeAccounts(tc.want)
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("accountsMerge() = %v, want %v", got, want)
			}
		})
	}
}

package accountsmerge

import "sort"

type dsu struct {
	parent map[string]string
	rank   map[string]int
}

func newDSU() *dsu {
	return &dsu{
		parent: make(map[string]string),
		rank:   make(map[string]int),
	}
}

func (d *dsu) find(x string) string {
	p, ok := d.parent[x]
	if !ok {
		d.parent[x] = x
		d.rank[x] = 0
		return x
	}
	if p != x {
		d.parent[x] = d.find(p)
	}
	return d.parent[x]
}

func (d *dsu) union(a, b string) {
	ra := d.find(a)
	rb := d.find(b)
	if ra == rb {
		return
	}
	if d.rank[ra] < d.rank[rb] {
		ra, rb = rb, ra
	}
	d.parent[rb] = ra
	if d.rank[ra] == d.rank[rb] {
		d.rank[ra]++
	}
}

// accountsMerge merges accounts by common email using DSU.
func accountsMerge(accounts [][]string) [][]string {
	uf := newDSU()
	emailToName := make(map[string]string)

	for _, acc := range accounts {
		if len(acc) < 2 {
			continue
		}
		name := acc[0]
		firstEmail := acc[1]
		emailToName[firstEmail] = name
		for i := 2; i < len(acc); i++ {
			email := acc[i]
			emailToName[email] = name
			uf.union(firstEmail, email)
		}
	}

	groups := make(map[string][]string)
	for email := range emailToName {
		root := uf.find(email)
		groups[root] = append(groups[root], email)
	}

	result := make([][]string, 0, len(groups))
	for root, emails := range groups {
		sort.Strings(emails)
		name := emailToName[root]
		result = append(result, append([]string{name}, emails...))
	}

	return result
}

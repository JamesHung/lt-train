package timebasedkeyvaluestore

import (
	"sort"
)

type ValueData struct {
	value string
	ts    int
}

// TimeMap stores versions of values per key; left unimplemented for practice.
type TimeMap struct {
	data map[string][]ValueData
}

func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string][]ValueData),
	}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	tm.data[key] = append(tm.data[key], ValueData{value: value, ts: timestamp})
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	values, ok := tm.data[key]
	if !ok {
		return ""
	}

	idx := sort.Search(len(values), func(i int) bool {
		return values[i].ts > timestamp
	})

	if idx == 0 {
		return ""
	}

	return ValueData[idx-1].val
}

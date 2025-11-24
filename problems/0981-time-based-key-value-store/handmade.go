package timebasedkeyvaluestore

import "sort"

type ValueObject struct {
	value string
	ts    int
}

type TimeMap struct {
	data map[string][]ValueObject
}

func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string][]ValueObject),
	}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	tm.data[key] = append(tm.data[key],
		ValueObject{value: value,
			ts: timestamp})
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	value, ok := tm.data[key]
	if !ok {
		return ""
	}

	index := sort.Search(len(value), func(i int) bool {
		return value[i].ts > timestamp
	})

	if index == 0 {
		return ""
	}

	return value[index-1].value
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

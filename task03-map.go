package homework

import (
	"sort"
)


func SortMapValues(input map[int]string) (result []string) {
	var values []string
	var keys []int
	
	for k := range input {
		keys = append(keys, k)
		sort.Ints(keys[:])
	}
	
	for _, v := range keys {
		values = append(values, input[v])
	}

	return values
}

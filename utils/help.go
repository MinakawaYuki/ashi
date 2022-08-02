package utils

func InArray(item string, items []string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func MergeMap(maps ...map[string]any) map[string]any {
	var mapData = make(map[string]any)
	for _, m := range maps {
		for k, v := range m {
			mapData[k] = v
		}
	}
	return mapData
}

func ArrayMerge(arrays ...[]any) []any {
	n := 0
	for _, v := range arrays {
		n += len(v)
	}
	s := make([]any, 0, n)
	for _, v := range arrays {
		s = append(s, v...)
	}
	return s
}

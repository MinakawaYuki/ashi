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

func AnyMapToStringMap(anyMap map[string]interface{}) map[string]string {
	convertedMap := make(map[string]string)

	for key, value := range anyMap {
		if strValue, ok := value.(string); ok {
			convertedMap[key] = strValue
		}
	}

	return convertedMap
}

//func AnyToStringMap(anyData any) map[string]string {
//	convertedMap := make(map[string]string)
//
//	for i := 0; i < len(anyData)-1; i += 2 {
//		if key, ok := anyData[i].(string); ok {
//			if value, ok := anyData[i+1].(string); ok {
//				convertedMap[key] = value
//			}
//		}
//	}
//	return convertedMap
//}

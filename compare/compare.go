package compare

import (
	"fmt"
	"reflect"
)

// compareMaps 比较两个 map[string]interface{} 是否相等，
func compareMaps(m1, m2 map[string]interface{}) bool {
	// 忽略字段个数不同的情况
	// if len(m1) != len(m2) {
	// 	return false
	// }
	for k, v1 := range m1 {
		v2, ok := m2[k]
		if !ok {
			return false
		}
		if !compareValues(v1, v2) {
			return false
		}
	}
	return true
}

func compareValues(v1, v2 interface{}) bool {
	// 如果 v1 和 v2 类型不一样则直接返回 false
	if reflect.TypeOf(v1) != reflect.TypeOf(v2) {
		return false
	}
	// 接下来只判断 map or slice 即可
	switch v1 := v1.(type) {
	case map[string]interface{}:
		v2, ok := v2.(map[string]interface{})
		fmt.Println("======>", v2, ok)
		if !ok {
			return false
		}
		return compareMaps(v1, v2)
	case []interface{}:
		v2, ok := v2.([]interface{})
		if !ok {
			return false
		}
		if len(v1) != len(v2) {
			return false
		}
		for i, v1i := range v1 {
			if !compareValues(v1i, v2[i]) {
				return false
			}
		}
		return true
	default:
		return v1 == v2
	}
}

// mergeMaps 合并两个 map[string]interface{}
//
// 如果 m1 和 m2 中的 key 对应的值都是 map[string]interface{}，那么会递归合并这两个 map[string]interface{}
// 如果 m1 和 m2 中的 key 对应的值都是 []interface{}，那么会将 m2 中的 []interface{} 追加到 m1 中的 []interface{} 后面,
// 如果 m1 和 m2 中的 key 对应的值都是 其他类型，那么会将 m2 中的值覆盖 m1 中的值,
// 返回合并后的 map
func mergeMaps(m1, m2 map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	// Copy m1 to the result map
	for k, v := range m1 {
		result[k] = v
	}

	// Merge m2 into the result map
	for k, v := range m2 {
		// If the key already exists in the result map, check if the value is also a map or slice
		if val, ok := result[k]; ok {
			if m2, ok := v.(map[string]interface{}); ok {
				// If both values are maps, merge them recursively
				if m1, ok := val.(map[string]interface{}); ok {
					result[k] = mergeMaps(m1, m2)
				} else {
					// Otherwise, overwrite the value with the new map
					result[k] = m2
				}
			} else if s2, ok := v.([]map[string]interface{}); ok {
				// If the new value is a slice, append its elements to the existing slice (if any)
				if s1, ok := val.([]map[string]interface{}); ok {

					// iterate the element from s2, if it has the same id with a element in s1, then merge them, otherwise append it to s1
					for i := 0; i < len(s2); i++ {
						// iterate the element from s2
						s2i := s2[i]
						// find the element in s1, find the id
						pos := -1
						for j := 0; j < len(s1); j++ {
							if s2i["id"] != nil && s1[j]["id"] != nil && s1[j]["id"] == s2i["id"] {
								// s1i = s1[j]
								pos = j
								break
							}
						}
						if pos != -1 {
							// 如果存在相同的 id，则直接 merge s1i and s2i
							s1[pos] = mergeMaps(s1[pos], s2i)
						} else {
							// append s2i to s1
							s1 = append(s1, s2i)
						}
					}
					result[k] = s1
				} else {
					// if the old value is not a map or slice, overwrite the value with the new slice
					result[k] = s2
				}
			} else {
				// If the new value is not a map or slice, overwrite the old value
				result[k] = v
			}
		} else {
			// If the key does not exist in the result map, add it
			result[k] = v
		}
	}

	return result
}

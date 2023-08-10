package helpers

import "strings"

func PrependSliceWithLimitAvoidDuplicates(array []string, newElem string, limit int) []string {
	for i, item := range array {
		if strings.ToLower(item) == strings.ToLower(newElem) {
			array = RemoveFromSliceByIndex(array, i)
			break
		}
	}
	return PrependSliceWithLimit(array, newElem, limit)
}

func PrependSliceWithLimit(array []string, newElem string, limit int) []string {
	if len(array) < limit {
		return append([]string{strings.ToLower(newElem)}, array...)
	} else {
		return append([]string{strings.ToLower(newElem)}, array[:limit-1]...)
	}
}

func RemoveFromSliceByIndex(array []string, index int) []string {
	return append(array[:index], array[index+1:]...)
}

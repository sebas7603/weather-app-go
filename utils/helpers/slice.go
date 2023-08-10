package helpers

import "strings"

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

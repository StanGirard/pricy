package helpers

import "strings"

func ConvertStringToInterface(array [][]string) [][]interface{} {
	old := array
	new := make([][]interface{}, len(old))
	for i, v := range old {
		new[i] = make([]interface{}, len(v))
		for j, v2 := range v {
			new[i][j] = strings.Replace(v2, ".", ",", -1)
		}
	}
	return new
}

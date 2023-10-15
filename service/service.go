package service

import "strings"

// ConvertStringToStrList ...  split by ','
func ConvertStringToStrList(str string) *[]string {
	res := strings.Split(str, ",")
	return &res
}

func ConvertStrListToString(list *[]string) string {
	res := ""
	for key, item := range *list {
		if key == 0 {
			res += item
		}
		res += "," + item
	}

	return res
}

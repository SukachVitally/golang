package main

import (
	"fmt"
)

func main() {
	strList := []string{"first", "second"}
	intList := []int{1, 10, 8, 7, 5, 9, 7}
	boolList := []bool{true, false, true}

	list := make([]interface{}, len(strList)+len(intList)+len(boolList))
	for i, s := range strList {
		list[i] = s
	}

	for i, s := range intList {
		list[i+len(strList)] = s
	}

	for i, s := range boolList {
		list[i+len(strList)+len(intList)] = s
	}

	for _, num := range sort(list) {
		fmt.Println(num)
	}
}

func sort(list []interface{}) []interface{} {
	if len(list) < 2 {
		return list
	}

	centerElementValue := getValue(list[int(len(list)/2)])

	var leftList []interface{}
	var centerList []interface{}
	var rightList []interface{}

	for _, num := range list {
		value := getValue(num)
		if value == centerElementValue {
			centerList = append(centerList, num)
		} else if value > centerElementValue {
			rightList = append(rightList, num)
		} else {
			leftList = append(leftList, num)
		}
	}
	leftList = sort(leftList)
	rightList = sort(rightList)

	return append(leftList, append(centerList, rightList...)...)
}

func getValue(element interface{}) string {
	switch element.(type) {
	case int:
		return fmt.Sprintf("%v", element)
	case float64:
		return fmt.Sprintf("%v", element)
	case string:
		return fmt.Sprintf("%v", element)
	case bool:
		return fmt.Sprintf("%v", element)
	default:
		panic("Unsupported type")
	}
}

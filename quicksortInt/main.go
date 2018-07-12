package main

import (
	"fmt"
)

func main() {
	list := []int{1, 10, 8, 7, 5, 9, 7}
	for _, num := range sort(list) {
		fmt.Println(num)
	}
}

func sort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	centerIndex := int(len(list) / 2)
	centerElement := list[centerIndex]

	leftList := []int{}
	centerList := []int{}
	rightList := []int{}

	for _, num := range list {
		if num == centerElement {
			centerList = append(centerList, num)
		} else if num > centerElement {
			rightList = append(rightList, num)
		} else {
			leftList = append(leftList, num)
		}
	}
	leftList = sort(leftList)
	rightList = sort(rightList)

	return append(leftList, append(centerList, rightList...)...)
}

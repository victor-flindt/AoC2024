package main

import (
	"fmt"
)

func main() {

	list1 := make([]int, 3, 3)
	list2 := make([]int, 3, 3)
	list3 := make([]int, 6, 6)

	for i := 0; i < len(list1); i++ {

		list1[i] = i
		fmt.Println(i)
	}

	for i, element := range list1 {
		list2[i] = element + 1
	}

	list3 = append(list1, list1...)

	fmt.Println("list1", list1)

	fmt.Println("list2", list2)

	fmt.Println("list3", list3)

	list3[3] = 1000
	list3[2] = 1001
	list3[3] = 1002
	list3[3] = 1003

	fmt.Println("list3", list3)

	fmt.Println("list1", list1)

}

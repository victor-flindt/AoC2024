package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func is_safe(list1 []float64) (bool, string) {

	// checking logic from exercise description

	// decreasing case
	if list1[0] > list1[1] {
		for index := 0; index < len(list1)-1; index++ {
			if (list1[index] > list1[index+1]) && (list1[index]-3 <= list1[index+1]) {
				continue
			} else {
				return false, fmt.Sprintf("[Decreasing] %v > %v  & %v <= %v ", list1[index], list1[index+1], list1[index]-3, list1[index+1])
			}
		}
		return true, ""
	}

	// equal case
	if list1[0] == list1[1] {
		return false, fmt.Sprintf("[Equal] %v == %v", list1[0], list1[1])
	} else {
		// increasing case
		for index := 0; index < len(list1)-1; index++ {
			if (list1[index] < list1[index+1]) && (list1[index+1] <= list1[index]+3) {
				continue
			} else {
				return false, fmt.Sprintf("[Increasing] %v < %v & %v < %v", list1[index], list1[index+1], list1[index+1], list1[index]+3)
			}
		}
		return true, ""
	}
}

func main() {

	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	total_save := 0
	for {
		// read row in from csv file
		row, err := reader.Read()
		// preallocate array size
		list1 := make([]float64, 0, len(row))

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}
		// converting string to float 64
		for _, element := range row {
			float_element, err := strconv.ParseFloat(element, 64)

			if err != nil {
				panic(err.Error())
			}

			list1 = append(list1, float_element)

		}
		safe, msg := is_safe(list1)

		if safe {
			fmt.Println("row is safe ")
			total_save += 1
		} else {
			fmt.Println("row is unsafe msg:", msg)
		}

	}

	fmt.Println(total_save)

}

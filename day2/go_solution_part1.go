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
	// Open the first file
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// solving problem1
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	var total_safe int = 0
	var total_safe_part2 int = 0
	for {
		// read row in from csv file
		row, err := reader.Read()
		// preallocate array size
		list1 := make([]float64, 0, len(row))
		list2 := make([]float64, 0, len(row))

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

		// Check if it's safe for the original row
		safe, msg := is_safe(list1)
		if safe {
			fmt.Println("row is safe ")
			total_safe += 1
			total_safe_part2 += 1
		} else {
			for index := 0; index < len(list1); index++ {
				// Create a new slice by removing the element at 'index'
				// fmt.Println("INDEX", list1[:index], list1[index+1:])

				list2 = append(list2, list1[:index]...)
				list2 = append(list2, list1[index+1:]...)

				safe, msg = is_safe(list2)
				if safe {
					total_safe_part2 += 1
					break
				} else {
					list2 = list2[:0]
					continue
				}
			}
			fmt.Println("row is unsafe msg:", msg)
		}
	}
	// Print totals after processing both files
	fmt.Println("Total safe part1: ", total_safe)
	fmt.Println("Total safe part2: ", total_safe_part2)
}

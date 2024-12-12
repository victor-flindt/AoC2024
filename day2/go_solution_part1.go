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

		// Check if it's safe for the original row
		safe, msg := is_safe(list1)
		if safe {
			fmt.Println("row is safe ")
			total_safe += 1
		} else {
			fmt.Println("row is unsafe msg:", msg)
		}
	}

	// Open the second file
	file2, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file2.Close()

	// solving problem2
	reader2 := csv.NewReader(file2)
	reader2.FieldsPerRecord = -1
	var total_safe2 int = 0
	for {
		// read row in from csv file
		row, err := reader2.Read()
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

		// Check if the original row is safe
		safe, msg := is_safe(list1)
		if safe {
			fmt.Println("row is safe ")
			total_safe2 += 1
		} else {
			// Iterate over the list and "pop" an element at each index to check safety
			for index := 0; index < len(list1); index++ {
				// Create a new slice by removing the element at 'index'
				var row_popped_element = append(list1[:index], list1[index+1:]...)

				// Print the updated slice
				fmt.Println("INDE")
				fmt.Println(row_popped_element)

				// Check if it's safe after removing the element
				safe2, _ := is_safe(row_popped_element)
				if safe2 {
					total_safe2 += 1
					break // Exit the loop if the condition is satisfied
				}
			}
			// Print the unsafe message if no "safe" configuration was found
			fmt.Println("row is unsafe msg:", msg)
		}
	}
	// Print totals after processing both files
	fmt.Println("Total safe part1: ", total_safe)
	fmt.Println("Total safe part2: ", total_safe2)
}

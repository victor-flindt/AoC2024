package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

func main() {

	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2
	data, err := reader.ReadAll()

	var list1 []float64
	var list2 []float64

	for _, row := range data {

		if len(row) != 0 {
			var el1, _ = strconv.ParseFloat(row[0], 32)
			var el2, _ = strconv.ParseFloat(row[1], 32)
			list1 = append(list1, el1)
			list2 = append(list2, el2)
		}
	}

	slices.Sort(list1)
	slices.Sort(list2)

	var total_distance float64
	for index, _ := range list1 {
		total_distance += math.Abs(list1[index] - list2[index])
	}

	fmt.Println(total_distance)

	defer file.Close()

}

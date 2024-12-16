package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// func pad_matrix(matrix []string) ([]string){
// 	return
// }

func main() {

	// kernal := [][]string{{"S", "*", "*", "S", "*", "*", "S"},
	// 	{"*", "A", "*", "A", "*", "A", "*"},
	// 	{"*", "*", "M", "M", "M", "*", "*"},
	// 	{"S", "A", "M", "X", "M", "A", "S"},
	// 	{"*", "*", "M", "M", "M", "*", "*"},
	// 	{"*", "A", "*", "A", "*", "A", "*"},
	// 	{"S", "*", "*", "S", "*", "*", "S"}}

	// Open the first file
	file, err := os.Open("test_data.data")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// solving problem1
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	var matrix [][]string
	var row_index int = 0

	matrix = append(matrix, []string{"*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*"})
	matrix = append(matrix, []string{"*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*"})
	matrix = append(matrix, []string{"*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*"})
	matrix = append(matrix, []string{"*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*"})

	for {
		row_pad := []string{"*", "*", "*", "*"}
		// read row in from csv file
		row, err := reader.Read()
		if err != nil {
			break
		}

		row_pad = append(row_pad, strings.Split(row[0], "")...)
		row_pad = append(row_pad, []string{"*", "*", "*", "*"}...)

		matrix = append(matrix, row_pad)

		row_pad = []string{}
		row_index += 1
	}
	matrix = append(matrix, []string{"*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*"})
	matrix = append(matrix, []string{"*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*"})
	matrix = append(matrix, []string{"*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*"})
	matrix = append(matrix, []string{"*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*"})

	for i := 4; i < len(matrix[0])-4; i++ {

		for j := 4; j < len(matrix)-4; j++ {
			fmt.Print(matrix[i][j])

		}
		fmt.Println("")
	}

}

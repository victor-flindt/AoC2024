package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func pad_matrix(matrix [][]string, padding_sequence []string, amount_of_rows int) [][]string {

	for i := 0; i < amount_of_rows; i++ {
		matrix = append(matrix, padding_sequence)
	}

	return matrix
}

func main() {

	// Open the first file
	file, err := os.Open("data.data")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// solving problem1
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	var matrix [][]string
	var row_index int = 0
	padding_sequence := []string{"*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*", "*"}

	matrix = pad_matrix(matrix, padding_sequence, 4)

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
	matrix = pad_matrix(matrix, padding_sequence, 4)

	var total_part1 int = 0
	var total_part2 int = 0

	x_mas_slice := []string{"X", "M", "A", "S"}
	x_max_invert_slice := []string{"S", "A", "M", "X"}

	mas_slice := []string{"M", "A", "S"}
	mas_invert_slice := []string{"S", "A", "M"}

	// row
	for i := 4; i < len(matrix[0])-4; i++ {
		// col
		for j := 4; j < len(matrix)-4; j++ {
			if (matrix[i][j] == "X") || (matrix[i][j] == "S") || (matrix[i][j] == "M") {

			} else {
				continue
			}

			// PART 1 exercise

			// horizontal case
			var horizontal_slice []string = matrix[i][j : j+4]
			if reflect.DeepEqual(horizontal_slice, x_mas_slice) || reflect.DeepEqual(horizontal_slice, x_max_invert_slice) {
				// fmt.Println("Horizontal", horizontal_slice, "row:", i-4, "col", j-4)
				total_part1 += 1
			}
			// vertical case
			var vertical_slice []string = []string{matrix[i][j], matrix[i+1][j], matrix[i+2][j], matrix[i+3][j]}
			if reflect.DeepEqual(vertical_slice, x_mas_slice) || reflect.DeepEqual(vertical_slice, x_max_invert_slice) {
				// fmt.Println("vertical", vertical_slice, "row:", i-4, "col", j-4)
				total_part1 += 1
			}
			// tilted up case
			var titled_slice_up []string = []string{matrix[i][j], matrix[i+1][j+1], matrix[i+2][j+2], matrix[i+3][j+3]}
			if reflect.DeepEqual(titled_slice_up, x_mas_slice) || reflect.DeepEqual(titled_slice_up, x_max_invert_slice) {
				// fmt.Println("tilted up", titled_slice_up, "row:", i-4, "col", j-4)
				total_part1 += 1
			}
			// tilted down case
			var titled_slice_down []string = []string{matrix[i][j], matrix[i-1][j+1], matrix[i-2][j+2], matrix[i-3][j+3]}
			if reflect.DeepEqual(titled_slice_down, x_mas_slice) || reflect.DeepEqual(titled_slice_down, x_max_invert_slice) {
				// fmt.Println("tilted down", titled_slice_down, "row:", i-4, "col", j-4)
				total_part1 += 1
			}

			// PART 2 exercise
			var title_slice_up_part2 []string = []string{matrix[i+2][j], matrix[i+1][j+1], matrix[i][j+2]}
			var title_slice_down_part2 []string = []string{matrix[i][j], matrix[i+1][j+1], matrix[i+2][j+2]}

			if (reflect.DeepEqual(title_slice_up_part2, mas_slice) || reflect.DeepEqual(title_slice_up_part2, mas_invert_slice)) &&
				(reflect.DeepEqual(title_slice_down_part2, mas_slice) || reflect.DeepEqual(title_slice_down_part2, mas_invert_slice)) {
				// fmt.Println("tilted up", title_slice_up_part2, "row:", i-4, "col", j-4)
				total_part2 += 1
			}

		}

	}
	fmt.Println("total part1", total_part1)
	fmt.Println("total part2", total_part2)

}

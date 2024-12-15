package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func find_digits(str string) (a int64, b int64) {

	regex, _ := regexp.Compile("\\d{1,3}")

	matches := regex.FindAllString(str, 2)

	a, _ = strconv.ParseInt(matches[0], 10, 32)

	b, _ = strconv.ParseInt(matches[1], 10, 32)

	return a, b

}

func find_dont(data string) [][]int {

	regex, _ := regexp.Compile("do\\(\\)")

	matches := regex.FindAllStringIndex(data, -1)

	return matches

}

func main() {

	bytes, err := os.ReadFile("test_data_part2.data")
	if err != nil {
		panic(err)
	}

	data_as_string := string(bytes)

	// donts := find_dont(data_as_string)

	fmt.Println(data_as_string[59:60], data_as_string[62:63])

	// var total int64 = 0

	// for {
	// 	// read row in from csv file
	// 	row, err := reader.Read()

	// 	if err != nil {
	// 		if err.Error() == "EOF" {
	// 			break
	// 		}
	// 		panic(err)
	// 	}

	// 	// regex, _ := regexp.Compile("mul\\(\\d{1,3}\\,\\d{1,3}\\)")

	// 	// matches := regex.FindAllString(strings.Join(row[:], ","), -1)
	// 	// for _, match := range matches {
	// 	// 	a, b := find_digits(match)
	// 	// 	total += a * b

	// 	// }
	// }

	// fmt.Println(total)

}

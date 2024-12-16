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

func find_mul(data string) []string {
	regex, _ := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)")

	matches := regex.FindAllString(data, -1)

	return matches

}

func find_dont(data string) [][]int {

	regex, _ := regexp.Compile("don't\\(\\)")

	matches := regex.FindAllStringIndex(data, -1)

	return matches

}

func find_do(data string) [][]int {

	regex, _ := regexp.Compile("do\\(\\)")

	matches := regex.FindAllStringIndex(data, -1)

	return matches

}

func main() {

	bytes, err := os.ReadFile("data.data")
	if err != nil {
		panic(err)
	}

	data_as_string := string(bytes)

	donts := find_dont(data_as_string)
	dos := find_do(data_as_string)

	var valid_indexes [][]int
	valid_indexes = append(valid_indexes, []int{0, donts[0][0]})

	dont := 0
	var do int = 0
	for _, do_array := range dos {
		do = do_array[0]
		if do < dont {
			continue
		}
		for _, dont_array := range donts {
			dont = dont_array[0]
			if dont < do {
				continue
			}
			if do < dont {
				start_stop := []int{do, dont}
				valid_indexes = append(valid_indexes, start_stop)
				break
			}
		}
		if dont < do {
			start_stop := []int{do, len(data_as_string)}
			valid_indexes = append(valid_indexes, start_stop)
		}
	}
	var total int64 = 0
	for _, valid_index := range valid_indexes {

		results := find_mul(data_as_string[valid_index[0]:valid_index[1]])

		for _, result := range results {
			a, b := find_digits(result)
			total += a * b
		}
	}
	fmt.Println(total)
}

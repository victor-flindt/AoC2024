package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func extrac_nums(str string) (a int64, b int64) {

	regex, _ := regexp.Compile("\\d{1,3}")

	matches := regex.FindAllString(str, 2)

	a, _ = strconv.ParseInt(matches[0], 10, 32)

	b, _ = strconv.ParseInt(matches[1], 10, 32)

	return a, b

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
	var total int64 = 0

	for {
		// read row in from csv file
		row, err := reader.Read()

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			panic(err)
		}

		regex, _ := regexp.Compile("mul\\(\\d{1,3}\\,\\d{1,3}\\)")

		matches := regex.FindAllString(strings.Join(row[:], ","), -1)
		for _, match := range matches {
			a, b := extrac_nums(match)
			total += a * b

		}
	}

	fmt.Println(total)

}

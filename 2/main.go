package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	data := utils.ParseFile("input", utils.StrToInt)
	i := 1
	results := utils.ArrayMap(data, func(row []int) bool {

		diff := row[0] - row[1]
		prevVal := row[1]

		if diff < 0 && diff > -4 {
			for _, val := range row[2:] {
				diff := prevVal - val

				if diff > -1 || diff < -3 {
					i++
					return false
				}
				prevVal = val
			}

		} else if diff > 0 && diff < 4 {
			for _, val := range row[2:] {
				diff := prevVal - val

				if diff > 3 || diff < 1 {
					i++
					return false
				}
				prevVal = val
			}
		} else {
			i++
			return false
		}

		i++
		return true
	})

	countSafe := utils.ArrayCountOccurances(results)[true]

	fmt.Println(countSafe)
}

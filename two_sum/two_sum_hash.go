import (
	"fmt"
	"strconv"
)

func twoSum(nums []int, target int) []int {

	var hash = make(map[string]int)
	var result []int

	for index, currentNumber := range nums {
		strNumber := strconv.Itoa(currentNumber)
		_, isFilled := hash[strNumber]

		if !isFilled {
			newKey := strconv.Itoa(target - currentNumber)
			fmt.Println(index)
			hash[newKey] = index
		} else {
			result = append(result, hash[strNumber], index)
			break
		}
	}

	return result
}

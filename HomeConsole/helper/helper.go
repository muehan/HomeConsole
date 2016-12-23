package helper

import (
	"fmt"
	"strconv"
)

func StringToInt(val string) (number int) {

	id, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println("Error during parsing")
	}

	return id
}

package payload

import (
	error_vars "dudos/error-vars"
	"fmt"
	"strconv"
	"strings"
)

type Payload struct {
	Type   string
	Object string
}

func New(typePayload string) (string, error) {

	object := ""
	lenPayload, err := getLength(typePayload)
	if err != nil || lenPayload <= 0 {
		return "", fmt.Errorf("%v:%v", error_vars.ErrorOfBadTypePayload, err)
	}

	if countNonDigit(typePayload) != 2 {
		return "", error_vars.ErrorOfBadTypePayload
	}

	switch {
	case strings.Contains(typePayload, "mb"):
		object = "mb"
	case strings.Contains(typePayload, "kb"):
		object = "kb"
	default:
		return "", error_vars.ErrorOfBadTypePayload
	}

	for i := 0; i < lenPayload; i++ {
		object += "x"
	}

	return object, nil
}

func countNonDigit(line string) int {
	cnt := 0
	for i := range line {
		if line[i] < '0' || line[i] > '9' {
			cnt++
		}
	}

	return cnt
}

func getLength(line string) (int, error) {
	for i := range line {
		if line[i] < '0' || line[i] > '9' {
			return strconv.Atoi(line[:i])
		}
	}
	return 0, error_vars.ErrorOfBadTypePayload
}

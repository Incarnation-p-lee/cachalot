package sampling

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	invalidSamplingIntValue = -1
	invalidSamplingStringValue = ""
)

func getFirstStringValue(line string) string {
	if len(line) == 0 {
		return invalidSamplingStringValue
	}

	separator := regexp.MustCompile(`[: ]`)
	values := separator.Split(line, -1)

	if len(values) == 0 || len(values[0]) == 0 {
		return invalidSamplingStringValue
	}

	return values[0]
}

func getFirstIntValue(line string) int {
	if len(line) == 0 {
		return invalidSamplingIntValue
	}

	separator := regexp.MustCompile(`[: ]`)
	values, intValue := separator.Split(line, -1), invalidSamplingIntValue

	values = values[1:]

	for _, v := range values {
		v := strings.Trim(v, " \t")

		if len(v) > 0 {
			if value, err := strconv.Atoi(v); err != nil {
				log.Printf("Failed to convert integer from %s due to %+v\n", v, err)
			} else {
				intValue = value
				break
			}
		}
	}

	return intValue
}

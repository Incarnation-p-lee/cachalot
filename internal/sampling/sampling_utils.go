package sampling

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	invalidSamplingIntValue = -1
)

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

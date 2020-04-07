package durationutil

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// ParseTime parses the string representation of the duration and returns a time.Duration type
// Supports min, h, day, week, month and year in the form of integer + unit like 3day, 5month, etc
func ParseTime(s string) (time.Duration, error) {
	baseRegex := "^(\\d+)"
	regInputToHour := map[string]int{
		baseRegex + "min":   1,
		baseRegex + "h":     60 * 1,
		baseRegex + "day":   60 * 24,
		baseRegex + "week":  60 * 24 * 7,
		baseRegex + "month": 60 * 24 * 30,
		baseRegex + "year":  60 * 24 * 365,
	}
	hours, err := func() (float64, error) {
		for k, v := range regInputToHour {
			re := regexp.MustCompile(k)
			match := re.FindStringSubmatch(s)
			if len(match) == 2 {
				numStr := match[1]
				num, err := strconv.ParseFloat(numStr, 64)
				if err == nil {
					return num * float64(v), nil
				}
			}
		}
		return 0, fmt.Errorf("cannot parse %s into a valid duration", s)
	}()
	if err != nil {
		return 0, err
	}

	return time.Duration(time.Duration(hours) * time.Minute), nil
}

package durationutil

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

// ParseTime parses the string representation of the duration and returns a time.Duration type
func ParseTime(s string) (time.Duration, error) {
	regInputToHour := map[string]int{
		"^(\\d+)d$": 24,
		"^(\\d+)w$": 24 * 7,
		"^(\\d+)m$": 24 * 30,
		"^(\\d+)y$": 24 * 365,
	}
	hours, err := func() (int, error) {
		for k, v := range regInputToHour {
			re := regexp.MustCompile(k)
			match := re.FindStringSubmatch(s)
			if len(match) == 2 {
				numStr := match[1]
				num, err := strconv.Atoi(numStr)
				if err == nil {
					return num * v, nil
				}
			}
		}
		return 0, fmt.Errorf("cannot parse %s into a valid duration", s)
	}()
	if err != nil {
		return 0, err
	}
	return time.Duration(time.Duration(hours) * time.Hour), nil
}

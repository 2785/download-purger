package fsutil

import (
	"fmt"
	"os/user"
	"path/filepath"
	"strings"
)

// ParsePath parses an absolute path from the input
func ParsePath(s string) (string, error) {
	trimmed := strings.TrimSpace(s)

	if strings.HasPrefix(trimmed, "~/") {
		currUser, err := user.Current()
		if err != nil {
			return "", fmt.Errorf("cannot expand ~ in path %s, error getting user", trimmed)
		}
		return strings.Replace(trimmed, "~", currUser.HomeDir, 1), nil
	} else if !filepath.IsAbs(trimmed) {
		p, err := filepath.Abs(trimmed)
		if err != nil {
			return "", fmt.Errorf("cannot get absolute path, error getting working directory: %s", err.Error())
		}
		return p, nil
	}
	return trimmed, nil
}

package common

import (
	"fmt"
	"strings"
)

func parseEndpoint(ep string) (proto, addr string, _ error) {
	if strings.HasPrefix(strings.ToLower(ep), "tcp://") {
		s := strings.SplitN(ep, "://", 2)
		if s[1] != "" {
			return s[0], s[1], nil
		}
	}
	return "", "", fmt.Errorf("invalid endpoint: %v", ep)
}

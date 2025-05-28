package core

import (
	"fmt"
)

// ModifyMinRunBandwidth 修改混跑带宽
func ModifyMinRunBandwidth(hostname string, bandwidth float64, action int64, networkCard, appID string) ([]byte, error) {
	if appID == "" {
		return nil, fmt.Errorf("the appid cannot be empty")
	}

	if hostname == "" {
		return nil, fmt.Errorf("the hostname cannot be empty")
	}

	if action == 0 {
		return nil, fmt.Errorf("the action cannot be zero")
	}

	return nil, nil
}

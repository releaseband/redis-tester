package results

import (
	"encoding/json"
	"fmt"

	"github.com/ayupov-ayaz/redis-tester/hooks/params"
)

type ConnectionsCounter struct {
	Set    uint64 `json:"Set,omitempty"`
	Get    uint64 `json:"Get,omitempty"`
	Del    uint64 `json:"Del,omitempty"`
	RPush  uint64 `json:"RPush,omitempty"`
	LTrim  uint64 `json:"LTRim,omitempty"`
	LRange uint64 `json:"LRange,omitempty"`
}

var (
	ConnectionsCount *ConnectionsCounter
)

func (c *ConnectionsCounter) Add(name string) {
	switch name {
	case params.Set:
		ConnectionsCount.Set++
	case params.Get:
		ConnectionsCount.Get++
	case params.Del:
		ConnectionsCount.Del++
	case params.RPush:
		ConnectionsCount.RPush++
	case params.LRange:
		ConnectionsCount.LRange++
	case params.LTrim:
		ConnectionsCount.LTrim++
	}
}

func init() {
	ConnectionsCount = &ConnectionsCounter{}
}

func GetConnectionsInfo() (string, error) {
	raw, err := json.Marshal(ConnectionsCount)
	if err != nil {
		return "", fmt.Errorf("marshaling ConnectionsCounter failed: %w", err)
	}

	return string(raw), err
}

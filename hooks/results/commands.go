package results

import (
	"encoding/json"
	"fmt"
)

type addresses map[string]uint32

type operationsCounter struct {
	Read  addresses `json:"read"`
	Write addresses `json:"write"`
}

var oCounter *operationsCounter

func init() {
	oCounter = &operationsCounter{
		Read:  make(addresses),
		Write: make(addresses),
	}
}

func AddRead(address string) {
	if _, ok := oCounter.Read[address]; !ok {
		oCounter.Read[address] = 0
	}

	oCounter.Read[address]++
}

func AddWrite(address string) {
	if _, ok := oCounter.Write[address]; !ok {
		oCounter.Write[address] = 0
	}

	oCounter.Write[address]++
}

func GetOperationsCounter() (string, error) {
	raw, err := json.Marshal(oCounter)
	if err != nil {
		return "", fmt.Errorf("marshal failed: %w", err)
	}

	return string(raw), nil
}

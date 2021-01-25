package results

import (
	"encoding/json"
	"fmt"
	"sync"
)

type addresses map[string]uint32

type operationsCounter struct {
	my    *sync.Mutex
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

func (c *operationsCounter) addRead(address string) {
	c.my.Lock()
	if _, ok := oCounter.Read[address]; !ok {
		oCounter.Read[address] = 0
	}

	oCounter.Read[address]++
	c.my.Unlock()
}

func AddRead(address string) {
	oCounter.addRead(address)
}

func (c *operationsCounter) addWrite(address string) {
	c.my.Lock()
	if _, ok := oCounter.Write[address]; !ok {
		oCounter.Write[address] = 0
	}

	oCounter.Write[address]++
	c.my.Unlock()
}

func AddWrite(address string) {
	oCounter.addWrite(address)
}

func GetOperationsCounter() (string, error) {
	raw, err := json.Marshal(oCounter)
	if err != nil {
		return "", fmt.Errorf("marshal failed: %w", err)
	}

	return string(raw), nil
}

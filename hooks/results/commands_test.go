package results

import (
	"fmt"
	"testing"
)

func TestAddRead(t *testing.T) {
	AddRead("123")
	AddRead("123")
	AddRead("123")
	AddRead("123")
	AddRead("123")
	AddRead("123")
	AddRead("123")

	fmt.Println(oCounter.Read)
	fmt.Println(oCounter.Write)
}

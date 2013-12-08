package httputil2

import (
	"testing"
	"fmt"
)

func TestIdHandlerRandomGenerator(t *testing.T) {
	g := RandomGenerator(32)
	id := g()
	if len(id) != 32 {
		fmt.Println(len(id))
		t.Fail()
	}
}

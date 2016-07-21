package httputil2

import (
	"fmt"
	"testing"
)

func TestRequestIDRandomGenerator(t *testing.T) {
	g := RandomIDGenerator(32)
	id := g()
	if len(id) != 32 {
		fmt.Println(len(id))
		t.Fail()
	}
}

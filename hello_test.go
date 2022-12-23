package hello

import (
	"testing"
	"fmt"
)

func TestHello( t *testing.T) {
	
	expected := "Hello, World"
	
	if ret := Hello(); ret != expected {
		t.Error(fmt.Sprintf("Hello() = %q, want %q", ret, expected))
	}
	
}

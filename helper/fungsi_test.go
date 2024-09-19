package helper

import (
	"fmt"
	"testing"
)

func TestSayHello(t *testing.T) {
	result := SayHello("Dhimas")
	if result != "Hallo Dhimas" {
		// Error
		t.FailNow()
	}
}

func TestPenambahan(t *testing.T) {
	result := Penambahan(5.5, 6.0)
	if result != 11.5 {
		// Fail
		t.FailNow()
	}
	fmt.Printf("%.2F", result)
}

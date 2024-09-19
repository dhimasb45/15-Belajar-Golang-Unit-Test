package helper

import "testing"

func TestSayHello(t *testing.T) {
	result := SayHello("Dhimas")
	if result != "Hallo Dhimas" {
		// Error
		panic("Hasil bukan = 'Hallo Dhimas'")
	}
}

package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSayHello(t *testing.T) {
	result := SayHello("Dhimas")
	if result != "Hallo Dhimas" {
		// Errorr
		t.Fatal("Seharusnya : Hallo Dhimas")
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

func TestPerkalian(t *testing.T) {
	hasil := Perkalian(4.0, 5.0)
	YangDiInginkan := 20.0

	assert.Equal(t, YangDiInginkan, hasil, "Hasil perkalian seharusnya : 20")
}

func TestPembagian(t *testing.T) {
	hasil := Pembagian(10.0, 2.0)
	YangDiInginkan := 5.0

	require.Equal(t, YangDiInginkan, hasil, "Hasil pembagian seharusnya : %.2f", YangDiInginkan)
}

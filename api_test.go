package pinata

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestClient_PinFile(t *testing.T) {
	client := New(
		DefaultNode,
		"",
		"",
		"",
	)
	_, err := client.PinFile("uploads/receipt-1636015649-potrait.jpeg")
	assert.Equal(t, err, nil)
}

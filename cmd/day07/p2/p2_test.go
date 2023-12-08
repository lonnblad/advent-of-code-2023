package p2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_apa(t *testing.T) {
	h := hand{hand: "3K2TJ"}
	actualVal := h.value()
	assert.Equal(t, 2, actualVal)
}

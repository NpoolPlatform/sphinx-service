package fil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecomposeStringInt(t *testing.T) {
	amountInt, amountDigits, amountString := DecomposeStringInt("11364195385307179586438")
	assert.Equal(t, int64(11364195385307), amountInt)
	assert.Equal(t, int32(9), amountDigits)
	assert.Equal(t, "11364195385307.179586438"[0:18], amountString[0:18])
}

package fil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringValue2BigInt(t *testing.T) {
	bi := StringValue2BigInt("123456789.123456789123456789")
	assert.Equal(t, "123456789123456789123456789"[0:11], bi.String()[0:11])
	bi = StringValue2BigInt("123456789.123")
	assert.Equal(t, "123456789123000000000000000"[0:11], bi.String()[0:11])
}

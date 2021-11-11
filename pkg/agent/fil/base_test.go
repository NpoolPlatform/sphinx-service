package fil

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecomposeStringInt(t *testing.T) {
	amountInt, amountDigits, amountString := DecomposeStringInt("11364195385307179586438")
	assert.Equal(t, int64(11364195385307), amountInt)
	assert.Equal(t, int32(9), amountDigits)
	assert.Equal(t, "11364195385307.179586438"[0:18], amountString[0:18])
}

func TestSetHostWithToken(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	SetHostWithToken("172.16.30.117", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.ppK_nggwygh6kCPDlktdBtkGaqQXxoXM99iNx3-tZ9E")
}

func TestDecomposeStringFloat64(t *testing.T) {
	amountFloat64 := DecomposeStringFloat64("11364195385307179586438")
	assert.Equal(t, "11364.195385307179586438"[0:12], fmt.Sprintf("%v", amountFloat64)[0:12])
}

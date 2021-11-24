package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"
)

func TestPatchGRPCError(t *testing.T) {
	// nil to nil
	assert.Nil(t, PatchGRPCError(nil, "anything"))
	// be able to handle wrapped error
	tmpErr := xerrors.Errorf("some package %v",
		xerrors.New("some info"),
	)
	assert.NotEqual(t, tmpErr, PatchGRPCError(tmpErr, "UT in process"))
}

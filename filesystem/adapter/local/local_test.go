package local

import (
	"testing"

	"github.com/deatil/go-filesystem/filesystem/interfaces"
)

func Test_Local(t *testing.T) {
	local := new(Local)

	var _ interfaces.Adapter = local
}
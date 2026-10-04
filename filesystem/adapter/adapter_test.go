package adapter

import (
	"testing"

	"github.com/deatil/go-filesystem/filesystem/interfaces"
)

func Test_Adapter(t *testing.T) {
	ad := new(Adapter)

	var _ interfaces.Adapter = ad
}
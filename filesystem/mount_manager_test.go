package filesystem

import (
	"testing"

	local_adapter "github.com/deatil/go-filesystem/filesystem/adapter/local"
)

func getManager() *MountManager {
	root := "./testdata"
	adapter := local_adapter.New(root)

	fs := New(adapter)

	return NewMountManager().MountFilesystem("local", fs)
}

func Test_Manager_Read(t *testing.T) {
	assertEqual := assertEqualT(t)

	manager := getManager()

	res, err := manager.Read("local://test.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res), "testdata", "Test_Read")
}
package filesystem

import (
	"bytes"
	"io"
	"testing"

	local_adapter "github.com/deatil/go-filesystem/filesystem/adapter/local"
)

func getManager() *MountManager {
	root := "./testdata"
	adapter := local_adapter.New(root)

	fs := New(adapter)

	return NewMountManager().MountFilesystem("local", fs)
}

func getManager2() *MountManager {
	root := "./testdata"
	adapter := local_adapter.New(root)

	fs := New(adapter)

	return NewMountManager().MountFilesystems(map[string]*Filesystem{
		"local": fs,
		"local2": fs,
	})
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

func Test_Manager_ListContents(t *testing.T) {
	assertEqual := assertEqualT(t)

	manager := getManager()

	res, err := manager.ListContents("local://")
	if err != nil {
		t.Fatal(err.Error())
	}

	check := map[string]any{
		"path":      "test.txt",
		"size":      int64(8),
		"timestamp": int64(1733803713),
		"type":      "file",
	}

	useRes := map[string]any{}
	for _, v := range res {
		if path, ok := v["path"].(string); ok && path == "test.txt" {
			useRes = v
		}
	}

	assertEqual(useRes["path"], check["path"], "Test_Manager_ListContents path")
	assertEqual(useRes["size"], check["size"], "Test_Manager_ListContents size")
	assertEqual(useRes["type"], check["type"], "Test_Manager_ListContents type")

	if useRes["timestamp"].(int64) <= 0 {
		t.Errorf("timestamp get error, got %d", useRes["timestamp"])
	}
}

func Test_Manager_Copy(t *testing.T) {
	assertEqual := assertEqualT(t)

	manager := getManager()

	res, err := manager.Copy("local://testcopy.txt", "local://newtestcopy.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(res, true, "Test_Manager_Copy")

	res2 := manager.Has("local://newtestcopy.txt")
	assertEqual(res2, true, "Test_Manager_Copy Has")

	res3, _ := manager.Delete("local://newtestcopy.txt")
	assertEqual(res3, true, "Test_Manager_Copy Delete")

	res33 := manager.Has("local://newtestcopy.txt")
	assertEqual(res33, false, "Test_Manager_Copy Delete after Has")
}

func Test_Manager_Move(t *testing.T) {
	assertEqual := assertEqualT(t)

	manager := getManager2()

	{
		res, err := manager.Copy("local://testcopy.txt", "local://newtestcopy2.txt")
		if err != nil {
			t.Fatal(err.Error())
		}
	
		assertEqual(res, true, "Test_Manager_Move")
	
		res1, err := manager.Move("local://newtestcopy2.txt", "local://newtestcopy.txt")
		if err != nil {
			t.Fatal(err.Error())
		}
	
		assertEqual(res1, true, "Test_Manager_Move")
	
		res2 := manager.Has("local://newtestcopy.txt")
		assertEqual(res2, true, "Test_Manager_Move Has")
	
		res3, _ := manager.Delete("local://newtestcopy.txt")
		assertEqual(res3, true, "Test_Manager_Move Delete")
	
		res33 := manager.Has("local://newtestcopy.txt")
		assertEqual(res33, false, "Test_Manager_Move Delete after Has")	
	}

	{
		res, err := manager.Copy("local://testcopy.txt", "local://newtestcopy2.txt")
		if err != nil {
			t.Fatal(err.Error())
		}
	
		assertEqual(res, true, "Test_Manager_Move")
	
		res1, err := manager.Move("local://newtestcopy2.txt", "local2://newtestcopy.txt")
		if err != nil {
			t.Fatal(err.Error())
		}
	
		assertEqual(res1, true, "Test_Manager_Move")
	
		res2 := manager.Has("local2://newtestcopy.txt")
		assertEqual(res2, true, "Test_Manager_Move Has")
	
		res3, _ := manager.Delete("local2://newtestcopy.txt")
		assertEqual(res3, true, "Test_Manager_Move Delete")
	
		res33 := manager.Has("local2://newtestcopy.txt")
		assertEqual(res33, false, "Test_Manager_Move Delete after Has")	
	}
}

func Test_Manager_ReadStream(t *testing.T) {
	assertEqual := assertEqualT(t)

	manager := getManager()

	res, err := manager.ReadStream("local://test.txt")
	if err != nil {
		t.Fatal(err)
	}

	defer res.Close()

	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, res)
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(string(buf.Bytes()), "testdata", "Test_ReadStream")
}

func Test_Manager_GetMimetype(t *testing.T) {
	assertEqual := assertEqualT(t)

	manager := getManager()

	res, err := manager.GetMimetype("local://test.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(res, "application/octet-stream", "Test_Manager_GetMimetype")
}

func Test_Manager_GetTimestamp(t *testing.T) {
	manager := getManager()

	res, err := manager.GetTimestamp("local://test.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	if res <= 0 {
		t.Errorf("GetTimestamp() error, got %d", res)
	}
}

func Test_Manager_GetVisibility(t *testing.T) {
	manager := getManager()

	{
		res, err := manager.GetVisibility("local://test.txt")
		if err != nil {
			t.Fatal(err.Error())
		}
	
		if res != "666" && res != "public" {
			t.Error("GetVisibility fail")
		}
	}

	{
		_, err := manager.SetVisibility("local://test.txt", "private")
		if err != nil {
			t.Fatal(err)
		}

		res, err := manager.GetVisibility("local://test.txt")
		if err != nil {
			t.Fatal(err)
		}
	
		if res != "private" {
			t.Error("GetVisibility fail")
		}

		_, err = manager.SetVisibility("local://test.txt", "public")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func Test_Manager_GetSize(t *testing.T) {
	assertEqual := assertEqualT(t)

	manager := getManager()

	res, err := manager.GetSize("local://test.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(res, int64(8), "Test_Manager_GetSize")
}

func Test_Manager_GetMetadata(t *testing.T) {
	assertEqual := assertEqualT(t)

	manager := getManager()

	res, err := manager.GetMetadata("local://test.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	check := map[string]any{
		"path":      "test.txt",
		"size":      int64(8),
		"timestamp": int64(1733803713),
		"type":      "file",
	}

	assertEqual(res["path"], check["path"], "Test_GetMetadata path")
	assertEqual(res["size"], check["size"], "Test_GetMetadata size")
	assertEqual(res["type"], check["type"], "Test_GetMetadata type")

	if res["timestamp"].(int64) <= 0 {
		t.Errorf("timestamp get error, got %d", res["timestamp"])
	}
}

func Test_Manager_Write(t *testing.T) {
	assertEqual := assertEqualT(t)

	manager := getManager()

	ok, err := manager.Write("local://testcopy.txt", []byte("testtestdata1111111"))
	if !ok {
		t.Fatal(err.Error())
	}

	res2, err := manager.Read("local://testcopy.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res2), "testtestdata1111111", "Test_Write")

	ok, err = manager.Write("local://testcopy.txt", []byte("testdata"))
	if !ok {
		t.Fatal(err.Error())
	}
}

func Test_Manager_WriteStream(t *testing.T) {
	assertEqual := assertEqualT(t)

	manager := getManager()

	buf := bytes.NewBufferString("testtestdata1111111")

	ok, err := manager.WriteStream("local://testcopy.txt", buf)
	if !ok {
		t.Fatal(err.Error())
	}

	res2, err := manager.Read("local://testcopy.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res2), "testtestdata1111111", "Test_Write")

	ok, err = manager.Write("local://testcopy.txt", []byte("testdata"))
	if !ok {
		t.Fatal(err.Error())
	}
}

func Test_Manager_Update(t *testing.T) {
	assertEqual := assertEqualT(t)

	manager := getManager()

	ok, err := manager.Update("local://testcopy.txt", []byte("222222222"))
	if !ok {
		t.Fatal(err.Error())
	}

	res2, err := manager.Read("local://testcopy.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res2), "222222222", "Test_Update")

	ok, err = manager.Write("local://testcopy.txt", []byte("testdata"))
	if !ok {
		t.Fatal(err.Error())
	}
}

func Test_Manager_UpdateStream(t *testing.T) {
	assertEqual := assertEqualT(t)

	manager := getManager()

	buf := bytes.NewBufferString("222222222")

	ok, err := manager.UpdateStream("local://testcopy.txt", buf)
	if !ok {
		t.Fatal(err.Error())
	}

	res2, err := manager.Read("local://testcopy.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res2), "222222222", "Test_UpdateStream")

	ok, err = manager.Write("local://testcopy.txt", []byte("testdata"))
	if !ok {
		t.Fatal(err.Error())
	}
}

func Test_Manager_Put(t *testing.T) {
	assertEqual := assertEqualT(t)

	manager := getManager()

	ok, err := manager.Put("local://testcopy.txt", []byte("222222222"))
	if !ok {
		t.Fatal(err.Error())
	}

	res2, err := manager.Read("local://testcopy.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res2), "222222222", "Test_Put")

	ok, err = manager.Write("local://testcopy.txt", []byte("testdata"))
	if !ok {
		t.Fatal(err.Error())
	}
}

func Test_Manager_PutStream(t *testing.T) {
	assertEqual := assertEqualT(t)

	manager := getManager()

	buf := bytes.NewBufferString("222222222")

	ok, err := manager.PutStream("local://testcopy.txt", buf)
	if !ok {
		t.Fatal(err.Error())
	}

	res2, err := manager.Read("local://testcopy.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res2), "222222222", "Test_PutStream")

	ok, err = manager.Write("local://testcopy.txt", []byte("testdata"))
	if !ok {
		t.Fatal(err.Error())
	}
}

func Test_Manager_Rename(t *testing.T) {
	assertEqual := assertEqualT(t)

	manager := getManager()

	ok, err := manager.Rename("local://testcopy.txt", "testcopy222.txt")
	if !ok {
		t.Fatal(err.Error())
	}

	res2 := manager.Has("local://testcopy222.txt")
	assertEqual(res2, true, "Test_Rename")

	ok, err = manager.Rename("local://testcopy222.txt", "testcopy.txt")
	if !ok {
		t.Fatal(err.Error())
	}

	res3 := manager.Has("local://testcopy222.txt")
	assertEqual(res3, false, "Test_Rename Rename 1")

	res33 := manager.Has("local://testcopy.txt")
	assertEqual(res33, true, "Test_Rename Rename 2")

}

func Test_Manager_Create(t *testing.T) {
	assertEqual := assertEqualT(t)

	manager := getManager()

	res, err := manager.Create("local://newtestcopy.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(res, true, "Test_Manager_Create")

	res2 := manager.Has("local://newtestcopy.txt")
	assertEqual(res2, true, "Test_Manager_Create Has")

	res3, _ := manager.Delete("local://newtestcopy.txt")
	assertEqual(res3, true, "Test_Manager_Create Delete")

	res33 := manager.Has("local://newtestcopy.txt")
	assertEqual(res33, false, "Test_Manager_Create Delete after Has")
}

func Test_Manager_CreateDir(t *testing.T) {
	assertEqual := assertEqualT(t)

	manager := getManager()

	res, err := manager.CreateDir("local://create_testdir")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(res, true, "Test_Manager_CreateDir")

	res2 := manager.Has("local://create_testdir")
	assertEqual(res2, true, "Test_Manager_CreateDir Has")

	res3, _ := manager.DeleteDir("local://create_testdir")
	assertEqual(res3, true, "Test_Manager_CreateDir Delete")

	res33 := manager.Has("local://create_testdir")
	assertEqual(res33, false, "Test_Manager_CreateDir Delete after Has")
}

func Test_Manager_HasDir(t *testing.T) {
	assertEqual := assertEqualT(t)

	manager := getManager()

	res := manager.Has("local://testdir222")
	assertEqual(res, true, "Test_Manager_HasDir")

	res2 := manager.Has("local://testdir333")
	assertEqual(res2, false, "Test_Manager_HasDir 2")
}

func Test_Manager_ReadAndDelete(t *testing.T) {
	assertEqual := assertEqualT(t)

	manager := getManager()

	res, err := manager.Copy("local://testcopy.txt", "local://testReadAndDelete.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(res, true, "Test_ReadAndDelete")

	res2 := manager.Has("local://testReadAndDelete.txt")
	assertEqual(res2, true, "Test_ReadAndDelete Has")

	res3, err := manager.ReadAndDelete("local://testReadAndDelete.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res3), "testdata", "Test_ReadAndDelete ReadAndDelete")

	res33 := manager.Has("local://testReadAndDelete.txt")
	assertEqual(res33, false, "Test_ReadAndDelete ReadAndDelete after Has")
}

func Test_Manager_Get_Read(t *testing.T) {
	assertEqual := assertEqualT(t)

	manager := getManager()

	res, err := manager.Get("local://test.txt").(*File).Read()
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res), "testdata", "Test_Read")
}

func Test_Manager_Get_Read2(t *testing.T) {
	assertEqual := assertEqualT(t)

	manager := getManager()

	res := manager.Get("local://test.txt", func(f *Filesystem, path string) any {
		res2, _ := f.Read(path)
		return res2
	})

	assertEqual(string(res.([]byte)), "testdata", "Test_Manager_Get_Read2")
}

func Test_Manager_Get_ListContents(t *testing.T) {
	assertEqual := assertEqualT(t)

	manager := getManager()

	res, err := manager.Get("local://").(*Directory).GetContents()
	if err != nil {
		t.Fatal(err.Error())
	}

	check := map[string]any{
		"path":      "test.txt",
		"size":      int64(8),
		"timestamp": int64(1733803713),
		"type":      "file",
	}

	useRes := map[string]any{}
	for _, v := range res {
		if path, ok := v["path"].(string); ok && path == "test.txt" {
			useRes = v
		}
	}

	assertEqual(useRes["path"], check["path"], "Test_ListContents path")
	assertEqual(useRes["size"], check["size"], "Test_ListContents size")
	assertEqual(useRes["type"], check["type"], "Test_ListContents type")

	if useRes["timestamp"].(int64) <= 0 {
		t.Errorf("timestamp get error, got %d", useRes["timestamp"])
	}
}

func Test_Manager_FilterPrefix(t *testing.T) {
	assertEqual := assertEqualT(t)

	arguments := []string{
		"aa://111",
		"222",
		"333",
	}

	manager := getManager()
	prefix, argument := manager.FilterPrefix(arguments)

	assertEqual(prefix, "aa", "Test_Manager_FilterPrefix")
	assertEqual(argument, []string{
		"111",
		"222",
		"333",
	}, "Test_Manager_FilterPrefix")

}

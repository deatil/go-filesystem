package filesystem

import (
	"bytes"
	"io"
	"testing"

	local_adapter "github.com/deatil/go-filesystem/filesystem/adapter/local"
)

func getFS() *Filesystem {
	root := "./testdata"
	adapter := local_adapter.New(root)

	fs := New(adapter)
	return fs
}

func Test_File_With(t *testing.T) {
	assertEqual := assertEqualT(t)

	fs := getFS()

	path := "/test"

	{
		file := NewFile(fs)
		file.WithPath(path)

		assertEqual(path, file.path, "Test_File_With")
	}

	{
		file := NewFile(fs, path)

		assertEqual(path, file.path, "Test_File_With")
	}

	{
		file := NewFile(nil)
		file.WithFilesystem(fs)

		assertEqual(fs, file.filesystem, "Test_File_With")
	}

	{
		file := NewFile(fs)

		assertEqual(fs, file.filesystem, "Test_File_With")
	}
}

func Test_File_Exists(t *testing.T) {
	assertEqual := assertEqualT(t)

	fs := getFS()

	res := NewFile(fs, "/test.txt").Exists()
	assertEqual(res, true, "Test_File_Exists")

	res2 := NewFile(fs, "/test2.txt").Exists()
	assertEqual(res2, false, "Test_File_Exists 2")
}

func Test_File_Read(t *testing.T) {
	assertEqual := assertEqualT(t)

	fs := getFS()

	res, err := NewFile(fs, "/test.txt").Read()
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res), "testdata", "Test_File_Read")
}

func Test_File_ReadStream(t *testing.T) {
	assertEqual := assertEqualT(t)

	fs := getFS()

	res, err := NewFile(fs, "/test.txt").ReadStream()
	if err != nil {
		t.Fatal(err)
	}

	defer res.Close()

	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, res)
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(string(buf.Bytes()), "testdata", "Test_File_ReadStream")
}

func Test_File_GetMimetype(t *testing.T) {
	assertEqual := assertEqualT(t)

	fs := getFS()

	res, err := NewFile(fs, "/test.txt").GetMimetype()
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(res, "application/octet-stream", "Test_File_GetMimetype")
}

func Test_File_GetTimestamp(t *testing.T) {
	fs := getFS()

	res, err := NewFile(fs, "/test.txt").GetTimestamp()
	if err != nil {
		t.Fatal(err.Error())
	}

	if res <= 0 {
		t.Errorf("GetTimestamp() error, got %d", res)
	}
}

func Test_File_GetVisibility(t *testing.T) {
	fs := getFS()

	{
		res, err := NewFile(fs, "/test.txt").GetVisibility()
		if err != nil {
			t.Fatal(err.Error())
		}
	
		if res != "666" && res != "public" {
			t.Error("GetVisibility fail")
		}
	}

}

func Test_File_GetSize(t *testing.T) {
	assertEqual := assertEqualT(t)

	fs := getFS()

	res, err := NewFile(fs, "/test.txt").GetSize()
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(res, int64(8), "Test_File_GetSize")
}

func Test_File_GetMetadata(t *testing.T) {
	assertEqual := assertEqualT(t)

	fs := getFS()

	res, err := NewFile(fs, "/test.txt").GetMetadata()
	if err != nil {
		t.Fatal(err.Error())
	}

	check := map[string]any{
		"path":      "test.txt",
		"size":      int64(8),
		"timestamp": int64(1733803713),
		"type":      "file",
	}

	assertEqual(res["path"], check["path"], "Test_File_GetMetadata path")
	assertEqual(res["size"], check["size"], "Test_File_GetMetadata size")
	assertEqual(res["type"], check["type"], "Test_File_GetMetadata type")

	if res["timestamp"].(int64) <= 0 {
		t.Errorf("timestamp get error, got %d", res["timestamp"])
	}
}

func Test_File_Write(t *testing.T) {
	assertEqual := assertEqualT(t)

	fs := getFS()

	ok, err := NewFile(fs, "/testcopy.txt").Write([]byte("testtestdata1111111"))
	if !ok {
		t.Fatal(err.Error())
	}

	res2, err := NewFile(fs, "/testcopy.txt").Read()
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res2), "testtestdata1111111", "Test_Write")

	ok, err = NewFile(fs, "/testcopy.txt").Write([]byte("testdata"))
	if !ok {
		t.Fatal(err.Error())
	}
}

func Test_File_WriteStream(t *testing.T) {
	assertEqual := assertEqualT(t)

	fs := NewFile(getFS())

	buf := bytes.NewBufferString("testtestdata1111111")

	ok, err := fs.WithPath("/testcopy.txt").WriteStream(buf)
	if !ok {
		t.Fatal(err.Error())
	}

	res2, err := fs.WithPath("/testcopy.txt").Read()
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res2), "testtestdata1111111", "Test_Write")

	ok, err = fs.WithPath("/testcopy.txt").Write([]byte("testdata"))
	if !ok {
		t.Fatal(err.Error())
	}
}

func Test_File_Put(t *testing.T) {
	assertEqual := assertEqualT(t)

	fs := NewFile(getFS())

	ok, err := fs.WithPath("/testcopy.txt").Put([]byte("222222222"))
	if !ok {
		t.Fatal(err.Error())
	}

	res2, err := fs.WithPath("/testcopy.txt").Read()
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res2), "222222222", "Test_Put")

	ok, err = fs.WithPath("/testcopy.txt").Write([]byte("testdata"))
	if !ok {
		t.Fatal(err.Error())
	}
}

func Test_File_PutStream(t *testing.T) {
	assertEqual := assertEqualT(t)

	fs := NewFile(getFS())

	buf := bytes.NewBufferString("222222222")

	ok, err := fs.WithPath("/testcopy.txt").PutStream(buf)
	if !ok {
		t.Fatal(err.Error())
	}

	res2, err := fs.WithPath("/testcopy.txt").Read()
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res2), "222222222", "Test_PutStream")

	ok, err = fs.WithPath("/testcopy.txt").Write([]byte("testdata"))
	if !ok {
		t.Fatal(err.Error())
	}
}

func Test_File_Update(t *testing.T) {
	assertEqual := assertEqualT(t)

	fs := NewFile(getFS())

	ok, err := fs.WithPath("/testcopy.txt").Update([]byte("222222222"))
	if !ok {
		t.Fatal(err.Error())
	}

	res2, err := fs.WithPath("/testcopy.txt").Read()
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res2), "222222222", "Test_Update")

	ok, err = fs.WithPath("/testcopy.txt").Write([]byte("testdata"))
	if !ok {
		t.Fatal(err.Error())
	}
}

func Test_File_UpdateStream(t *testing.T) {
	assertEqual := assertEqualT(t)

	fs := NewFile(getFS())

	buf := bytes.NewBufferString("222222222")

	ok, err := fs.WithPath("/testcopy.txt").UpdateStream(buf)
	if !ok {
		t.Fatal(err.Error())
	}

	res2, err := fs.WithPath("/testcopy.txt").Read()
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res2), "222222222", "Test_UpdateStream")

	ok, err = fs.WithPath("/testcopy.txt").Write([]byte("testdata"))
	if !ok {
		t.Fatal(err.Error())
	}
}

func Test_File_Rename(t *testing.T) {
	assertEqual := assertEqualT(t)

	fs := NewFile(getFS())

	ok, err := fs.WithPath("/testcopy.txt").Rename("/testcopy222.txt")
	if !ok {
		t.Fatal(err.Error())
	}

	res2 := fs.WithPath("/testcopy222.txt").Exists()
	assertEqual(res2, true, "Test_Rename")

	ok, err = fs.WithPath("/testcopy222.txt").Rename("/testcopy.txt")
	if !ok {
		t.Fatal(err.Error())
	}

	res3 := fs.WithPath("/testcopy222.txt").Exists()
	assertEqual(res3, false, "Test_Rename Rename 1")

	res33 := fs.WithPath("/testcopy.txt").Exists()
	assertEqual(res33, true, "Test_Rename Rename 2")

}

func Test_File_Copy(t *testing.T) {
	assertEqual := assertEqualT(t)

	fs := NewFile(getFS())

	res, err := fs.WithPath("/testcopy.txt").Copy("/newtestcopy.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	res2 := res.WithPath("/newtestcopy.txt").Exists()
	assertEqual(res2, true, "Test_Copy Exists")

	res3, _ := res.WithPath("/newtestcopy.txt").Delete()
	assertEqual(res3, true, "Test_Copy Delete")

	res33 := res.WithPath("/newtestcopy.txt").Exists()
	assertEqual(res33, false, "Test_Copy Delete after Exists")
}

func Test_File_Create(t *testing.T) {
	assertEqual := assertEqualT(t)

	fs := NewFile(getFS())

	_, err := fs.WithPath("/newtestcopy.txt").Create()
	if err != nil {
		t.Fatal(err.Error())
	}

	res2 := fs.WithPath("/newtestcopy.txt").Exists()
	assertEqual(res2, true, "Test_File_Create Exists")

	res3, _ := fs.WithPath("/newtestcopy.txt").Delete()
	assertEqual(res3, true, "Test_File_Create Delete")

	res33 := fs.WithPath("/newtestcopy.txt").Exists()
	assertEqual(res33, false, "Test_File_Create Delete after Exists")
}

func Test_File_Is(t *testing.T) {
	assertEqual := assertEqualT(t)

	f := getFS()
	fs := NewFile(f)

	res := fs.WithPath("/test.txt")

	assertEqual(res.IsDir(), false, "Test_File_Is")
	assertEqual(res.IsFile(), true, "Test_File_Is")
	assertEqual(res.GetType(), "file", "Test_File_Is")
	assertEqual(res.GetFilesystem(), f, "Test_File_Is")
	assertEqual(res.GetPath(), "/test.txt", "Test_File_Is")

	fs2 := NewFile(nil)
	fs2.SetFilesystem(f)
	fs2.SetPath("/test.txt")
	assertEqual(fs2.filesystem, f, "Test_File_Is")
	assertEqual(fs2.path, "/test.txt", "Test_File_Is")
}



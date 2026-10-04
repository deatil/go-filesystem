package filesystem

import (
	"bytes"
	"io"
	"reflect"
	"testing"

	local_adapter "github.com/deatil/go-filesystem/filesystem/adapter/local"
	"github.com/deatil/go-filesystem/filesystem/config"
)

func assertErrorT(t *testing.T) func(error, string) {
	return func(err error, msg string) {
		if err != nil {
			t.Errorf("Failed %s: error: %+v", msg, err)
		}
	}
}

func assertEqualT(t *testing.T) func(any, any, string) {
	return func(actual any, expected any, msg string) {
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Failed %s: actual: %v, expected: %v", msg, actual, expected)
		}
	}
}

func Test_With(t *testing.T) {
	assertEqual := assertEqualT(t)

	settings := map[string]any{
		"test": "data",
	}
	conf := config.New(settings)

	root := "./testdata"
	adapter := local_adapter.New(root)

	{
		fs := New(adapter)
		fs.WithConfig(conf)

		assertEqual(conf, fs.config, "Test_WithConfig")
	}

	{
		fs := New(adapter, settings)

		assertEqual(conf, fs.config, "Test_WithConfig")
		assertEqual(conf, fs.GetConfig(), "Test_WithConfig")
	}

	{
		fs := New(nil)
		fs.WithAdapter(adapter)

		assertEqual(adapter, fs.adapter, "Test_WithConfig")
	}

	{
		fs := New(adapter)

		assertEqual(adapter, fs.adapter, "Test_WithConfig")
		assertEqual(adapter, fs.GetAdapter(), "Test_WithConfig")
	}
}

func Test_ListContents(t *testing.T) {
	assertEqual := assertEqualT(t)

	// 根目录
	root := "./testdata"
	adapter := local_adapter.New(root)

	// 磁盘
	fs := New(adapter)

	// 使用
	res, err := fs.ListContents("/")
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

func Test_Has(t *testing.T) {
	assertEqual := assertEqualT(t)

	// 根目录
	root := "./testdata"
	adapter := local_adapter.New(root)

	fs := New(adapter)

	res := fs.Has("/test.txt")
	assertEqual(res, true, "Test_Has")

	res2 := fs.Has("/test2.txt")
	assertEqual(res2, false, "Test_Has 2")
}

func Test_Read(t *testing.T) {
	assertEqual := assertEqualT(t)

	// 根目录
	root := "./testdata"
	adapter := local_adapter.New(root)

	fs := New(adapter)

	res, err := fs.Read("/test.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res), "testdata", "Test_Read")
}

func Test_GetMimetype(t *testing.T) {
	assertEqual := assertEqualT(t)

	// 根目录
	root := "./testdata"
	adapter := local_adapter.New(root)

	fs := New(adapter)

	res, err := fs.GetMimetype("/test.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(res, "application/octet-stream", "Test_GetMimetype")
}

func Test_GetTimestamp(t *testing.T) {
	// 根目录
	root := "./testdata"
	adapter := local_adapter.New(root)

	fs := New(adapter)

	res, err := fs.GetTimestamp("/test.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	if res <= 0 {
		t.Errorf("GetTimestamp() error, got %d", res)
	}
}

func Test_GetVisibility(t *testing.T) {
	// 根目录
	root := "./testdata"
	adapter := local_adapter.New(root)

	fs := New(adapter)

	res, err := fs.GetVisibility("/test.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	if res != "666" && res != "public" {
		t.Error("GetVisibility fail")
	}
}

func Test_GetSize(t *testing.T) {
	assertEqual := assertEqualT(t)

	// 根目录
	root := "./testdata"
	adapter := local_adapter.New(root)

	fs := New(adapter)

	res, err := fs.GetSize("/test.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(res, int64(8), "Test_GetSize")
}

func Test_GetMetadata(t *testing.T) {
	assertEqual := assertEqualT(t)

	// 根目录
	root := "./testdata"
	adapter := local_adapter.New(root)

	// 磁盘
	fs := New(adapter)

	// 使用
	res, err := fs.GetMetadata("/test.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	check := map[string]any{
		"path":      "test.txt",
		"size":      int64(8),
		"timestamp": int64(1733803713),
		"type":      "file",
	}

	assertEqual(res["path"], check["path"], "Test_ListContents path")
	assertEqual(res["size"], check["size"], "Test_ListContents size")
	assertEqual(res["type"], check["type"], "Test_ListContents type")

	if res["timestamp"].(int64) <= 0 {
		t.Errorf("timestamp get error, got %d", res["timestamp"])
	}
}

func Test_Write(t *testing.T) {
	assertEqual := assertEqualT(t)

	// 根目录
	root := "./testdata"
	adapter := local_adapter.New(root)

	// 磁盘
	fs := New(adapter)

	// 使用
	ok, err := fs.Write("/testcopy.txt", []byte("testtestdata1111111"))
	if !ok {
		t.Fatal(err.Error())
	}

	res2, err := fs.Read("/testcopy.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res2), "testtestdata1111111", "Test_Write")

	// 使用
	ok, err = fs.Write("/testcopy.txt", []byte("testdata"))
	if !ok {
		t.Fatal(err.Error())
	}
}

func Test_WriteStream(t *testing.T) {
	assertEqual := assertEqualT(t)

	// 根目录
	root := "./testdata"
	adapter := local_adapter.New(root)

	// 磁盘
	fs := New(adapter)

	buf := bytes.NewBufferString("testtestdata1111111")

	// 使用
	ok, err := fs.WriteStream("/testcopy.txt", buf)
	if !ok {
		t.Fatal(err.Error())
	}

	res2, err := fs.Read("/testcopy.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res2), "testtestdata1111111", "Test_Write")

	// 使用
	ok, err = fs.Write("/testcopy.txt", []byte("testdata"))
	if !ok {
		t.Fatal(err.Error())
	}
}

func Test_Put(t *testing.T) {
	assertEqual := assertEqualT(t)

	// 根目录
	root := "./testdata"
	adapter := local_adapter.New(root)

	// 磁盘
	fs := New(adapter)

	// 使用
	ok, err := fs.Put("/testcopy.txt", []byte("222222222"))
	if !ok {
		t.Fatal(err.Error())
	}

	res2, err := fs.Read("/testcopy.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res2), "222222222", "Test_Put")

	// 使用
	ok, err = fs.Write("/testcopy.txt", []byte("testdata"))
	if !ok {
		t.Fatal(err.Error())
	}
}

func Test_Prepend(t *testing.T) {
	assertEqual := assertEqualT(t)

	// 根目录
	root := "./testdata"
	adapter := local_adapter.New(root)

	// 磁盘
	fs := New(adapter)

	// 使用
	ok, err := fs.Prepend("/testcopy.txt", []byte("222222222"))
	if !ok {
		t.Fatal(err.Error())
	}

	res2, err := fs.Read("/testcopy.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res2), "222222222testdata", "Test_Prepend")

	// 使用
	ok, err = fs.Write("/testcopy.txt", []byte("testdata"))
	if !ok {
		t.Fatal(err.Error())
	}
}

func Test_PrependStream(t *testing.T) {
	assertEqual := assertEqualT(t)

	// 根目录
	root := "./testdata"
	adapter := local_adapter.New(root)

	// 磁盘
	fs := New(adapter)

	prependDdata := bytes.NewBuffer([]byte("222222222"))

	// 使用
	ok, err := fs.PrependStream("/testcopy.txt", prependDdata)
	if !ok {
		t.Fatal(err.Error())
	}

	res2, err := fs.Read("/testcopy.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res2), "222222222testdata", "Test_PrependStream")

	// 使用
	ok, err = fs.Write("/testcopy.txt", []byte("testdata"))
	if !ok {
		t.Fatal(err.Error())
	}
}

func Test_Append(t *testing.T) {
	assertEqual := assertEqualT(t)

	// 根目录
	root := "./testdata"
	adapter := local_adapter.New(root)

	// 磁盘
	fs := New(adapter)

	// 使用
	ok, err := fs.Append("/testcopy.txt", []byte("222222222"))
	if !ok {
		t.Fatal(err.Error())
	}

	res2, err := fs.Read("/testcopy.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res2), "testdata222222222", "Test_Append")

	// 使用
	ok, err = fs.Write("/testcopy.txt", []byte("testdata"))
	if !ok {
		t.Fatal(err.Error())
	}
}

func Test_AppendStream(t *testing.T) {
	assertEqual := assertEqualT(t)

	// 根目录
	root := "./testdata"
	adapter := local_adapter.New(root)

	// 磁盘
	fs := New(adapter)

	appendData := bytes.NewBuffer([]byte("222222222"))

	// 使用
	ok, err := fs.AppendStream("/testcopy.txt", appendData)
	if !ok {
		t.Fatal(err.Error())
	}

	res2, err := fs.Read("/testcopy.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res2), "testdata222222222", "Test_AppendStream")

	// 使用
	ok, err = fs.Write("/testcopy.txt", []byte("testdata"))
	if !ok {
		t.Fatal(err.Error())
	}
}

func Test_Rename(t *testing.T) {
	assertEqual := assertEqualT(t)

	// 根目录
	root := "./testdata"
	adapter := local_adapter.New(root)

	// 磁盘
	fs := New(adapter)

	// 使用
	ok, err := fs.Rename("/testcopy.txt", "/testcopy222.txt")
	if !ok {
		t.Fatal(err.Error())
	}

	res2 := fs.Has("/testcopy222.txt")
	assertEqual(res2, true, "Test_Rename")

	// 使用
	ok, err = fs.Rename("/testcopy222.txt", "/testcopy.txt")
	if !ok {
		t.Fatal(err.Error())
	}

	res3 := fs.Has("/testcopy222.txt")
	assertEqual(res3, false, "Test_Rename Rename 1")

	res33 := fs.Has("/testcopy.txt")
	assertEqual(res33, true, "Test_Rename Rename 2")

}

func Test_Copy(t *testing.T) {
	assertEqual := assertEqualT(t)

	// 根目录
	root := "./testdata"
	adapter := local_adapter.New(root)

	fs := New(adapter)

	res, err := fs.Copy("/testcopy.txt", "/newtestcopy.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(res, true, "Test_Copy")

	res2 := fs.Has("/newtestcopy.txt")
	assertEqual(res2, true, "Test_Copy Has")

	res3, _ := fs.Delete("/newtestcopy.txt")
	assertEqual(res3, true, "Test_Copy Delete")

	res33 := fs.Has("/newtestcopy.txt")
	assertEqual(res33, false, "Test_Copy Delete after Has")
}

func Test_CreateDir(t *testing.T) {
	assertEqual := assertEqualT(t)

	// 根目录
	root := "./testdata"
	adapter := local_adapter.New(root)

	fs := New(adapter)

	res, err := fs.CreateDir("/testdir")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(res, true, "Test_CreateDir")

	res2 := fs.Has("/testdir")
	assertEqual(res2, true, "Test_CreateDir Has")

	res3, _ := fs.DeleteDir("/testdir")
	assertEqual(res3, true, "Test_CreateDir Delete")

	res33 := fs.Has("/testdir")
	assertEqual(res33, false, "Test_CreateDir Delete after Has")
}

func Test_HasDir(t *testing.T) {
	assertEqual := assertEqualT(t)

	// 根目录
	root := "./testdata"
	adapter := local_adapter.New(root)

	fs := New(adapter)

	res := fs.Has("/testdir222")
	assertEqual(res, true, "Test_HasDir")

	res2 := fs.Has("/testdir333")
	assertEqual(res2, false, "Test_HasDir 2")
}

func Test_ReadAndDelete(t *testing.T) {
	assertEqual := assertEqualT(t)

	// 根目录
	root := "./testdata"
	adapter := local_adapter.New(root)

	fs := New(adapter)

	res, err := fs.Copy("/testcopy.txt", "/testReadAndDelete.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(res, true, "Test_ReadAndDelete")

	res2 := fs.Has("/testReadAndDelete.txt")
	assertEqual(res2, true, "Test_ReadAndDelete Has")

	res3, err := fs.ReadAndDelete("/testReadAndDelete.txt")
	if err != nil {
		t.Fatal(err.Error())
	}

	assertEqual(string(res3), "testdata", "Test_ReadAndDelete ReadAndDelete")

	res33 := fs.Has("/testReadAndDelete.txt")
	assertEqual(res33, false, "Test_ReadAndDelete ReadAndDelete after Has")
}

func Test_ReadStream(t *testing.T) {
	assertEqual := assertEqualT(t)

	// 根目录
	root := "./testdata"
	adapter := local_adapter.New(root)

	// 磁盘
	fs := New(adapter)

	res, err := fs.ReadStream("/test.txt")
	if err != nil {
		t.Fatal(err)
	}

	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, res)
	if err != nil {
		t.Fatal(err)
	}

	assertEqual(string(buf.Bytes()), "testdata", "Test_ReadStream")
}

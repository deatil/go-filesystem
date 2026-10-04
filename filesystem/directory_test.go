package filesystem

import (
	"testing"
)

func Test_Directory_With(t *testing.T) {
	assertEqual := assertEqualT(t)

	fs := getFS()

	path := "/test"

	{
		file := NewDirectory(fs)
		file.WithPath(path)

		assertEqual(path, file.path, "Test_Directory_With")
	}

	{
		file := NewDirectory(fs, path)

		assertEqual(path, file.path, "Test_Directory_With")
	}

	{
		file := NewDirectory(nil)
		file.WithFilesystem(fs)

		assertEqual(fs, file.filesystem, "Test_Directory_With")
	}

	{
		file := NewDirectory(fs)

		assertEqual(fs, file.filesystem, "Test_Directory_With")
	}
}

func Test_Directory_Exists(t *testing.T) {
	assertEqual := assertEqualT(t)

	fs := NewDirectory(getFS())

	res1 := fs.WithPath("/test_dir").Exists()
	assertEqual(res1, true, "Test_Directory_Exists")

	res2 := fs.WithPath("/test_dir222").Exists()
	assertEqual(res2, false, "Test_Directory_Exists")
}

func Test_Directory_Rename(t *testing.T) {
	assertEqual := assertEqualT(t)

	fs := NewDirectory(getFS())

	ok, err := fs.WithPath("/test_dir").Rename("/test_dir222")
	if !ok {
		t.Fatal(err.Error())
	}

	res2 := fs.WithPath("/test_dir222").Exists()
	assertEqual(res2, true, "Test_Rename")

	ok, err = fs.WithPath("/test_dir222").Rename("/test_dir")
	if !ok {
		t.Fatal(err.Error())
	}

	res3 := fs.WithPath("/test_dir222").Exists()
	assertEqual(res3, false, "Test_Rename Rename 1")

	res33 := fs.WithPath("/test_dir").Exists()
	assertEqual(res33, true, "Test_Rename Rename 2")

}

func Test_Directory_Copy(t *testing.T) {
	assertEqual := assertEqualT(t)

	fs := NewDirectory(getFS())

	_, err := fs.WithPath("/test_dir").Copy("/newtest_dir")
	if err != nil {
		t.Fatal(err)
	}

	res2 := fs.WithPath("/newtest_dir").Exists()
	assertEqual(res2, true, "Test_Directory_Copy Exists")

	res3, err := fs.WithPath("/newtest_dir").Delete()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(res3, true, "Test_Directory_Copy Delete")

	res33 := fs.WithPath("/newtest_dir").Exists()
	assertEqual(res33, false, "Test_Directory_Copy Delete after Exists")
}

func Test_Directory_Create(t *testing.T) {
	assertEqual := assertEqualT(t)

	fs := NewDirectory(getFS())

	_, err := fs.WithPath("/create_test_dir").Create()
	if err != nil {
		t.Fatal(err)
	}

	res2 := fs.WithPath("/create_test_dir").Exists()
	assertEqual(res2, true, "Test_Directory_Create Exists")

	res3, err := fs.WithPath("/create_test_dir").Delete()
	if err != nil {
		t.Fatal(err)
	}
	assertEqual(res3, true, "Test_Directory_Create Delete")

	res33 := fs.WithPath("/create_test_dir").Exists()
	assertEqual(res33, false, "Test_Directory_Create Delete after Exists")
}

func Test_Directory_ListContents(t *testing.T) {
	assertEqual := assertEqualT(t)

	fs := NewDirectory(getFS())

	res, err := fs.WithPath("/").GetContents()
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

	assertEqual(useRes["path"], check["path"], "Test_Directory_ListContents path")
	assertEqual(useRes["size"], check["size"], "Test_Directory_ListContents size")
	assertEqual(useRes["type"], check["type"], "Test_Directory_ListContents type")

	if useRes["timestamp"].(int64) <= 0 {
		t.Errorf("timestamp get error, got %d", useRes["timestamp"])
	}
}

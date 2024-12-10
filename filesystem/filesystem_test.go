package filesystem_test

import (
    "reflect"
    "testing"

    "github.com/deatil/go-filesystem/filesystem"
    local_adapter "github.com/deatil/go-filesystem/filesystem/adapter/local"
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

func Test_ListContents(t *testing.T) {
    assertEqual := assertEqualT(t)

    // 根目录
    root := "./testdata"
    adapter := local_adapter.New(root)

    // 磁盘
    fs := filesystem.New(adapter)

    // 使用
    res, err := fs.ListContents("/")
    if err != nil {
        t.Fatal(err.Error())
    }

    check := map[string]any{
        "path": "test.txt",
        "size": int64(8),
        "timestamp": int64(1733803713),
        "type": "file",
    }

    useRes := map[string]any{}
    for _, v := range res {
        if path, ok := v["path"].(string); ok && path == "test.txt" {
            useRes = v
        }
    }

    assertEqual(useRes, check, "Test_ListContents")
}

func Test_Has(t *testing.T) {
    assertEqual := assertEqualT(t)

    // 根目录
    root := "./testdata"
    adapter := local_adapter.New(root)

    fs := filesystem.New(adapter)

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

    fs := filesystem.New(adapter)

    res, err := fs.Read("/test.txt")
    if err != nil {
        t.Fatal(err.Error())
    }

    assertEqual(res, "testdata", "Test_Read")
}

func Test_GetMimetype(t *testing.T) {
    assertEqual := assertEqualT(t)

    // 根目录
    root := "./testdata"
    adapter := local_adapter.New(root)

    fs := filesystem.New(adapter)

    res, err := fs.GetMimetype("/test.txt")
    if err != nil {
        t.Fatal(err.Error())
    }

    assertEqual(res, "application/octet-stream", "Test_GetMimetype")
}

func Test_GetTimestamp(t *testing.T) {
    assertEqual := assertEqualT(t)

    // 根目录
    root := "./testdata"
    adapter := local_adapter.New(root)

    fs := filesystem.New(adapter)

    res, err := fs.GetTimestamp("/test.txt")
    if err != nil {
        t.Fatal(err.Error())
    }

    assertEqual(res, int64(1733803713), "Test_GetTimestamp")
}

func Test_GetVisibility(t *testing.T) {
    assertEqual := assertEqualT(t)

    // 根目录
    root := "./testdata"
    adapter := local_adapter.New(root)

    fs := filesystem.New(adapter)

    res, err := fs.GetVisibility("/test.txt")
    if err != nil {
        t.Fatal(err.Error())
    }

    assertEqual(res, "666", "Test_GetTimestamp")
}

func Test_GetSize(t *testing.T) {
    assertEqual := assertEqualT(t)

    // 根目录
    root := "./testdata"
    adapter := local_adapter.New(root)

    fs := filesystem.New(adapter)

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
    fs := filesystem.New(adapter)

    // 使用
    res, err := fs.GetMetadata("/test.txt")
    if err != nil {
        t.Fatal(err.Error())
    }

    check := map[string]any{
        "path": "test.txt",
        "size": int64(8),
        "timestamp": int64(1733803713),
        "type": "file",
    }

    assertEqual(res, check, "Test_GetMetadata")
}

func Test_Write(t *testing.T) {
    assertEqual := assertEqualT(t)

    // 根目录
    root := "./testdata"
    adapter := local_adapter.New(root)

    // 磁盘
    fs := filesystem.New(adapter)

    // 使用
    ok, err := fs.Write("/testcopy.txt", "testtestdata1111111")
    if !ok {
        t.Fatal(err.Error())
    }

    res2, err := fs.Read("/testcopy.txt")
    if err != nil {
        t.Fatal(err.Error())
    }

    assertEqual(res2, "testtestdata1111111", "Test_Write")

    // 使用
    ok, err = fs.Write("/testcopy.txt", "testdata")
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
    fs := filesystem.New(adapter)

    // 使用
    ok, err := fs.Put("/testcopy.txt", "222222222")
    if !ok {
        t.Fatal(err.Error())
    }

    res2, err := fs.Read("/testcopy.txt")
    if err != nil {
        t.Fatal(err.Error())
    }

    assertEqual(res2, "222222222", "Test_Put")

    // 使用
    ok, err = fs.Write("/testcopy.txt", "testdata")
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
    fs := filesystem.New(adapter)

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

    fs := filesystem.New(adapter)

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

    fs := filesystem.New(adapter)

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

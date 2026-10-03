## Go-filesystem

A filesystem library for golang.


### Adapter

*  `local`: local storage


### Download

~~~go
go get -u github.com/deatil/go-filesystem
~~~


### Get Starting

~~~go
import (
    "fmt"

    "github.com/deatil/go-filesystem/filesystem"
    local_adapter "github.com/deatil/go-filesystem/filesystem/adapter/local"
)

func main() {
    root := "/storage"
    adapter := local_adapter.New(root)

    fs := filesystem.New(adapter)

    path := "/path.txt"
    contents := []byte("testdata")

    ok, err := fs.Write(path, contents)
    if err != nil {
        fmt.Println(err.Error())
    }
}
~~~


### Functions

~~~go
Write(path, contents []byte) (bool, error)

WriteStream(path string, resource io.Reader) (bool, error)

Put(path, contents []byte) (bool, error)

PutStream(path string, resource io.Reader) (bool, error)

ReadAndDelete(path string) (any, error)

Update(path, contents []byte) (bool, error)

Read(path string) ([]byte, error)

Rename(path, newpath string) (bool, error)

Copy(path, newpath string) (bool, error)

Delete(path string) (bool, error)

DeleteDir(dirname string) (bool, error)

CreateDir(dirname string) (bool, error)

ListContents(dirname string) ([]map[string]any, error)
~~~


### LICENSE

*  The library LICENSE is `Apache2`, using the library need keep the LICENSE.


### Copyright

*  Copyright deatil(https://github.com/deatil).


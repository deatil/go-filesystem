package filesystem

import (
	"fmt"
	"io"
	"strings"
)

/**
 * MountManager
 *
 * @create 2021-8-7
 * @author deatil
 */
type MountManager struct {
	filesystems map[string]*Filesystem
}

func NewMountManager(filesystems ...map[string]any) *MountManager {
	mm := &MountManager{
		filesystems: make(map[string]*Filesystem),
	}

	if len(filesystems) > 0 {
		mm.MountFilesystems(filesystems[0])
	}

	return mm
}

func (this *MountManager) MountFilesystems(filesystems map[string]any) *MountManager {
	for prefix, filesystem := range filesystems {
		this.MountFilesystem(prefix, filesystem.(*Filesystem))
	}

	return this
}

func (this *MountManager) MountFilesystem(prefix string, filesystem *Filesystem) *MountManager {
	this.filesystems[prefix] = filesystem

	return this
}

func (this *MountManager) GetFilesystem(prefix string) *Filesystem {
	if _, ok := this.filesystems[prefix]; !ok {
		panic(fmt.Sprintf("go-filesystem: [%s] prefix not exists", prefix))
	}

	return this.filesystems[prefix]
}

// [:prefix, :arguments]
func (this *MountManager) FilterPrefix(arguments []string) (string, []string) {
	if len(arguments) < 1 {
		panic("go-filesystem: arguments slice not empty")
	}

	path := arguments[0]

	prefix, path := this.GetPrefixAndPath(path)

	newArguments := make([]string, len(arguments))
	newArguments = append(newArguments, path)
	newArguments = append(newArguments, arguments[1:]...)

	return prefix, newArguments
}

// [:prefix, :path]
func (this *MountManager) GetPrefixAndPath(path string) (string, string) {
	paths := strings.SplitN(path, "://", 2)

	if len(paths) < 1 {
		panic(fmt.Sprintf("go-filesystem: [%s] prefix not exists", path))
	}

	return paths[0], paths[1]
}

func (this *MountManager) ListContents(directory string, recursive ...bool) ([]map[string]any, error) {
	prefix, dir := this.GetPrefixAndPath(directory)

	filesystem := this.GetFilesystem(prefix)

	result, err := filesystem.ListContents(dir, recursive...)
	if err != nil {
		return nil, err
	}

	for key, item := range result {
		item["filesystem"] = prefix
		result[key] = item
	}

	return result, nil
}

func (this *MountManager) Copy(from string, to string, conf ...map[string]any) (bool, error) {
	prefixFrom, pathFrom := this.GetPrefixAndPath(from)

	buffer, err := this.GetFilesystem(prefixFrom).ReadStream(pathFrom)
	if err != nil {
		return false, err
	}

	defer buffer.Close()

	prefixTo, pathTo := this.GetPrefixAndPath(to)

	result, err2 := this.GetFilesystem(prefixTo).WriteStream(pathTo, buffer, conf...)
	if err2 != nil {
		return false, err2
	}

	return result, nil
}

func (this *MountManager) Move(from string, to string, conf ...map[string]any) (bool, error) {
	prefixFrom, pathFrom := this.GetPrefixAndPath(from)
	prefixTo, pathTo := this.GetPrefixAndPath(to)

	if prefixFrom == prefixTo {
		filesystem := this.GetFilesystem(prefixFrom)

		renamed, err := filesystem.Rename(pathFrom, pathTo)
		if err != nil {
			return false, err
		}

		if len(conf) > 0 {
			if visibility, ok := conf[0]["visibility"]; ok && renamed {
				return filesystem.SetVisibility(pathTo, visibility.(string))
			}
		}

		return renamed, nil
	}

	copied, err := this.Copy(from, to, conf...)
	if copied {
		return this.Delete(from)
	}

	return false, err
}

func (this *MountManager) Has(path string) bool {
	prefix, newPath := this.GetPrefixAndPath(path)

	return this.GetFilesystem(prefix).Has(newPath)
}

func (this *MountManager) Read(path string) ([]byte, error) {
	prefix, newPath := this.GetPrefixAndPath(path)

	return this.GetFilesystem(prefix).Read(newPath)
}

func (this *MountManager) ReadStream(path string) (io.ReadCloser, error) {
	prefix, newPath := this.GetPrefixAndPath(path)

	return this.GetFilesystem(prefix).ReadStream(newPath)
}

func (this *MountManager) GetMetadata(path string) (map[string]any, error) {
	prefix, newPath := this.GetPrefixAndPath(path)

	return this.GetFilesystem(prefix).GetMetadata(newPath)
}

func (this *MountManager) GetSize(path string) (int64, error) {
	prefix, newPath := this.GetPrefixAndPath(path)

	return this.GetFilesystem(prefix).GetSize(newPath)
}

func (this *MountManager) GetMimetype(path string) (string, error) {
	prefix, newPath := this.GetPrefixAndPath(path)

	return this.GetFilesystem(prefix).GetMimetype(newPath)
}

func (this *MountManager) GetTimestamp(path string) (int64, error) {
	prefix, newPath := this.GetPrefixAndPath(path)

	return this.GetFilesystem(prefix).GetTimestamp(newPath)
}

func (this *MountManager) GetVisibility(path string) (string, error) {
	prefix, newPath := this.GetPrefixAndPath(path)

	return this.GetFilesystem(prefix).GetVisibility(newPath)
}

func (this *MountManager) Write(path string, contents []byte, conf ...map[string]any) (bool, error) {
	prefix, newPath := this.GetPrefixAndPath(path)

	return this.GetFilesystem(prefix).Write(newPath, contents, conf...)
}

func (this *MountManager) WriteStream(path string, resource io.Reader, conf ...map[string]any) (bool, error) {
	prefix, newPath := this.GetPrefixAndPath(path)

	return this.GetFilesystem(prefix).WriteStream(newPath, resource, conf...)
}

func (this *MountManager) Update(path string, contents []byte, conf ...map[string]any) (bool, error) {
	prefix, newPath := this.GetPrefixAndPath(path)

	return this.GetFilesystem(prefix).Update(newPath, contents, conf...)
}

func (this *MountManager) UpdateStream(path string, resource io.Reader, conf ...map[string]any) (bool, error) {
	prefix, newPath := this.GetPrefixAndPath(path)

	return this.GetFilesystem(prefix).UpdateStream(newPath, resource, conf...)
}

func (this *MountManager) Rename(path string, newpath string) (bool, error) {
	prefix, pather := this.GetPrefixAndPath(path)

	return this.GetFilesystem(prefix).Rename(pather, newpath)
}

func (this *MountManager) Delete(path string) (bool, error) {
	prefix, newPath := this.GetPrefixAndPath(path)

	return this.GetFilesystem(prefix).Delete(newPath)
}

func (this *MountManager) DeleteDir(dirname string) (bool, error) {
	prefix, newDirname := this.GetPrefixAndPath(dirname)

	return this.GetFilesystem(prefix).DeleteDir(newDirname)
}

func (this *MountManager) CreateDir(dirname string, conf ...map[string]any) (bool, error) {
	prefix, newDirname := this.GetPrefixAndPath(dirname)

	return this.GetFilesystem(prefix).CreateDir(newDirname, conf...)
}

func (this *MountManager) SetVisibility(path string, visibility string) (bool, error) {
	prefix, newPath := this.GetPrefixAndPath(path)

	return this.GetFilesystem(prefix).SetVisibility(newPath, visibility)
}

func (this *MountManager) Put(path string, contents []byte, conf ...map[string]any) (bool, error) {
	prefix, newPath := this.GetPrefixAndPath(path)

	return this.GetFilesystem(prefix).Put(newPath, contents, conf...)
}

func (this *MountManager) PutStream(path string, resource io.Reader, conf ...map[string]any) (bool, error) {
	prefix, newPath := this.GetPrefixAndPath(path)

	return this.GetFilesystem(prefix).PutStream(newPath, resource, conf...)
}

func (this *MountManager) ReadAndDelete(path string) (any, error) {
	prefix, newPath := this.GetPrefixAndPath(path)

	return this.GetFilesystem(prefix).ReadAndDelete(newPath)
}

// file := Get("/file.txt").(*File)
// dir := Get("/dir").(*Directory)
func (this *MountManager) Get(path string, handler ...func(*Filesystem, string) any) any {
	prefix, newPath := this.GetPrefixAndPath(path)

	return this.GetFilesystem(prefix).Get(newPath, handler...)
}

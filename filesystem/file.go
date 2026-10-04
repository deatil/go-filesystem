package filesystem

import (
	"io"
)

/**
 * File
 *
 * @create 2021-8-1
 * @author deatil
 */
type File struct {
	Handler
}

func NewFile(filesystem *Filesystem, path ...string) *File {
	fs := &File{}
	fs.filesystem = filesystem

	if len(path) > 0 {
		fs.path = path[0]
	}

	return fs
}

func (this *File) WithFilesystem(filesystem *Filesystem) *File {
	this.filesystem = filesystem

	return this
}

func (this *File) WithPath(path string) *File {
	this.path = path

	return this
}

func (this *File) Exists() bool {
	return this.filesystem.Has(this.path)
}

func (this *File) Read() ([]byte, error) {
	return this.filesystem.Read(this.path)
}

func (this *File) ReadStream() (io.Reader, error) {
	return this.filesystem.ReadStream(this.path)
}

func (this *File) Write(content []byte) (bool, error) {
	return this.filesystem.Write(this.path, content)
}

func (this *File) WriteStream(resource io.Reader) (bool, error) {
	return this.filesystem.WriteStream(this.path, resource)
}

func (this *File) Update(content []byte) (bool, error) {
	return this.filesystem.Update(this.path, content)
}

func (this *File) UpdateStream(resource io.Reader) (bool, error) {
	return this.filesystem.UpdateStream(this.path, resource)
}

func (this *File) Put(content []byte) (bool, error) {
	return this.filesystem.Update(this.path, content)
}

func (this *File) PutStream(resource io.Reader) (bool, error) {
	return this.filesystem.PutStream(this.path, resource)
}

func (this *File) Rename(newpath string) (bool, error) {
	if _, err := this.filesystem.Rename(this.path, newpath); err != nil {
		return false, err
	}

	this.path = newpath
	return true, nil
}

func (this *File) Copy(newpath string) (*File, error) {
	_, err := this.filesystem.Copy(this.path, newpath)
	if err != nil {
		return nil, err
	}

	var file = &File{}
	file.filesystem = this.filesystem
	file.path = newpath

	return file, nil
}

func (this *File) Delete() (bool, error) {
	return this.filesystem.Delete(this.path)
}

func (this *File) GetTimestamp() (int64, error) {
	return this.filesystem.GetTimestamp(this.path)
}

func (this *File) GetMimetype() (string, error) {
	return this.filesystem.GetMimetype(this.path)
}

func (this *File) GetVisibility() (string, error) {
	return this.filesystem.GetVisibility(this.path)
}

func (this *File) GetMetadata() (map[string]any, error) {
	return this.filesystem.GetMetadata(this.path)
}

func (this *File) GetSize() (int64, error) {
	return this.filesystem.GetSize(this.path)
}

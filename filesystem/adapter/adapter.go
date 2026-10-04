package adapter

import (
	"io"

	"github.com/deatil/go-filesystem/filesystem/interfaces"
)

/**
 * Empty adapter
 *
 * @create 2021-8-1
 * @author deatil
 */
type Adapter struct {
	Abstract
}

func (this *Adapter) Has(path string) bool {
	return false
}

func (this *Adapter) Write(path string, contents []byte, conf interfaces.Config) (map[string]any, error) {
	panic("go-filesystem: Write does not implement")
}

func (this *Adapter) WriteStream(path string, stream io.Reader, conf interfaces.Config) (map[string]any, error) {
	panic("go-filesystem: WriteStream does not implement")
}

func (this *Adapter) Update(path string, contents []byte, conf interfaces.Config) (map[string]any, error) {
	panic("go-filesystem: Update does not implement")
}

func (this *Adapter) UpdateStream(path string, stream io.Reader, conf interfaces.Config) (map[string]any, error) {
	panic("go-filesystem: UpdateStream does not implement")
}

func (this *Adapter) Read(path string) (map[string]any, error) {
	panic("go-filesystem: Read does not implement")
}

func (this *Adapter) ReadStream(path string) (map[string]any, error) {
	panic("go-filesystem: ReadStream does not implement")
}

func (this *Adapter) Rename(path string, newpath string) error {
	panic("go-filesystem: Rename does not implement")
}

func (this *Adapter) Copy(path string, newpath string) error {
	panic("go-filesystem: Copy does not implement")
}

func (this *Adapter) Delete(path string) error {
	panic("go-filesystem: Delete does not implement")
}

func (this *Adapter) Create(path string, conf interfaces.Config) (map[string]string, error) {
	panic("go-filesystem: Create does not implement")
}

func (this *Adapter) CopyDir(path string, newpath string) error {
	panic("go-filesystem: CopyDir does not implement")
}

func (this *Adapter) DeleteDir(dirname string) error {
	panic("go-filesystem: DeleteDir does not implement")
}

func (this *Adapter) CreateDir(dirname string, conf interfaces.Config) (map[string]string, error) {
	panic("go-filesystem: CreateDir does not implement")
}

func (this *Adapter) ListContents(directory string, recursive ...bool) ([]map[string]any, error) {
	panic("go-filesystem: ListContents does not implement")
}

func (this *Adapter) GetMetadata(path string) (map[string]any, error) {
	panic("go-filesystem: GetMetadata does not implement")
}

func (this *Adapter) GetSize(path string) (map[string]any, error) {
	panic("go-filesystem: GetSize does not implement")
}

func (this *Adapter) GetMimetype(path string) (map[string]any, error) {
	panic("go-filesystem: GetMimetype does not implement")
}

func (this *Adapter) GetTimestamp(path string) (map[string]any, error) {
	panic("go-filesystem: GetTimestamp does not implement")
}

func (this *Adapter) GetVisibility(path string) (map[string]string, error) {
	panic("go-filesystem: GetVisibility does not implement")
}

func (this *Adapter) SetVisibility(path string, visibility string) (map[string]string, error) {
	panic("go-filesystem: SetVisibility does not implement")
}

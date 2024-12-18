package adapter

import(
    "io"

    "github.com/deatil/go-filesystem/filesystem/interfaces"
)

/**
 * 空适配器
 *
 * @create 2021-8-1
 * @author deatil
 */
type Adapter struct {
    Abstract
}

// 判断
func (this *Adapter) Has(string) bool {
    return false
}

// 上传
func (this *Adapter) Write(path string, contents []byte, conf interfaces.Config) (map[string]any, error) {
    panic("go-filesystem: Write does not implement")
}

// 上传 Stream 文件类型
func (this *Adapter) WriteStream(path string, stream io.Reader, conf interfaces.Config) (map[string]any, error) {
    panic("go-filesystem: WriteStream does not implement")
}

// 更新
func (this *Adapter) Update(path string, contents []byte, conf interfaces.Config) (map[string]any, error) {
    panic("go-filesystem: Update does not implement")
}

// 更新
func (this *Adapter) UpdateStream(path string, stream io.Reader, conf interfaces.Config) (map[string]any, error) {
    panic("go-filesystem: UpdateStream does not implement")
}

// 读取
func (this *Adapter) Read(path string) (map[string]any, error) {
    panic("go-filesystem: Read does not implement")
}

// 读取数据为数据流
func (this *Adapter) ReadStream(path string) (map[string]any, error) {
    panic("go-filesystem: ReadStream does not implement")
}

// 重命名
func (this *Adapter) Rename(path string, newpath string) error {
    panic("go-filesystem: Rename does not implement")
}

// 复制
func (this *Adapter) Copy(path string, newpath string) error {
    panic("go-filesystem: Copy does not implement")
}

// 删除
func (this *Adapter) Delete(path string) error {
    panic("go-filesystem: Delete does not implement")
}

// 删除文件夹
func (this *Adapter) DeleteDir(dirname string) error {
    panic("go-filesystem: DeleteDir does not implement")
}

// 创建文件夹
func (this *Adapter) CreateDir(dirname string, conf interfaces.Config) (map[string]string, error) {
    panic("go-filesystem: CreateDir does not implement")
}

// 列出内容
func (this *Adapter) ListContents(directory string, recursive ...bool) ([]map[string]any, error) {
    panic("go-filesystem: ListContents does not implement")
}

//
func (this *Adapter) GetMetadata(path string) (map[string]any, error) {
    panic("go-filesystem: GetMetadata does not implement")
}

//
func (this *Adapter) GetSize(path string) (map[string]any, error) {
    panic("go-filesystem: GetSize does not implement")
}

//
func (this *Adapter) GetMimetype(path string) (map[string]any, error) {
    panic("go-filesystem: GetMimetype does not implement")
}

//
func (this *Adapter) GetTimestamp(path string) (map[string]any, error) {
    panic("go-filesystem: GetTimestamp does not implement")
}

// 获取文件的权限
func (this *Adapter) GetVisibility(path string) (map[string]string, error) {
    panic("go-filesystem: GetVisibility does not implement")
}

// 设置文件的权限
func (this *Adapter) SetVisibility(path string, visibility string) (map[string]string, error) {
    panic("go-filesystem: SetVisibility does not implement")
}

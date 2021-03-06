package filesystem

// new 文件管理器
func NewDirectory(filesystem *Fllesystem, path ...string) *Directory {
    fs := &Directory{}
    fs.filesystem = filesystem

    if len(path) > 0{
        fs.path = path[0]
    }

    return fs
}

/**
 * 文件管理器文件夹操作扩展
 *
 * @create 2021-8-1
 * @author deatil
 */
type Directory struct {
    Handler
}

// 设置管理器
func (this *Directory) SetFilesystem(filesystem *Fllesystem) *Directory {
    this.filesystem = filesystem

    return this
}

// 设置目录
func (this *Directory) SetPath(path string) *Directory {
    this.path = path

    return this
}

// 删除文件夹
func (this *Directory) Delete() (bool, error) {
    return this.filesystem.DeleteDir(this.path)
}

// 列出文件
func (this *Directory) GetContents(recursive ...bool) ([]map[string]any, error) {
    return this.filesystem.ListContents(this.path, recursive...)
}

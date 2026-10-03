package filesystem

/**
 * Directory
 *
 * @create 2021-8-1
 * @author deatil
 */
type Directory struct {
	Handler
}

func NewDirectory(filesystem *Filesystem, path ...string) *Directory {
	fs := &Directory{}
	fs.filesystem = filesystem

	if len(path) > 0 {
		fs.path = path[0]
	}

	return fs
}

func (this *Directory) WithFilesystem(filesystem *Filesystem) *Directory {
	this.filesystem = filesystem

	return this
}

func (this *Directory) WithPath(path string) *Directory {
	this.path = path

	return this
}

func (this *Directory) Delete() (bool, error) {
	return this.filesystem.DeleteDir(this.path)
}

func (this *Directory) GetContents(recursive ...bool) ([]map[string]any, error) {
	return this.filesystem.ListContents(this.path, recursive...)
}

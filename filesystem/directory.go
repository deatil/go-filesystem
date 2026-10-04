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

func (this *Directory) Exists() bool {
	return this.filesystem.Has(this.path)
}

func (this *Directory) Rename(newpath string) (bool, error) {
	if _, err := this.filesystem.Rename(this.path, newpath); err != nil {
		return false, err
	}

	this.path = newpath
	return true, nil
}

func (this *Directory) Copy(newpath string) (*Directory, error) {
	_, err := this.filesystem.CopyDir(this.path, newpath)
	if err != nil {
		return nil, err
	}

	var file = &Directory{}
	file.filesystem = this.filesystem
	file.path = newpath

	return file, nil
}

func (this *Directory) Delete() (bool, error) {
	return this.filesystem.DeleteDir(this.path)
}

func (this *Directory) Create(conf ...map[string]any) (bool, error) {
	return this.filesystem.CreateDir(this.path, conf...)
}

func (this *Directory) GetContents(recursive ...bool) ([]map[string]any, error) {
	return this.filesystem.ListContents(this.path, recursive...)
}

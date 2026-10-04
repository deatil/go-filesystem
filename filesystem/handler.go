package filesystem

/**
 * Handler
 *
 * @create 2021-8-1
 * @author deatil
 */
type Handler struct {
	filesystem *Filesystem
	path       string
}

func (this *Handler) IsDir() bool {
	return this.GetType() == "dir"
}

func (this *Handler) IsFile() bool {
	return this.GetType() == "file"
}

func (this *Handler) GetType() string {
	metadata, _ := this.filesystem.GetMetadata(this.path)
	if metadata == nil {
		return "dir"
	}

	if typ, ok := metadata["type"].(string); ok {
		return typ
	}

	return "dir"
}

func (this *Handler) SetFilesystem(filesystem *Filesystem) any {
	this.filesystem = filesystem

	return this
}

func (this *Handler) GetFilesystem() *Filesystem {
	return this.filesystem
}

func (this *Handler) SetPath(path string) any {
	this.path = path

	return this
}

func (this *Handler) GetPath() string {
	return this.path
}

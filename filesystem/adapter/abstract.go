package adapter

import (
	"strings"
)

/**
 * Abstract
 *
 * @create 2021-8-1
 * @author deatil
 */
type Abstract struct {
	pathPrefix    string
	pathSeparator string
}

func (this *Abstract) SetPathPrefix(prefix string) {
	if prefix == "" {
		this.pathPrefix = ""

		return
	}

	this.pathSeparator = "/"
	this.pathPrefix = strings.TrimSuffix(prefix, "/") + this.pathSeparator
}

func (this *Abstract) GetPathPrefix() string {
	return this.pathPrefix
}

func (this *Abstract) ApplyPathPrefix(path string) string {
	return this.GetPathPrefix() + strings.TrimPrefix(path, "/")
}

func (this *Abstract) RemovePathPrefix(path string) string {
	prefix := this.GetPathPrefix()
	return strings.TrimPrefix(path, prefix)
}

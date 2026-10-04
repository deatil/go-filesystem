package interfaces

import (
	"io"
)

/**
 * Adapter interface
 *
 * @create 2021-8-1
 * @author deatil
 */
type Adapter interface {
	// exists
	Has(path string) bool

	// Write
	Write(path string, contents []byte, conf Config) (map[string]any, error)

	// WriteStream
	WriteStream(path string, stream io.Reader, conf Config) (map[string]any, error)

	// Update
	Update(path string, contents []byte, conf Config) (map[string]any, error)

	// UpdateStream
	UpdateStream(path string, stream io.Reader, conf Config) (map[string]any, error)

	// Read
	Read(path string) (map[string]any, error)

	// ReadStream
	ReadStream(path string) (map[string]any, error)

	// Rename
	Rename(path string, newpath string) error

	// Copy
	Copy(path string, newpath string) error

	// Delete
	Delete(path string) error

	// Create file
	Create(path string, conf Config) (map[string]string, error)

	// Copy dir
	CopyDir(path string, newpath string) error

	// DeleteDir
	DeleteDir(path string) error

	// CreateDir
	CreateDir(path string, conf Config) (map[string]string, error)

	// ListContents
	ListContents(path string, recursive ...bool) ([]map[string]any, error)

	// GetMetadata
	GetMetadata(path string) (map[string]any, error)

	// GetSize
	GetSize(path string) (map[string]any, error)

	// GetMimetype
	GetMimetype(path string) (map[string]any, error)

	// GetTimestamp
	GetTimestamp(path string) (map[string]any, error)

	// GetVisibility
	GetVisibility(path string) (map[string]string, error)

	// SetVisibility
	SetVisibility(path string, visibility string) (map[string]string, error)
}

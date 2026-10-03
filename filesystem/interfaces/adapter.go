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
	// SetPathPrefix
	SetPathPrefix(string)

	// GetPathPrefix
	GetPathPrefix() string

	// ApplyPathPrefix
	ApplyPathPrefix(string) string

	// RemovePathPrefix
	RemovePathPrefix(string) string

	// exists
	Has(string) bool

	// Write
	Write(string, []byte, Config) (map[string]any, error)

	// WriteStream
	WriteStream(string, io.Reader, Config) (map[string]any, error)

	// Update
	Update(string, []byte, Config) (map[string]any, error)

	// UpdateStream
	UpdateStream(string, io.Reader, Config) (map[string]any, error)

	// Read
	Read(string) (map[string]any, error)

	// ReadStream
	ReadStream(string) (map[string]any, error)

	// Rename
	Rename(string, string) error

	// Copy
	Copy(string, string) error

	// Delete
	Delete(string) error

	// DeleteDir
	DeleteDir(string) error

	// CreateDir
	CreateDir(string, Config) (map[string]string, error)

	// ListContents
	ListContents(string, ...bool) ([]map[string]any, error)

	// GetMetadata
	GetMetadata(string) (map[string]any, error)

	// GetSize
	GetSize(string) (map[string]any, error)

	// GetMimetype
	GetMimetype(string) (map[string]any, error)

	// GetTimestamp
	GetTimestamp(string) (map[string]any, error)

	// GetVisibility
	GetVisibility(string) (map[string]string, error)

	// SetVisibility
	SetVisibility(string, string) (map[string]string, error)
}

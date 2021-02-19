package gadget

import (
	"embed"
	"errors"
	"io"
	"io/fs"
	"net/http"
	"path"
	"path/filepath"
	"strings"
)

type ioFile struct {
	io.Seeker
	fs.File
}

func (*ioFile) Readdir(count int) ([]fs.FileInfo, error) {
	return nil, nil
}

func NewEmbed(fs embed.FS, path string) *embedFS {
	return &embedFS{
		embed: fs,
		path:  path,
	}
}

type embedFS struct {
	embed embed.FS
	path  string
}

func (f *embedFS) Open(name string) (http.File, error) {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) {
		return nil, errors.New("http: invalid character in file path")
	}
	fullName := filepath.Join(f.path, filepath.FromSlash(path.Clean(name)))
	file, err := f.embed.Open(fullName)
	return &ioFile{
		File: file,
	}, err
}

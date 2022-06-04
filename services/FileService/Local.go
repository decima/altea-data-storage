package FileService

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var _ FileServiceInterface = (*Local)(nil)

type Local struct {
	basePath string
}

func (l *Local) getFullPath(path string) string {

	clean := filepath.Clean(l.basePath + "/" + path)

	return clean
}

func (l *Local) convertToOSPath(path string) string {
	clean := filepath.Clean(l.basePath + "/" + path)

	return filepath.FromSlash(clean)
}

func (l *Local) convertToWebPath(path string) string {

	clean := filepath.Clean(l.getRelativePath(path))

	return filepath.ToSlash(clean)
}

func (l *Local) getRelativePath(fullPath string) string {

	return l.getRelativePathFromBase(fullPath, l.basePath)
}

func (l *Local) getRelativePathFromBase(fullPath string, base string) string {

	return strings.TrimPrefix(strings.Replace(fullPath, base, "", 1), "/")

}

func NewLocal(basePath string) *Local {
	return &Local{basePath: basePath}
}

func (l Local) Delete(path string) error {
	return os.RemoveAll(l.convertToOSPath(path))
}

func (l Local) Write(path string, content []byte) error {
	return ioutil.WriteFile(l.convertToOSPath(path), content, 0700)
}

func (l Local) Mkdir(path string, recursive bool) error {
	if !recursive {
		return os.Mkdir(l.convertToOSPath(path), 0700)
	}
	return os.MkdirAll(l.convertToOSPath(path), 0700)
}

func (l Local) Read(path string) ([]byte, error) {
	return ioutil.ReadFile(l.convertToOSPath(path))
}

func (l Local) Exists(path string) bool {
	_, err := os.Open(l.convertToOSPath(path))
	return err == nil
}

func (l Local) ReadDir(path string) (*DirectoryContent, error) {
	f := NewDirectoryContent(path)
	files, err := os.ReadDir(l.convertToOSPath(path))
	if err != nil {
		return nil, err
	}
	for _, element := range files {
		if strings.HasPrefix(element.Name(), ".") {
			continue
		}
		fileType := File
		if element.IsDir() {
			fileType = Directory
		}
		f.Items[element.Name()] = FileInfo{
			Name: element.Name(),
			Path: l.convertToWebPath(path + "/" + element.Name()),
			Type: fileType,
		}
	}
	return f, nil
}

func (l Local) IsDir(path string) bool {
	file, err := os.Open(l.convertToOSPath(path))

	if err != nil {
		return false
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return false
	}
	return fileInfo.IsDir()

}

package managers

import (
	"Altea/services/FileService"
	"io"
	"io/ioutil"
	"path/filepath"
)

type FileManager struct {
	fileService *FileService.FileServiceInterface
}

func NewFileManager(fileService *FileService.FileServiceInterface) *FileManager {
	return &FileManager{fileService: fileService}
}

func (c *FileManager) GetPath(path string) (interface{}, FileService.ItemType, error) {
	if !(*c.fileService).Exists(path) {
		return nil, FileService.NotFound, nil
	}

	if (*c.fileService).IsDir(path) {
		directoryContent, err := (*c.fileService).ReadDir(path)
		return directoryContent, FileService.Directory, err
	}

	fileContent, err := (*c.fileService).Read(path)
	return fileContent, FileService.File, err

}

func (c *FileManager) WritePath(path string, content io.Reader) error {
	byteContent, err := ioutil.ReadAll(content)
	if err != nil {
		return err
	}
	err = (*c.fileService).Mkdir(filepath.Dir(path), true)
	if err != nil {
		return err
	}
	return (*c.fileService).Write(path, byteContent)
}

func (c FileManager) DeletePath(path string) error {
	return (*c.fileService).Delete(path)

}

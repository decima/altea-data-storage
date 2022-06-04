package FileService

import (
	"errors"
	"path/filepath"
	"strings"
	"sync"
)

var _ FileServiceInterface = (*InMemory)(nil)

type InMemory struct {
	directoryListing map[string]DirectoryContent
	fileContent      map[string][]byte
	mux              sync.Mutex
}

func (i InMemory) Delete(path string) error {
	for k := range i.fileContent {
		if strings.HasPrefix(k, path) {
			delete(i.fileContent, k)
		}
	}
	if i.IsDir(path) {
		for k := range i.directoryListing {
			if strings.HasPrefix(k, path) {
				delete(i.directoryListing, k)
			}
		}
	}
	name := filepath.Base(path)

	delete(i.directoryListing[filepath.Dir(path)].Items, name)

	return nil
}

func NewInMemory() *InMemory {
	return &InMemory{
		directoryListing: map[string]DirectoryContent{"/": *NewDirectoryContent("/")},
		fileContent:      map[string][]byte{},
		mux:              sync.Mutex{},
	}
}

func (i InMemory) Write(path string, content []byte) error {
	i.mux.Lock()
	i.fileContent[path] = content
	i.mux.Unlock()
	directory := filepath.Dir(path)
	if _, ok := i.directoryListing[directory]; !ok {
		i.Mkdir(directory, true)
	}
	i.mux.Lock()
	defer i.mux.Unlock()
	name := filepath.Base(path)
	i.directoryListing[directory].AddFile(name, FileInfo{Name: name, Type: File, Path: path})

	return nil
}

func (i InMemory) Mkdir(path string, recursive bool) error {
	i.mux.Lock()

	if _, ok := i.directoryListing[path]; ok {
		i.mux.Unlock()

		return nil
	}
	i.directoryListing[path] = *NewDirectoryContent(path)
	i.mux.Unlock()
	if path == "/" {

		return nil
	}
	if recursive {
		directory := filepath.Dir(path)
		err := i.Mkdir(directory, recursive)
		if err != nil {
			return err
		}
		name := filepath.Base(path)
		i.mux.Lock()
		i.directoryListing[filepath.Dir(path)].AddFile(name, FileInfo{Name: name, Type: Directory, Path: path})
		i.mux.Unlock()

	}
	return nil
}

func (i InMemory) Read(path string) ([]byte, error) {
	defer i.mux.Unlock()
	i.mux.Lock()
	if content, ok := i.fileContent[path]; ok {
		return content, nil
	}
	return nil, errors.New("not found")
}

func (i InMemory) Exists(path string) bool {
	defer i.mux.Unlock()
	i.mux.Lock()
	if _, ok := i.fileContent[path]; ok {
		return true
	}
	if _, ok := i.directoryListing[path]; ok {
		return true
	}
	return false
}

func (i InMemory) ReadDir(path string) (*DirectoryContent, error) {

	defer i.mux.Unlock()
	i.mux.Lock()
	if content, ok := i.directoryListing[path]; ok {
		return &content, nil
	}
	return nil, errors.New("not found")
}

func (i InMemory) IsDir(path string) bool {
	defer i.mux.Unlock()
	i.mux.Lock()
	_, ok := i.directoryListing[path]
	return ok

}

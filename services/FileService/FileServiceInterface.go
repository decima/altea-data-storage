package FileService

const (
	Directory ItemType = "dir"
	File      ItemType = "file"
	NotFound  ItemType = "notFound"
)

type ItemType string

func (t ItemType) IsDirectory() bool {
	return t == Directory
}
func (t ItemType) IsFile() bool {
	return t == File
}
func (t ItemType) IsNotFound() bool {
	return t == NotFound
}

type FileServiceInterface interface {
	Write(path string, content []byte) error
	Mkdir(path string, recursive bool) error
	Read(path string) ([]byte, error)
	Exists(path string) bool
	ReadDir(path string) (*DirectoryContent, error)
	IsDir(path string) bool
	Delete(path string) error
}

type FileInfo struct {
	Name string   `json:"name"`
	Path string   `json:"path"`
	Type ItemType `json:"type"`
}

type DirectoryContent struct {
	Path  string              `json:"path"`
	Items map[string]FileInfo `json:"items"`
}

func NewDirectoryContent(path string) *DirectoryContent {
	return &DirectoryContent{
		Path:  path,
		Items: map[string]FileInfo{},
	}
}

func (directory DirectoryContent) AddFile(name string, file FileInfo) {
	directory.Items[name] = file

}

package storage

type Storage interface {
	GetFile(id string) (string, error)
}

type FileStorage struct{}

func (fs *FileStorage) GetFile(id string) (string, error) {
	return "Content from storage", nil
}

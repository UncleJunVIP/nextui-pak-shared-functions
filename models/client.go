package models

type Client interface {
	Close() error
	ListDirectory(subdirectory string) ([]Item, error)
	DownloadFile(remotePath, localPath, filename string) (string, error)
	DownloadFileRename(remotePath, localPath, filename, rename string) (string, error)
}

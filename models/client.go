package models

type Client interface {
	Close() error
	ListDirectory(subdirectory string) ([]Items, error)
	DownloadFile(remotePath, localPath, filename string) (string, error)
	DownloadFileRename(remotePath, localPath, filename, rename string) (string, error)
}

package domain

type AppConfig struct {
	FolderPath string
	Port       int
	AppPath    string
}

type DirEntry struct {
	Path          string
	IsDir         bool
	Name          string
	SizeInBytes   string
	Created       string
	Updated       string
	FileExtension string
}

type FileInfo struct {
	DirEntry
	IpfsHash string
}

type OpenFileInfo struct {
	Location string
}

type KeyPair struct {
	PublicKey  string
	PrivateKey string
}
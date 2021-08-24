package middlewares

type FileListItem struct {
	IsDirectory bool   `json:"isDirectory"`
	FileSize    uint   `json:"fileSize"`
	FileName    string `json:"fileName"`
	ModifiedAt  uint   `json:"modifiedAt"`
}

type FileList []FileListItem

type FileListResponse struct {
	Success  bool     `json:"success"`
	Message  string   `json:"message"`
	FileList FileList `json:"fileList"`
}

package models

type UploadQueryParams struct {
	Name string `json:"name,required"`
}

type DownloadArgs struct {
	Path string `json:"path"`
}

type DownloadResponse struct {
	ContentType string `json:"content_type"`
	File        []byte `json:"file"`
}

package models

type UploadQueryParams struct {
	Name string `json:"name,required"`
}

type DownloadArgs struct {
	Path string `json:"path"`
}

package send_to_js

import (
	"github.com/ganeshrvel/go-mtpfs/mtp"
	"github.com/ganeshrvel/go-mtpx"
)

type ErrorType string

type ErrorResult struct {
	ErrorType ErrorType   `json:"errorType"`
	Error     string      `json:"error"`
	Data      interface{} `json:"data"`
}

type InitializeResult struct {
	ErrorType ErrorType      `json:"errorType"`
	Error     string         `json:"error"`
	Data      mtp.DeviceInfo `json:"data"`
}

type DeviceInfoResult struct {
	ErrorType ErrorType      `json:"errorType"`
	Error     string         `json:"error"`
	Data      mtp.DeviceInfo `json:"data"`
}

type StoragesResult struct {
	ErrorType ErrorType          `json:"errorType"`
	Error     string             `json:"error"`
	Data      []mtpx.StorageData `json:"data"`
}

type MakeDirectoryResult struct {
	ErrorType ErrorType `json:"errorType"`
	Error     string    `json:"error"`
	Data      bool      `json:"data"`
}

type DisposeResult struct {
	ErrorType ErrorType `json:"errorType"`
	Error     string    `json:"error"`
	Data      bool      `json:"data"`
}

type FileInfo struct {
	Size       int64  `json:"size"`
	IsDir      bool   `json:"isDir"`
	ModTime    string `json:"modTime"`
	Name       string `json:"name"`
	FullPath   string `json:"fullPath"`
	ParentPath string `json:"parentPath"`
	Extension  string `json:"extension"`
	ParentId   uint32 `json:"parentId"`
	ObjectId   uint32 `json:"objectId"`
}

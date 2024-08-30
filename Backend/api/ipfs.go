package api

import "io"

type IPFSRepository interface {
	UploadData(data io.Reader) (string, error)
	GetData(hash string) ([]byte, error)
}

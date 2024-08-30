package repository

import (
	"errors"
	"io"
	"net/http"
)

type IPFSRepository struct {
	apiURL string
}

func NewIPFSRepository(apiURL string) *IPFSRepository {
	return &IPFSRepository{apiURL: apiURL}
}

func (c *IPFSRepository) UploadData(data io.Reader) (string, error) {
	resp, err := http.Post(c.apiURL+"/api/v0/add", "multipart/form-data", data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to upload data to IPFS")
	}

	return string(body), nil
}

func (c *IPFSRepository) GetData(hash string) ([]byte, error) {
	resp, err := http.Get(c.apiURL + "/api/v0/cat?arg=" + hash)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to retrieve data from IPFS")
	}

	return body, nil
}

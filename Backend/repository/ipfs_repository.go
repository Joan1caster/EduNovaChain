package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	configs "nftPlantform/config"
	"nftPlantform/models"
	"strings"
)

type IPFSRepository struct {
	apiURL string
}

func NewIPFSRepository(apiURL string) *IPFSRepository {
	return &IPFSRepository{apiURL: apiURL}
}

func (c *IPFSRepository) UploadData(createData interface{}) (*models.IpfsResponseData, error) {
	url := "https://api.pinata.cloud/pinning/pinJSONToIPFS"
	jsonString, _ := json.Marshal(createData)
	payload := strings.NewReader(string(jsonString))
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySW5mb3JtYXRpb24iOnsiaWQiOiJkNmM1Y2M0MS1hZmIxLTQwOGYtOGUzNi0yZTRjZWRlNTM0YTEiLCJlbWFpbCI6Imxvbmd5dTAxMTZAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsInBpbl9wb2xpY3kiOnsicmVnaW9ucyI6W3siZGVzaXJlZFJlcGxpY2F0aW9uQ291bnQiOjEsImlkIjoiRlJBMSJ9LHsiZGVzaXJlZFJlcGxpY2F0aW9uQ291bnQiOjEsImlkIjoiTllDMSJ9XSwidmVyc2lvbiI6MX0sIm1mYV9lbmFibGVkIjpmYWxzZSwic3RhdHVzIjoiQUNUSVZFIn0sImF1dGhlbnRpY2F0aW9uVHlwZSI6InNjb3BlZEtleSIsInNjb3BlZEtleUtleSI6IjZkMDQzNzc3MGQ1ZjlmMzI2ZWRkIiwic2NvcGVkS2V5U2VjcmV0IjoiNDg3YmI4ZWJmMDE0ZmQ2YjNlYmU1ZWFkODUwMDk4OWU4MTQ2ZjM3NjRkOGE0MjhiZjg2Y2IwN2Q3NzFmYmU3NyIsImV4cCI6MTc1NzY2MTY2N30.mk3j5EkArR5ar-_ExT6kjaP_tja6YzdZfshnH8uwgLo")
	fmt.Println(configs.AppConfig.IpfsApiKey)
	req.Header.Add("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()
	fmt.Println(res)
	fmt.Println(string(body))
	var ipfsResponse models.IpfsResponseData
	err := json.Unmarshal(body, &ipfsResponse)
	return &ipfsResponse, err
}

func (c *IPFSRepository) GetData(ipfsHash string) (*models.Metadata, error) {
	url := fmt.Sprintf("https://gateway.pinata.cloud/ipfs/%s", ipfsHash)
	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var ipfsData models.Metadata
	err = json.Unmarshal(body, &ipfsData)
	if err != nil {
		return nil, err
	}
	return &ipfsData, nil
}

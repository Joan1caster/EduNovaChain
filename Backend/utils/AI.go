package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RequestBody struct {
	Words []string `json:"words"`
}

// http调python服务获取文章特征值
func GetFeatures(words []string) ([][512]float32, error) {
	reqBody := RequestBody{Words: words}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %v", err)
	}

	resp, err := http.Post("http://localhost:5000/get_features", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// 解析响应数据
	var features [][512]float32
	err = json.Unmarshal(body, &features)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %v", err)
	}

	return features, nil
}

// 计算特征相似度
func CalculateSimilarity(a, b [512]float32) float32 {
	var result float32
	for i := 0; i < 512; i++ {
		result += a[i] * b[i]
	}
	return result
}

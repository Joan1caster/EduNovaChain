package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

type RequestBody struct {
	Words []string `json:"words"`
}

// http调python服务获取文章特征值
func GetFeatures(words []string) (*[]float32, error) {
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
	var embeddings *[]float32
	// 解析响应数据
	err = json.Unmarshal(body, &embeddings)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %v", err)
	}
	fmt.Println("get data succeed.")
	return embeddings, nil
}

// 计算特征相似度
func CalculateSimilarity(a, b *[]float32) float32 {
	var result float32
	for i := 0; i < 512; i++ {
		result += (*a)[i] * (*b)[i]
	}
	return result
}

type VectorWithSimilarity struct {
	Vector     []float32
	Similarity float32
}

// SortVectorsBySimilarity sorts the vectors by their similarity to the given vector
func SortVectorsBySimilarity(vectors *[][]float32, targetVector *[]float32, limit int) [][]float32 {
	// Create a slice of VectorWithSimilarity
	vectorsWithSimilarity := make([]VectorWithSimilarity, len(*vectors))
	for i, v := range *vectors {
		similarity := CalculateSimilarity(&v, targetVector)
		vectorsWithSimilarity[i] = VectorWithSimilarity{v, similarity}
	}

	// Sort the slice based on similarity (in descending order)
	sort.Slice(vectorsWithSimilarity, func(i, j int) bool {
		return vectorsWithSimilarity[i].Similarity > vectorsWithSimilarity[j].Similarity
	})

	// Take the top 'limit' results
	result := make([][]float32, 0, limit)
	for i := 0; i < limit && i < len(vectorsWithSimilarity); i++ {
		result = append(result, vectorsWithSimilarity[i].Vector)
	}

	return result
}

package utils

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func VerifySignature(address, signature, message string) bool {
	addr := common.HexToAddress(address)
	sig := common.FromHex(signature)
	msgHash := crypto.Keccak256([]byte(message))
	pubKey, err := crypto.SigToPub(msgHash, sig)
	if err != nil {
		return false
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return addr == recoveredAddr
}

type ContractABI struct {
	ABI json.RawMessage `json:"abi"`
}

func ReadABI(filePath string) (json.RawMessage, string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, "", fmt.Errorf("读取文件失败: %v", err)
	}
	var contractABI ContractABI
	err = json.Unmarshal(content, &contractABI)
	if err != nil {
		return nil, "", fmt.Errorf("解析JSON失败: %v", err)
	}
	jsonABI := contractABI.ABI
	stringABI := string(jsonABI)

	return jsonABI, stringABI, nil
}

// GenerateNonce 生成一个随机的 nonce
func GenerateNonce() string {
	nonce := make([]byte, 16)
	_, err := rand.Read(nonce)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(nonce)
}

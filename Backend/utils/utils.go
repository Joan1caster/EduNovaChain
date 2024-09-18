package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"

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
	file, err := os.Create("abi/abi.json")
	if err != nil {
		return nil, "", fmt.Errorf("创建文件失败: %v", err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // 设置缩进格式化输出
	err = encoder.Encode(jsonABI)
	if err != nil {
		return nil, "", fmt.Errorf("创建文件失败: %v", err)
	}
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

// Float32ArrayToBlob converts a [512]float32 to []byte
func Float32ArrayToBlob(arr [512]float32) ([]byte, error) {
	buf := new(bytes.Buffer)
	for _, v := range arr {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

// BlobToFloat32Array converts []byte to [512]float32
func BlobToFloat32Array(b []byte) ([]float32, error) {
	arr := make([]float32, 512)
	buf := bytes.NewReader(b)
	for i := 0; i < 512; i++ {
		var v float32
		err := binary.Read(buf, binary.LittleEndian, &v)
		if err != nil {
			return arr, err
		}
		arr[i] = v
	}
	return arr, nil
}

func SetupLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:          true,                  // 显示完整的时间戳
		ForceColors:            true,                  // 强制启用颜色输出
		DisableLevelTruncation: true,                  // 禁用日志级别的截断
		TimestampFormat:        "2006-01-02 15:04:05", // 设置时间戳格式为"年-月-日 时:分:秒"
	})

	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		// 只输出到文件
		// logrus.SetOutput(file)
		// 如果要同时输出到控制台和文件，请启用下面一行代码
		logrus.SetOutput(io.MultiWriter(file, os.Stdout))
	} else {
		logrus.SetOutput(os.Stdout)
	}

	logrus.SetLevel(logrus.InfoLevel)
}

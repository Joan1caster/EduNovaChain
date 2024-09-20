package utils

import (
	"io"
	"log"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

var Sh = shell.NewShell("localhost:5001")

func DownloadString(cid string) ([]byte, error) {
	// 从IPFS获取文件内容
	data, err := Sh.Cat(cid)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// 读取内容
	content, err := io.ReadAll(data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return content, nil
}

func UploadString(content string) (string, error) {
	cid, err := Sh.Add(strings.NewReader(content))
	if err != nil {
		return "", err
	}
	return cid, nil
}

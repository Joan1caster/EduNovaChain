package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"nftPlantform/config"
	routes "nftPlantform/router"
	"os"

	"nftPlantform/internal/database"
	_ "nftPlantform/service"
)

func setupLogger() {
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
func main() {
	setupLogger()
	err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Fatal("Failed to load configuration")
	}
	log.Println("connect rpc-url:", config.AppConfig.Contract.Eth_rpc_url)
	database.ConnectDB()
	defer database.CloseDB()
	database.InitTable()
	db := database.GetDB()
	router := routes.SetupRouter(db)
	err = router.Run(":8080")
	if err != nil {
		return
	}
}

package routes

import (
	"nftPlantform/config"
	"nftPlantform/handlers"
	"nftPlantform/middleware"
	"nftPlantform/repository"
	"nftPlantform/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	userRepo := repository.NewGormUserRepository(db)
	nftRepo := repository.NewGormNFTRepository(db)
	orderRepo := repository.NewGormOrderRepository(db)
	ipfsRepo := repository.NewIPFSRepository(config.AppConfig.IpfsApiKey)
	orderService := service.NewOrderService(nftRepo, orderRepo)
	nftService := service.NewNFTService(nftRepo)
	blockChainService := service.NewBlockchainservice()
	tradeService := service.NewNFTTrade(orderService, nftService, blockChainService)
	userHandler := handlers.NewUserHandler(service.NewUserService(userRepo))
	ipfsHandler := handlers.NewIPFSHandler(service.NewIPFSService(ipfsRepo))
	nftHandler := handlers.NewNFTHandler(nftService, service.NewIPFSService(ipfsRepo))
	orderHandler := handlers.NewOrderHandler(orderService, nftService, tradeService)
	router := gin.Default()

	nftService.UpdateNFTCategory(3600) // 启动一个异步函数来处理NFT的分类,每小时更新一次

	// 公开路由
	public := router.Group("/api/v1")
	{
		public.GET("/siweMessage", userHandler.GetSIWEMessage) // 签名
		public.POST("/login", userHandler.Login)               // 验证签名并登录

		public.GET("/nfts/id", nftHandler.GetNFTByID)            // 根据NFT id查NFT信息
		public.GET("/nfts/creator", nftHandler.GetNFTsByCreator) // 根据NFT作者查所有NFT列表
		// public.GET("/nfts/retrieval", nftHandler.GetNFTBySummary)   // 根据文字内容查相关NFT列表
		public.GET("/nfts/latest/:number", nftHandler.GetLatestNFT)         // 返回最新的number个NFT
		public.GET("/nfts/topicAndType", nftHandler.GetNFTByTopicAndType) // 根据主题和类型查询NFT
		public.GET("/grade", nftHandler.GetGradeList) // 查询年级
		public.GET("/subject/:grade", nftHandler.GetSubjectByGrade)    // 根据年级查学科
		public.GET("/topic/subjectandgrade", nftHandler.GetTopicBySubjectAndGrade)    // 根据年级\学科查主题

		public.GET("/order/history", orderHandler.GetHistoryByNFTId) // 根据NFT id查其交易记录
	}

	// 需要认证的路由
	authenticated := router.Group("/api/v1")
	authenticated.Use(middleware.AuthMiddleware())
	{
		//user's router
		authenticated.GET("/user/favorite", userHandler.GetFavoriteTopic) // Query users’ favorite topics
		//ipfs 相关路由
		authenticated.POST("/ipfs/upload", ipfsHandler.UploadData) // 上传数据到IPFS
		authenticated.GET("/ipfs/data/:hash", ipfsHandler.GetData) // 从下载数据到本地
		// NFT 相关路由
		authenticated.POST("/nfts", nftHandler.CreateNFT) // 创建NFT

		// 订单相关路由
		authenticated.POST("/orders", orderHandler.ListNFT)         // 上架NFT
		authenticated.PUT("/orders/delist", orderHandler.DelistNFT) // 下架NFT
		// authenticated.GET("/orders/:id/buy", orderHandler.BuyNFT) // 购买NFT
	}

	return router
}

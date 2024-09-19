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
	tracRepo := repository.NewGormTransactionRepository(db)
	userService := service.NewUserService(userRepo)
	orderService := service.NewOrderService(nftRepo, orderRepo)
	nftService := service.NewNFTService(nftRepo)
	ipfsService := service.NewIPFSService(ipfsRepo)
	blockChainService := service.NewBlockchainservice()
	tradeService := service.NewNFTTrade(userRepo, tracRepo, orderRepo, orderService, nftService, blockChainService)
	userHandler := handlers.NewUserHandler(userService, nftService)
	ipfsHandler := handlers.NewIPFSHandler(ipfsService)
	nftHandler := handlers.NewNFTHandler(nftService, ipfsService, userService)
	orderHandler := handlers.NewOrderHandler(orderService, nftService, tradeService)
	router := gin.Default()

	nftService.UpdateNFTCategory(86400) // 启动一个异步函数来处理NFT的分类,每24小时更新一次

	// 公开路由
	public := router.Group("/api/v1")
	{
		public.GET("/siweMessage", userHandler.GetSIWEMessage) // 签名 -- pass
		public.POST("/login", userHandler.Login)               // 验证签名并登录 -- pass

		public.GET("/nfts/:id", nftHandler.GetNFTByID)             // 根据NFT id查NFT信息  -- pass
		public.GET("/nfts/creatorID", nftHandler.GetNFTsByCreator) // 根据NFT作者查所有NFT列表 -- pass
		// public.GET("/nfts/retrieval", nftHandler.GetNFTBySummary)   // 根据文字内容查相关NFT列表
		public.GET("/nfts/latest/:number", nftHandler.GetLatestNFT)                 // 返回最新的number个NFT -- pass
		public.GET("/nfts/hottest/:number", nftHandler.GetHottestNFT)               // 返回最新的number个NFT -- pass
		public.GET("/nfts/HighTrading/:number", nftHandler.GetHighTradingNFT)       // 返回最新的number个NFT -- pass
		public.GET("/nfts/topicAndType", nftHandler.GetNFTByTopicAndType)           // 根据主题和类型查询NFT -- pass
		public.POST("/nfts/feature", nftHandler.GetFeatures)                         // 根据输入查询特征值 -- pass
		public.POST("/nfts/details", nftHandler.GetNFTByDetails)                    // 根据关键词等一系列信息查询 -- pass
		public.GET("/grade", nftHandler.GetGradeList)                               // 查询年级 -- pass
		public.GET("/subject/:grade", nftHandler.GetSubjectByGrade)                 // 根据年级查学科 -- pass
		public.POST("/topic/subjectAndGrade", nftHandler.GetTopicBySubjectAndGrade) // 根据年级\学科查主题 -- pass
		public.GET("/order/history", orderHandler.GetHistoryByNFTId)                // 根据NFT id查其交易记录
	}

	// 需要认证的路由
	authenticated := router.Group("/api/v1")
	authenticated.Use(middleware.AuthMiddleware())
	{
		//user's router
		authenticated.GET("/check-auth/:nftID", userHandler.CheckAuth)    // 检查用户认证状态  -- pass
		authenticated.GET("/user/favorite", userHandler.GetFavoriteTopic) // Query users’ favorite topics
		//ipfs 相关路由
		authenticated.POST("/ipfs/upload", ipfsHandler.UploadData) // 上传数据到IPFS -- bad
		authenticated.GET("/ipfs/data/:hash", ipfsHandler.GetData) // 从下载数据到本地 -- bad
		// NFT 相关路由
		authenticated.POST("/nfts/create", nftHandler.CreateNFT) // 创建NFT -- problem: no subject
		authenticated.POST("/nfts/like/:nftID", nftHandler.LikeNFT)

		// 订单相关路由
		authenticated.POST("/orders", orderHandler.ListNFT)                         // 上架NFT -- pass
		authenticated.GET("/orders/status/:txHash", orderHandler.TransactionStatus) // 监听
		authenticated.PUT("/orders/delist", orderHandler.DelistNFT)                 // 下架NFT -- pass
		authenticated.POST("/orders/buy", orderHandler.BuyNFT)                      // 购买NFT -- pass

	}

	return router
}

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
	userHandler := handlers.NewUserHandler(service.NewUserService(userRepo))
	ipfsHandler := handlers.NewIPFSHandler(service.NewIPFSService(ipfsRepo))
	nftHandler := handlers.NewNFTHandler(service.NewNFTService(nftRepo), service.NewIPFSService(ipfsRepo))
	orderHandler := handlers.NewOrderHandler(service.NewOrderService(nftRepo, orderRepo), service.NewNFTService(nftRepo))
	router := gin.Default()

	// 公开路由
	public := router.Group("/api/v1")
	{
		public.GET("/siweMessage", userHandler.GetSIWEMessage)
		public.POST("/login", userHandler.Login)
		//public.POST("/auth", h.AuthenticateUser)
		//public.GET("/nfts", h.GetNFTs)
		//public.GET("/nfts/:id", h.GetNFTByID)
		//public.GET("/orders", h.GetOpenOrders)
	}

	// 需要认证的路由
	authenticated := router.Group("/api/v1")
	authenticated.Use(middleware.AuthMiddleware())
	{
		authenticated.POST("/ipfs/upload", ipfsHandler.UploadData)
		authenticated.GET("/ipfs/data/:hash", ipfsHandler.GetData)
		// NFT 相关路由
		authenticated.POST("/nfts", nftHandler.CreateNFT)
		//authenticated.DELETE("/nfts/:id", h.DeleteNFT)

		// 订单相关路由
		authenticated.POST("/orders", orderHandler.ListNFT)
		authenticated.PUT("/orders/:id", orderHandler.DelistNFT)
		// authenticated.POST("/orders/:id/buy", orderHandler.BuyNFT)

	}

	return router
}

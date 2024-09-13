package service

import (
	"fmt"
	"log"
	"sync"
	"time"

	"nftPlantform/common"
	"nftPlantform/repository"
)

type TransactionListener struct {
	NFTID        uint
	OrderID      uint
	BuyerID      uint
	TxHash       string
	StopChan     chan struct{}
	CompleteChan chan bool
}

type NFTTrade struct {
	userRepo          *repository.GormUserRepository
	tracRepo          *repository.GormTransactionRepository
	orderRepo         *repository.GormOrderRepository
	orderService      *OrderService
	nftService        *NFTService
	blockchainService *Blockchainservice
	listeners         map[string]*TransactionListener
	mu                sync.Mutex
}

func NewNFTTrade(
	userRepo *repository.GormUserRepository,
	tracRepo *repository.GormTransactionRepository,
	orderRepo *repository.GormOrderRepository,
	orderService *OrderService,
	nftService *NFTService,
	blockchainService *Blockchainservice,
) *NFTTrade {
	return &NFTTrade{
		userRepo:          userRepo,
		tracRepo:          tracRepo,
		orderRepo:         orderRepo,
		orderService:      orderService,
		nftService:        nftService,
		blockchainService: blockchainService,
		listeners:         make(map[string]*TransactionListener),
	}
}

func (s *NFTTrade) CreateTransaction(orderID, nftID, buyerID uint, txHash string, price string) error {
	_, err := s.tracRepo.CreateTransaction(orderID, txHash, price, "0", "PENDING")
	if err != nil {
		return fmt.Errorf("创建交易记录失败: %w", err)
	}
	return nil
}

func (s *NFTTrade) StartTransactionListener(nftID, orderID, buyerID uint, txHash string) {
	listener := &TransactionListener{
		NFTID:        nftID,
		OrderID:      orderID,
		BuyerID:      buyerID,
		TxHash:       txHash,
		StopChan:     make(chan struct{}),
		CompleteChan: make(chan bool),
	}

	s.mu.Lock()
	s.listeners[txHash] = listener
	s.mu.Unlock()

	go s.monitorTransaction(listener)
}

func (s *NFTTrade) monitorTransaction(listener *TransactionListener) {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		varifiedInfo, err := s.blockchainService.MonitorTransaction(listener)
		if err != nil {
			// 处理错误，可能需要重试或放弃
			log.Printf("Error monitoring transaction: %v", err)
			// 考虑是否需要在这里添加一个短暂的延迟，以避免在出错时过于频繁地重试
			time.Sleep(time.Second * 5)
			continue
		}
		log.Println(varifiedInfo)

		if varifiedInfo.status != "UNCONFIRMED" {
			s.completeTransaction(listener.OrderID, listener.BuyerID)
			close(listener.CompleteChan)
			s.nftService.nftRepo.IncrementNFTCount(listener.NFTID, "transaction_count")
			statusChannel, exists := common.TxStatusChannels.Get(listener.TxHash)
			if !exists {
				// 记录错误或返回错误
				log.Printf("No status channel found for transaction %s", listener.TxHash)
				return
			}
			select {
			case statusChannel <- "confirmed":
				log.Printf("send confirmed succeed")
			default:
				log.Printf("send confirmed failed, channel is full")
			}
			return
		}

		// 在循环的最后等待下一个 tick 或停止信号
		select {
		case <-ticker.C:
			// 继续下一次循环
		case <-listener.StopChan:
			return
		}
	}
}

func (s *NFTTrade) completeTransaction(orderID, buyerID uint) error {
	order, err := s.orderRepo.GetOrderByID(orderID)
	if err != nil {
		return err
	}
	// 3. 更新NFT表单
	if err := s.nftService.TransferNFT(order.NFT.ID, order.SellerID, buyerID); err != nil {
		return err
	}

	// 4. 更新订单表单
	if err := s.orderRepo.CompleteOrder(orderID, buyerID); err != nil {
		return err
	}

	// 更新交易状态
	if err := s.tracRepo.UpdateTransactionStatus(orderID, "COMPLETED"); err != nil {
		return err
	}

	return nil
}

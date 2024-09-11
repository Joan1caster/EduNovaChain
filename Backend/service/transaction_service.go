package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"nftPlantform/api"
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
	orderRepo         api.OrderRepository
	transactionRepo   api.TransactionRepository
	userRepo          api.UserRepository
	orderService      *OrderService
	nftService        *NFTService
	blockchainService *Blockchainservice
	listeners         map[string]*TransactionListener
	mu                sync.Mutex
}

func NewNFTTrade(
	orderService *OrderService,
	nftService *NFTService,
	blockchainService *Blockchainservice,
) *NFTTrade {
	return &NFTTrade{
		orderService: orderService,
		nftService:   nftService,
	}
}

func (s *NFTTrade) ExecuteTrade(ctx context.Context, orderID uint, buyerID uint, txHash string) error {
	// Step 1: 验证订单状态
	order, err := s.orderService.GetOrderByID(orderID)
	if err != nil {
		return fmt.Errorf("获取订单失败: %w", err)
	}
	if err := s.orderService.ValidateOrderStatus(orderID, order.SellerID); err != nil {
		return fmt.Errorf("验证订单状态失败: %w", err)
	}

	// Step 2: 创建交易记录并启动交易监听
	_, err = s.transactionRepo.CreateTransaction(orderID, txHash, strconv.FormatFloat(order.Price, 'f', -1, 64), "0", "PENDING")
	if err != nil {
		return fmt.Errorf("创建交易记录失败: %w", err)
	}
	go s.startTransactionListener(order.NFTID, orderID, buyerID, txHash)

	// Step 3: 在数据库中转移NFT
	buyer, err := s.userRepo.GetUserByID(buyerID)
	if err != nil {
		// 如果获取买家信息失败，将交易标记为失败
		if updateErr := s.transactionRepo.UpdateTransactionStatus(orderID, "FAILED"); updateErr != nil {
			return fmt.Errorf("获取买家信息失败且更新交易状态失败: %v, %w", updateErr, err)
		}
		return fmt.Errorf("获取买家信息失败: %w", err)
	}
	if err := s.nftService.TransferNFT(order.NFTID, orderID, buyer.ID); err != nil {
		// 如果NFT转移失败，将交易标记为失败
		if updateErr := s.transactionRepo.UpdateTransactionStatus(orderID, "FAILED"); updateErr != nil {
			return fmt.Errorf("NFT转移失败且更新交易状态失败: %v, %w", updateErr, err)
		}
		return fmt.Errorf("NFT转移失败: %w", err)
	}

	// Step 4: 标记订单为完成
	if err := s.orderRepo.CompleteOrder(orderID, buyerID); err != nil {
		// 如果标记订单完成失败，尝试重新打开订单
		if reopenErr := s.orderRepo.ReopenOrder(orderID); reopenErr != nil {
			return fmt.Errorf("标记订单完成失败且重新打开订单失败: %v, %w", reopenErr, err)
		}
		return fmt.Errorf("标记订单完成失败: %w", err)
	}

	return nil
}

func (s *NFTTrade) startTransactionListener(nftID, orderID, buyerID uint, txHash string) {
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
		select {
		case <-ticker.C:
			varifiedInfo, err := s.blockchainService.MonitorTransaction(listener)
			if err != nil {
				// 处理错误，可能需要重试或放弃
				continue
			}
			log.Println(varifiedInfo)
			if varifiedInfo.status != "UNCONFIRMED" {
				s.completeTransaction(listener.OrderID, listener.BuyerID)
				close(listener.CompleteChan)
				s.nftService.nftRepo.IncrementNFTCount(listener.NFTID, "transaction_count")
				return
			}
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
	if err := s.transactionRepo.UpdateTransactionStatus(orderID, "COMPLETED"); err != nil {
		return err
	}

	return nil
}

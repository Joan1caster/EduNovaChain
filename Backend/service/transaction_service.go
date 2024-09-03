package service

import (
	"context"
	"errors"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/lysu/go-saga"

	"nftPlantform/api"
)

type TransactionListener struct {
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

	sagaInstance := saga.NewSEC()

	sagaInstance.AddSubTxDef(
		"tradition step 1: verify order status",
		func(ctx context.Context) error {
			order, err := s.orderRepo.GetOrderByID(buyerID)
			if err != nil {
				return err
			}
			return s.orderService.ValidateOrderStatus(orderID, order.Buyer.WalletAddress)
		},
		nil, // 无需补偿操作
	)

	sagaInstance.AddSubTxDef(
		"tradition step 2: lesson to blockchain, update database",
		func(ctx context.Context) error {
			order, err := s.orderRepo.GetOrderByID(orderID)
			if err != nil {
				return err
			}

			_, err = s.transactionRepo.CreateTransaction(orderID, txHash, strconv.FormatFloat(order.Price, 'f', -1, 64), "0", "PENDING")
			if err != nil {
				return err
			}
			s.startTransactionListener(orderID, buyerID, txHash)
			return err
		},
		func(ctx context.Context) error {
			// 补偿操作：将交易标记为失败
			return s.transactionRepo.UpdateTransactionStatus(orderID, "FAILED")
		},
	)

	sagaInstance.AddSubTxDef(
		"tradition step 3: transfer nft in database",
		func(ctx context.Context) error {
			order, err := s.orderRepo.GetOrderByID(orderID)
			if err != nil {
				return err
			}
			buyer, err := s.userRepo.GetUserByID(buyerID)
			if err != nil {
				return err
			}
			return s.nftService.TransferNFT(order.NFTID, orderID, buyer.ID)
		},
		nil,
	)

	sagaInstance.AddSubTxDef(
		"tradition step 4: mark order as complete",
		func(ctx context.Context) error {
			return s.orderRepo.CompleteOrder(orderID, buyerID)
		},
		func(ctx context.Context) error {
			// 补偿操作：重新打开订单
			return s.orderRepo.ReopenOrder(orderID)
		},
	)

	// 执行Saga
	err := sagaInstance.StartSaga(ctx, 0)
	if err != nil {
		// Saga执行失败，触发补偿操作
		compensateErr := sagaInstance.StartCoordinator()
		if compensateErr != nil {
			return errors.New("交易失败，补偿操作也失败: " + compensateErr.Error())
		}
		return errors.New("transaction failed")
	}

	return nil
}

func (s *NFTTrade) startTransactionListener(orderID, buyerID uint, txHash string) {
	listener := &TransactionListener{
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

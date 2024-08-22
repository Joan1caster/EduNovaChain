package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"nftPlantform/api"
)

type NFTMarketplaceService struct {
	nftRepo         api.NFTRepository
	orderRepo       api.OrderRepository
	transactionRepo api.TransactionRepository
	userRepo        api.UserRepository
}

func NewNFTMarketplaceService(
	nftRepo api.NFTRepository,
	orderRepo api.OrderRepository,
	transactionRepo api.TransactionRepository,
	userRepo api.UserRepository,
) *NFTMarketplaceService {
	return &NFTMarketplaceService{
		nftRepo:         nftRepo,
		orderRepo:       orderRepo,
		transactionRepo: transactionRepo,
		userRepo:        userRepo,
	}
}

func (s *NFTMarketplaceService) ListNFTForSale(sellerID, nftID uint, price float64) (uint, error) {
	// Check if the NFT exists and belongs to the seller
	nft, err := s.nftRepo.GetNFTByID(nftID)
	if err != nil {
		return 0, err
	}
	if nft.OwnerID != sellerID {
		return 0, errors.New("seller does not own this NFT")
	}

	// Create a new order
	orderID, err := s.orderRepo.CreateOrder(sellerID, nftID, price)
	if err != nil {
		return 0, err
	}

	return orderID, nil
}

func (s *NFTMarketplaceService) CreateTransaction(orderID uint, txHash, amount string) error {
	// 创建初始交易记录
	_, err := s.transactionRepo.CreateTransaction(orderID, txHash, amount, "0", "PENDING")
	if err != nil {
		return err
	}

	// 启动一个 goroutine 来监听交易确认
	go s.monitorTransaction(txHash, orderID)

	return nil
}

func (s *NFTMarketplaceService) monitorTransaction(txHash string, orderID uint) {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR-PROJECT-ID")
	if err != nil {
		log.Printf("Failed to connect to the Ethereum client: %v", err)
		return
	}

	for {
		time.Sleep(15 * time.Second) // 每15秒检查一次
		tx, isPending, err := client.TransactionByHash(context.Background(), common.HexToHash(txHash))
		if err != nil {
			log.Printf("Failed to fetch transaction: %v", err)
			continue
		}

		if !isPending {
			receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				log.Printf("Failed to get transaction receipt: %v", err)
				continue
			}

			var status string
			if receipt.Status == 1 {
				status = "COMPLETED"
			} else {
				status = "FAILED"
			}

			gasUsed := new(big.Int).SetUint64(receipt.GasUsed)
			gasPrice := tx.GasPrice()
			gasFee := new(big.Int).Mul(gasUsed, gasPrice)

			gasFeeFloat, _ := gasFee.Float64()
			gasFeeEther := gasFeeFloat / 1e18 // Convert wei to ether

			err = s.transactionRepo.UpdateTransactionStatus(orderID, status)
			if err != nil {
				log.Printf("Failed to update transaction status: %v", err)
			}

			err = s.transactionRepo.UpdateTransactionGasFee(orderID, fmt.Sprint(gasFeeEther, 'f', -1, 64))
			if err != nil {
				log.Printf("Failed to update transaction gas fee: %v", err)
			}

			// 交易已确认，退出循环
			break
		}
	}
}

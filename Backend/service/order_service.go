package service

import (
	"context"
	"errors"
	"math/big"
	"nftPlantform/api"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type NFTMarketplaceService struct {
	nftRepo         api.NFTRepository
	orderRepo       api.OrderRepository
	transactionRepo api.TransactionRepository
	userBalanceRepo api.UserBalanceRepository
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

func (s *NFTMarketplaceService) BuyNFT(buyerAddress string, orderID uint) error {
    // Get the order
    order, err := s.orderRepo.GetOrderByID(orderID)
    if err != nil {
        return err
    }
    if order.Status != "OPEN" {
        return errors.New("order is not open")
    }

    // Connect to Ethereum network
    client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR-PROJECT-ID")
    if err != nil {
        return err
    }

    // Check buyer's balance
    balance, err := client.BalanceAt(context.Background(), common.HexToAddress(buyerAddress), nil)
    if err != nil {
        return err
    }

    price := new(big.Int)
    priceStr := strconv.FormatFloat(order.Price, 'f', -1, 64)
	price.SetString(priceStr, 10)
    if balance.Cmp(price) < 0 {
        return errors.New("insufficient balance")
    }

    // On-Chain transaction
	

    // Update NFT ownership
    nft, err := s.nftRepo.GetNFTByID(order.NFTID)
    if err != nil {
        return err
    }	
    nft.OwnerID = *order.BuyerID
    err = s.nftRepo.UpdateNFT(nft)
    if err != nil {
        return err
    }

    // Complete the order
    err = s.orderRepo.CompleteOrder(orderID, *order.BuyerID)
    if err != nil {
        return err
    }

    return nil
}
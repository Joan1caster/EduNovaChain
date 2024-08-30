package service

import (
	"context"
	"errors"
	"log"
	"math/big"
	configs "nftPlantform/config"

	"github.com/ethereum/go-ethereum/common"

	"nftPlantform/models"
)

type Blockchainservice struct {
}

func NewBlockchainservice() *Blockchainservice {
	return &Blockchainservice{}
}

type VarifiedInfo struct {
	status      string
	gasUsed     *big.Int
	gasPrice    *big.Int
	gasFeeEther float64
}

func (s *Blockchainservice) MonitorTransaction(listener *TransactionListener) (VarifiedInfo, error) {
	var varifiedInfo VarifiedInfo
	client, err := models.GetClient(configs.AppConfig.Contract.Eth_rpc_url)
	if err != nil {
		log.Printf("Failed to connect to the Ethereum client: %v", err)
		return varifiedInfo, err
	}
	tx, isPending, err := client.TransactionByHash(context.Background(), common.HexToHash(listener.TxHash))
	if err != nil {
		return varifiedInfo, errors.New("failed to fetch transaction")
	}

	if !isPending {
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			return varifiedInfo, errors.New("failed to get transaction receipt")
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
		// 交易已确认，退出循环
		return VarifiedInfo{
			status:      status,
			gasUsed:     gasUsed,
			gasFeeEther: gasFeeEther,
			gasPrice:    gasPrice,
		}, err
	} else {
		varifiedInfo.status = "UNCONFIRMED"
		return varifiedInfo, nil
	}
}

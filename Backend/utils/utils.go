package utils
import (
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
)

func VerifySignature(address, signature, message string) bool {
    addr := common.HexToAddress(address)
    sig := common.FromHex(signature)
    msgHash := crypto.Keccak256([]byte(message))
    pubKey, err := crypto.SigToPub(msgHash, sig)
    if err != nil {
        return false
    }
    recoveredAddr := crypto.PubkeyToAddress(*pubKey)
    return addr == recoveredAddr
}
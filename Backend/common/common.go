package common

import (
    "sync"
)

var (
    // TxStatusChannels 存储每个交易的状态 channel
    TxStatusChannels = &TxChannelMap{
        channels: make(map[string]chan string),
    }
)

// TxChannelMap 提供了一个线程安全的方式来管理交易状态 channels
type TxChannelMap struct {
    mu       sync.RWMutex
    channels map[string]chan string
}

// Get 返回指定交易哈希的 channel，如果不存在则返回 false
func (tcm *TxChannelMap) Get(txHash string) (chan string, bool) {
    tcm.mu.RLock()
    defer tcm.mu.RUnlock()
    ch, exists := tcm.channels[txHash]
    return ch, exists
}

// Set 为指定的交易哈希设置一个新的 channel
func (tcm *TxChannelMap) Set(txHash string, ch chan string) {
    tcm.mu.Lock()
    defer tcm.mu.Unlock()
    tcm.channels[txHash] = ch
}

// Delete 删除指定交易哈希的 channel
func (tcm *TxChannelMap) Delete(txHash string) {
    tcm.mu.Lock()
    defer tcm.mu.Unlock()
    delete(tcm.channels, txHash)
}

// NewTxStatusChannel 创建一个新的交易状态 channel 并添加到 map 中
func NewTxStatusChannel(txHash string) chan string {
    ch := make(chan string, 1)
    TxStatusChannels.Set(txHash, ch)
    return ch
}
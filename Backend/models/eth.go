package models

import (
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/time/rate"
)

var (
	Client  *ethclient.Client
	once    sync.Once
	limiter *rate.Limiter
)

func GetClient(url string) (*ethclient.Client, error) {
	var err error
	once.Do(func() {
		Client, err = ethclient.Dial(url)
		// 创建一个限流器，每秒允许100个请求，最多允许积累500个令牌
		limiter = rate.NewLimiter(rate.Limit(100), 500)
	})
	return Client, err
}

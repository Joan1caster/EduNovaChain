package repository

import (
	"context"
	"time"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type HeatRepository struct {
	redisDB *redis.Client
}

func NewHeatRepository(redisDB *redis.Client) *HeatRepository {
	return &HeatRepository{redisDB: redisDB}
}

var ctx = context.Background()

// Increase the number of views for productID products
func(r *HeatRepository) IncrementViewCount(productID string) {
    // 当前日期
    dayKey := time.Now().Format("2006-01-02")

    // 使用Redis哈希表存储不同时间周期的浏览量
    r.redisDB.HIncrBy(ctx, productID, "view:day:"+dayKey, 1)
	r.redisDB.Expire(ctx, dayKey, 30*24*time.Hour)
}

// Increase the number of transactions for productID products
func(r *HeatRepository) IncrementTraditionCount(productID string) {
    // 当前日期
    dayKey := time.Now().Format("2006-01-02")

    // 使用Redis哈希表存储不同时间周期的浏览量
    r.redisDB.HIncrBy(ctx, productID, "transaction:day:"+dayKey, 1)
	r.redisDB.Expire(ctx, dayKey, 30*24*time.Hour)
}

// query the number of views/transactions in some day.
func (r *HeatRepository) GetCountForDays(productID,content string, days int) (int64, error) {
    currentDate := time.Now()
    
    totalViews := int64(0)
    
    for i := 0; i < days; i++ {
        date := currentDate.AddDate(0, 0, -i)
        dayKey := date.Format("2006-01-02")
        
        key := fmt.Sprintf("%s:day:%s",content ,dayKey)
        
        views, err := r.redisDB.HGet(context.Background(), productID, key).Int64()
        if err != nil {
            if err == redis.Nil {
                continue
            }
            return 0, err
        }
        
        totalViews += views
    }
    
    return totalViews, nil
}
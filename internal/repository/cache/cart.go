package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hespecial/gin-mall/global"
	"github.com/hespecial/gin-mall/internal/api/response"
	"github.com/hespecial/gin-mall/internal/common/constant"
	"github.com/redis/go-redis/v9"
)

func GetCartList(userID uint) (items []*response.CartItem, exist bool, err error) {
	key := fmt.Sprintf("%s:%d", constant.CartKey, userID)

	ctx := context.Background()
	itemsJson, err := global.Redis.Get(ctx, key).Result()
	if err != nil {
		// key不存在，将err置为nil，允许mysql进行查询
		if errors.Is(err, redis.Nil) {
			return nil, false, nil
		}
		// 查询出错，返回错误
		return nil, false, err
	}

	err = json.Unmarshal([]byte(itemsJson), &items)

	return items, true, err
}

func SaveCartItems(userID uint, items []*response.CartItem) error {
	key := fmt.Sprintf("%s:%d", constant.CartKey, userID)

	// 序列化购物项列表为JSON
	itemsJson, err := json.Marshal(items)
	if err != nil {
		return err
	}

	// 将购物车数据存储到Redis
	ctx := context.Background()
	// 设置过期时间保证最终一致性
	return global.Redis.Set(ctx, key, itemsJson, constant.CartKeyExpire).Err()
}

func DeleteCartCache(userID uint) error {
	key := fmt.Sprintf("%s:%d", constant.CartKey, userID)

	return global.Redis.Del(context.Background(), key).Err()
}

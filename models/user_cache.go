package models

import (
	"context"
	"gochat/utils"
	"time"
)

/*
*
redis cache
*
*/
func SetUserOnlineInfo(key string, val []byte, timeTTL time.Duration) {
	ctx := context.Background()
	utils.Red.Set(ctx, key, val, timeTTL)
}

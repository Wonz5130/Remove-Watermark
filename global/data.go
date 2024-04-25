package global

import (
	"math/rand"
	"os"
	"strconv"
	"time"
	// dataHandler "gitlab.yc345.tv/backend/course-pkg/data"
	// redis_pkg "gitlab.yc345.tv/backend/course-pkg/redis"
	// "gorm.io/gorm"
)

var (
// DB *gorm.DB
// Redis       *redis_pkg.RedisClient
// DataHandler *dataHandler.DataHandler
// Vm *gorm.DB
)

// GetWorkerNum 并发数控制
func GetWorkerNum() int {
	workerNum := 5
	if GoEnv == `stage` {
		workerNum = 10
	}
	if GoEnv == `production` {
		workerNum = 10
	}
	return workerNum
}

func GetCacheTime(productionExpireTime int) time.Duration {
	rand.Seed(time.Now().Unix())
	expireTime := 5
	if getEnv() == `production` {
		expireTime = productionExpireTime
	}
	if getEnv() == `stage` {
		expireTime = 5
	}
	REDIS_EXPIRED_MINUTES := os.Getenv(`REDIS_EXPIRED_MINUTES`)
	if REDIS_EXPIRED_MINUTES != `` {
		expireTime, _ = strconv.Atoi(REDIS_EXPIRED_MINUTES)
	}
	return time.Minute * time.Duration(expireTime)
}

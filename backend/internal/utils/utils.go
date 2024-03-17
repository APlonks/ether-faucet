package utils

import (
	"context"
	"fmt"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/redis/go-redis/v9"
)

// IsValidAddress validate hex address
func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

func GetSimuRunning(redis_url string) (bool, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redis_url,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println("The redis url:", redis_url)
	ctx := context.Background()

	RedisSimuRunning, err := redisClient.Get(ctx, "RedisSimuRunning").Result()
	if err == redis.Nil {
		fmt.Println("The key doesn't exist ")
		errSet := redisClient.Set(ctx, "RedisSimuRunning", "false", 0).Err()
		if errSet != nil {
			fmt.Println("Problem while initializing RedisSimuRunning in Redis")
		}
		RedisSimuRunning, err = redisClient.Get(ctx, "RedisSimuRunning").Result()

	} else if err != nil {
		fmt.Println("Error while to get RedisSimuRunning's value in Redis :", err)
	} else {
		fmt.Println("RedisSimuRunning's value:", RedisSimuRunning)
	}

	return RedisSimuRunning == "true", err
}

func SetSimuRunning(redis_url string, simuRunning string) error {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redis_url,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println("The redis url:", redis_url)
	ctx := context.Background()

	err := redisClient.Set(ctx, "RedisSimuRunning", simuRunning, 0).Err()
	if err != nil {
		fmt.Println("Problem while initializing RedisSimuRunning in Redis")
		return err
	}
	return err
}

// Not used
func ErrManagement(err error) {
	if err != nil {
		// log.Fatal("!! ERROR !!:", err)
		fmt.Println("!! ERROR !!:", err)
	}
}

package redis

import (
	"github.com/go-redis/redis"
)

func keyNotExists(service, url string) bool {
	_, err := client.ZRank(service, url).Result()
	if err != nil {
		return true
	}
	return false
}

// RegisterService puts a service URL in its respective sorted set if it doesn't exist
// for service discovery
func RegisterService(service, url string, score float64) error {
	if keyNotExists(service, url) {
		_, err := client.ZAdd(service, redis.Z{Score: score, Member: url}).Result()
		return err
	}
	return nil
}

// IncrementServiceLoad increments the number of apps deployed on a service host by 1
func IncrementServiceLoad(service, url string) error {
	_, err := client.ZIncrBy(service, 1, url).Result()
	return err
}

// GetLeastLoadedInstance returns the URL of the host currently having the least number
// of apps of a particular service deployed
func GetLeastLoadedInstance(service string) (string, error) {
	data, err := client.ZRangeByScore(
		service,
		redis.ZRangeBy{
			Min:    "-inf",
			Max:    "+inf",
			Offset: 0,
			Count:  1,
		}).Result()
	if err != nil {
		return "", err
	}
	if len(data) == 0 {
		return "Empty Set", nil
	}
	return data[0], nil
}

// FetchServiceInstances returns all instances of a given service
func FetchServiceInstances(service string) ([]string, error) {
	data, err := client.ZRangeByScore(
		service,
		redis.ZRangeBy{
			Min:    "-inf",
			Max:    "+inf",
			Offset: 0,
		}).Result()
	if err != nil {
		return []string{}, err
	}
	if len(data) == 0 {
		return []string{}, nil
	}
	return data, nil
}

// RemoveServiceInstance removes an instance of a particular service
func RemoveServiceInstance(service, member string) error {
	_, err := client.ZRem(service, member).Result()
	if err != nil {
		return err
	}
	return nil
}
package cache

import (
	"log"
	cache "vacanciesParser/internal/core/repository/redis"
)

func GetITRolesIDs() []string {
	cachedResult, err := cache.Rdb.LRange(cache.Ctx, "hhITRolesIDs", 0, -1).Result()
	if err != nil {
		log.Printf("Error while getting cache: %v", err)
	}

	return cachedResult
}

func SetITRolesIDs(vals []string) {
	cache.Rdb.RPush(cache.Ctx, "hhITRolesIDs", vals)
}

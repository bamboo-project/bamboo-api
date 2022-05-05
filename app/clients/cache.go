package clients

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var Localcache *cache.Cache

func InitCache() {
	Localcache = cache.New(5*time.Minute, 10*time.Minute)
}

package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	Entry map[string]cacheEntry
	Mutex	sync.Mutex
}

type cacheEntry struct {
	createdAt	time.Time
	val	[]byte
}

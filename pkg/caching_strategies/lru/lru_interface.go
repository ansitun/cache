package lru

type LRUCache interface {
	Get(key string) (value string, ok bool) // value of type interface and ok-true or false
	Set(key string, value string) bool      // send true if eviction occurred
	Len() int                               // number of elements int the cache
	Print()
}

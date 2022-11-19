package main

import (
	"cache/pkg/caching_strategies/lru"
	"fmt"
)

// the directory name for each application should match the name of the executable you want to have (e.g., /cmd/cache).

func main() {
	fmt.Println("hello cache")
	cache, _ := lru.NewLRU(4)
	cache.Set("1", "hello")
	cache.Set("12", "2hello")
	cache.Set("123", "3hello")
	cache.Set("1234", "4hello")
	cache.Set("1235", "5hello")

	cache.Delete("1")

	cache.Delete("12")

	cache.Print()
}

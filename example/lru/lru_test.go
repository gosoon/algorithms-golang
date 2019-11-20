package lru

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache(2)
	cache.Put("name", "1")
	fmt.Println("--------------> list  1")
	cache.List()

	cache.Put("age", "2")
	fmt.Println("--------------> list  2")
	cache.List()

	cache.Put("name", "2")
	fmt.Println("--------------> list  2")
	cache.List()

	cache.Put("data", "3")
	fmt.Println("--------------> list  2")
	cache.List()

	cache.Put("data2", "3")
	fmt.Println("--------------> list  2")
	cache.List()
}

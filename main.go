package main

import (
	"fmt"
	// "hw1/cache"
	"github.com/Aizengott/_udemi_hw1/cache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 44)
	cache.Set("Alice", 22)
	userId := cache.Get("Alice")

	fmt.Println(userId)

	cache.Delete("userId")
	userId = cache.Get("userId")

	fmt.Println(userId)
}

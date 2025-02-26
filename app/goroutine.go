package main

import (
	"fmt"
	cache "github.com/kiricle/in-memory-cache"
	"time"
)

func main() {
	c := cache.New()

	c.Set("key", "value", time.Second*2)

	time.Sleep(time.Second * 3)

	fmt.Println("val:", c.Get("key"))
}

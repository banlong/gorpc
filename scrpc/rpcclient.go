package main

import (
	"log"
	"time"
	"gorpc/context"
)

var (
	c   *context.Client
	err error
	dsn       = "localhost:9876"
	cacheItem = &context.CacheItem{Key: "some key", Value: "some value"}
)

func main() {
	c, err = context.NewClient(dsn, time.Millisecond*500)
	if err != nil {
		log.Fatal(err)
	}else{
		log.Println("Connected to the remote server")
	}

	//Put Item
	_, err := c.Put(cacheItem)
	if err != nil {
		log.Println(err)
	}

	//Get Item
	item, _ := c.Get(cacheItem.Key)
	if item == nil {
		log.Printf("Cache key should not exist: %s\n", cacheItem.Key)
	}else{
		log.Printf("Cache Get: %s\n", item.Value)
	}

}


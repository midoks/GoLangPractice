package main

import (
	"fmt"

	redis "github.com/alphazero/Go-Redis"
)

func main() {

	rdb := redis.DefaultSpec().Host("127.0.0.1").Db(0).Password("")
	client, err := redis.NewSynchClientWithSpec(rdb)

	if err != nil {
		fmt.Println("connect redis server fail")
		return
	}

	dbkey := "test_1"
	value := []byte("hello world")
	client.Set(dbkey, value)

	getValue, err := client.Get(dbkey)

	if err != nil {
		fmt.Println("Get Key Fail")
	} else {
		str := string(getValue)
		fmt.Println(str)
	}

	fmt.Println("test")
}

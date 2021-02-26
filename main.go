package main

import (
	"fmt"
	"goft-redis/StringOperation"
	"time"
)

var (
	RedisString = StringOperation.NewStringOperation()
	withExpire  = StringOperation.WithExpire(time.Second * 20) //过期时间
	withNX      = StringOperation.WithNX()
	withXX      = StringOperation.WithXX()
)

func main() {
	no3()
}

func no3() {
	fmt.Println(RedisString.Set("test1", "eeeee", withExpire, withXX).Unwrap())
}

func no2() {
	iter := StringOperation.NewStringOperation().
		Mget("name", "age", "abc").Iter()

	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
}

func no1() {
	fmt.Println(StringOperation.NewStringOperation().Get("abc").Unwrap_Or("default：sunhaiming"))
	fmt.Println(StringOperation.NewStringOperation().Get("name").Unwrap())
}

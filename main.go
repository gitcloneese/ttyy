package main

import (
	"context"
	"fmt"
	v9 "github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
	"log"
	"ttyy/proto/user"
)

func main0001() {

	client := v9.NewClient(&v9.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Println(err)
		return
	}

	//pas := "nhao"
	t := &user.User{
		//UserName: "nihao",
		//Age:      12,
		//Sex:      []bool{true, false},
		//Password: &pas,
	}
	b, err := proto.Marshal(t)
	if err != nil {
		fmt.Println("proto:", err)
		return
	}
	t1 := new(user.User)
	err1 := proto.Unmarshal(b, t1)
	fmt.Println("err1:", err1)

	err2 := client.Set(context.Background(), "test", string(b), 0).Err()
	if err2 != nil {
		log.Println(err2)
	}
}

func main() {

	client := v9.NewClient(&v9.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Println(err)
		return
	}
	b, err2 := client.Get(context.Background(), "test").Bytes()
	if err2 != nil {
		log.Println(err2)
		return
	}
	t1 := new(user.User)
	err3 := proto.Unmarshal(b, t1)
	if err3 != nil {
		log.Println(err3)
	}
}

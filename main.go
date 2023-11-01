package main

import (
	"context"
	"fmt"
	v9 "github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
	anypb "google.golang.org/protobuf/types/known/anypb"
	"log"
	"ttyy/proto/user"
)

func main() {

	//client := v9.NewClient(&v9.Options{
	//	Addr:     "localhost:6379",
	//	Password: "", // no password set
	//	DB:       0,  // use default DB
	//})
	//_, err := client.Ping(context.Background()).Result()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	//pas := "nhao"

	bt, _ := proto.Marshal(&user.Test{
		TestName: "test",
	})
	bt1, _ := proto.Marshal(&user.Test1{
		Age: 10000000,
	})

	t := &user.User{
		UserName: "nihao",
		Age:      12,
		Sex:      user.Sex_S_FEMALE,
		//Password: &pas,
		Data1: &anypb.Any{

			TypeUrl: "Test",
			Value:   bt,
		},
		Data2: &anypb.Any{

			TypeUrl: "Test1",
			Value:   bt1,
		},
	}
	b, err := proto.Marshal(t)
	if err != nil {
		fmt.Println("proto:", err)
		return
	}
	t1 := new(user.User)
	err1 := proto.Unmarshal(b, t1)
	fmt.Println("err1:", err1)

	tTest := new(user.Test)
	tTest1 := new(user.Test1)
	err3 := proto.Unmarshal(t1.Data1.Value, tTest)
	fmt.Println(err3)
	err4 := proto.Unmarshal(t1.Data2.Value, tTest1)
	fmt.Println(err4)
	//err2 := client.Set(context.Background(), "test", string(b), 0).Err()
	//if err2 != nil {
	//	log.Println(err2)
	//}
}

func main999() {

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

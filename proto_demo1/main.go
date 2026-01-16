package main

import (
	"fmt"
	"proto_demo1/proto/userService"

	"google.golang.org/protobuf/proto"
)

func main() {
	userinfo := &userService.UserInfo{
		Age:     1,
		Name:    "张三",
		Hobbies: []string{"吃饭", "睡觉"},
	}
	fmt.Println(userinfo.GetAge())
	fmt.Println(userinfo.GetName())
	fmt.Println(userinfo.GetHobbies())

	byte, _ := proto.Marshal(userinfo)
	fmt.Println(byte)
	u := new(userService.UserInfo)
	proto.Unmarshal(byte, u)
	fmt.Println(u)
}

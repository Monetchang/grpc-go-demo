package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/api/protobuf-spec/user"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("connect failure %s", err)
		return
	}
	defer conn.Close()

	// create client
	cli := user.NewMainServiceClient(conn)

	ret, err := cli.GetUserName(context.Background(), &user.UserReq{UserId: 13})
	if err != nil {
		fmt.Printf("call srever func failure %s", err)
		return
	}
	fmt.Printf("call success %s /n", ret.UserName)
}

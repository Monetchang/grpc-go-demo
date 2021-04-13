package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/api/protobuf-spec/user"
	"net"
	"strconv"
)

type server struct {}

func (s *server) GetUserName(ctx context.Context, request *user.UserReq) (*user.UserResp, error){
	return &user.UserResp{UserName: "name" + strconv.FormatInt(int64(request.UserId), 10)}, nil
}

func main() {
	fmt.Print("start server...")
	// listen port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("listen port error: %s", err)
		return
	}
	
	// new api server
	grpcServer := grpc.NewServer()
	user.RegisterMainServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("Server start error: %s", err)
		return
	}
}

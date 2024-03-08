package client

import (
	"fmt"
	"github.com/15733012783/mysql/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Client(serverName, Address string) (*grpc.ClientConn, error) {
	conn, err := consul.GetClient(serverName, Address)
	if err != nil {
		return nil, err
	}
	fmt.Println(conn)
	return grpc.Dial(conn, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

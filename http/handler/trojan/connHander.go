package trojan

import (
	"github.com/p4gefau1t/trojan-go/api/service"
	"google.golang.org/grpc"
)

func NewHsClient() (*grpc.ClientConn, service.TrojanServerServiceClient) {
	conn, err := grpc.Dial("18.140.56.139:9999", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	hsClient := service.NewTrojanServerServiceClient(conn)
	return conn, hsClient
}

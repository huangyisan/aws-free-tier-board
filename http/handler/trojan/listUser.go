package trojan

import (
	"context"
	"fmt"
	"github.com/p4gefau1t/trojan-go/api/service"
	"google.golang.org/grpc"
	"io"
	"log"
)

func listUsers(conn *grpc.ClientConn, c service.TrojanServerServiceClient) {
	stream, err := c.ListUsers(context.Background(), &service.ListUsersRequest{})
	defer conn.Close()
	if err != nil {
		log.Printf("failed to call grpc command: %v", err)
	}
	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("faild to recv: %v", err)
		}
		fmt.Println(reply.Status.TrafficTotal.DownloadTraffic)
		break
	}
}

func PrintListUser() {
	listUsers(NewHsClient())
}

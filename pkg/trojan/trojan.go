package trojan

import (
	"context"
	"fmt"
	"github.com/p4gefau1t/trojan-go/api/service"
	"google.golang.org/grpc"
	"io"
	"log"
)

type trojanClient struct {
	ip                        string
	port                      string
	grpcClientConn            *grpc.ClientConn
	trojanServerServiceClient service.TrojanServerServiceClient
}

func (t *trojanClient) setIp(ip string) {
	t.ip = ip
}

func (t *trojanClient) setPort(port string) {
	t.port = port
}

func (t *trojanClient) setGrpcClientConn() {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", t.ip, t.port), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	t.grpcClientConn = conn
}

func (t *trojanClient) setTrojanServerServiceClient() {
	t.trojanServerServiceClient = service.NewTrojanServerServiceClient(t.grpcClientConn)
}

func (t *trojanClient) Close() {
	t.grpcClientConn.Close()
}

func NewTrojanClient(ip, port string) *trojanClient {
	tc := &trojanClient{}
	tc.setIp(ip)
	tc.setPort(port)
	tc.setGrpcClientConn()
	tc.setTrojanServerServiceClient()
	return tc
}

func (t *trojanClient) ListAllTraffic() (downloadTraffic uint64, uploadTraffic uint64) {

	stream, err := t.trojanServerServiceClient.ListUsers(context.Background(), &service.ListUsersRequest{})
	if err != nil {
		panic(err)
	}
	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("faild to recv: %v", err)
		}
		downloadTraffic += reply.Status.TrafficTotal.DownloadTraffic
		uploadTraffic += reply.Status.TrafficTotal.UploadTraffic
	}
	return
}

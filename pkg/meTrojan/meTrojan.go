package meTrojan

import (
	"fmt"
	"github.com/p4gefau1t/trojan-go/api/service"
	"google.golang.org/grpc"
)

type trojanClient struct {
	tag                       string
	ip                        string
	port                      string
	grpcClientConn            *grpc.ClientConn
	trojanServerServiceClient service.TrojanServerServiceClient
}

func (t *trojanClient) setTag(tag string) {
	t.tag = tag
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

package meTrojan

import (
	"context"
	"github.com/p4gefau1t/trojan-go/api/service"
	"io"
	"time"
	"trojan-dashboard/pkg/common"
	"trojan-dashboard/pkg/logger"
	"trojan-dashboard/pkg/meInfluxdb"
)

func (t *trojanClient) ListAllTraffic() (downloadTraffic uint64, uploadTraffic uint64) {
	stream, err := t.trojanServerServiceClient.ListUsers(context.Background(), &service.ListUsersRequest{})
	if err != nil {
		logger.Failed.Msgf(err.Error())
		return 0, 0
	}
	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Failed.Msgf(err.Error())
		}
		downloadTraffic += reply.Status.TrafficTotal.DownloadTraffic
		uploadTraffic += reply.Status.TrafficTotal.UploadTraffic
	}
	return
}

func QueryAllTraffic(i *meInfluxdb.InfluxClient, startTime string) []queryAllTrafficResponse {
	defer i.Close()
	var results []queryAllTrafficResponse
	allTrafficQuery := meInfluxdb.MakeQuery(common.AllTrafficQuery, i.GetBucket(), startTime)
	result, err := i.QueryRecord(context.Background(), allTrafficQuery)
	if err != nil {
		logger.Failed.Msgf(err.Error())
	}
	for result.Next() {
		record := result.Record()
		results = append(results, queryAllTrafficResponse{
			Time:  record.Time().Format(time.RFC3339),
			Tag:   record.ValueByKey("tag").(string),
			Ip:    record.ValueByKey("ip").(string),
			Value: record.ValueByKey("total").(float64),
		})
	}
	return results
}

type queryAllTrafficResponse struct {
	Time  string  `json:"time"`
	Tag   string  `json:"tag"`
	Ip    string  `json:"ip"`
	Value float64 `json:"value"`
}

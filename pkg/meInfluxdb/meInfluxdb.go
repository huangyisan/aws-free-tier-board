package meInfluxdb

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type InfluxClient struct {
	url       string
	authToken string
	org       string
	bucket    string
	client    influxdb2.Client
	writeApi  api.WriteAPI
	queryAPI  api.QueryAPI
}

func (i *InfluxClient) setURL(url string) {
	i.url = url
}

func (i *InfluxClient) setAuthToken(authToken string) {
	i.authToken = authToken
}

func (i *InfluxClient) setOrg(org string) {
	i.org = org
}

func (i *InfluxClient) setBucket(bucket string) {
	i.bucket = bucket
}

func (i *InfluxClient) GetBucket() string {
	return i.bucket
}

func (i *InfluxClient) setClient() {
	i.client = influxdb2.NewClient(i.url, i.authToken)
}

func (i *InfluxClient) setWriteApi() {
	i.writeApi = i.client.WriteAPI(i.org, i.bucket)
}

func (i *InfluxClient) setQueryApi() {
	i.queryAPI = i.client.QueryAPI(i.org)
}

func (i *InfluxClient) WriteLineRecord(line string) {
	i.writeApi.WriteRecord(line)
}

func (i *InfluxClient) QueryRecord(cxt context.Context, query string) (*api.QueryTableResult, error) {
	result, err := i.queryAPI.Query(cxt, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (i *InfluxClient) Close() {
	i.client.Close()
}

func MakeQuery(queryString string, args ...interface{}) string {
	return fmt.Sprintf(queryString, args...)
}

func MakeLine(queryString string, args ...interface{}) string {
	return fmt.Sprintf(queryString, args...)
}

func NewInfluxDBClient(url, authToken, org, bucket string) *InfluxClient {
	i := &InfluxClient{}
	i.setURL(url)
	i.setAuthToken(authToken)
	i.setOrg(org)
	i.setBucket(bucket)
	i.setClient()
	i.setWriteApi()
	i.setQueryApi()
	return i
}

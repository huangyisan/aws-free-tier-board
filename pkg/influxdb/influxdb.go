package influxdb

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type influxClient struct {
	url       string
	authToken string
	org       string
	bucket    string
	client    influxdb2.Client
	writeApi  api.WriteAPI
	queryAPI  api.QueryAPI
}

func (i *influxClient) setURL(url string) {
	i.url = url
}

func (i *influxClient) setAuthToken(authToken string) {
	i.authToken = authToken
}

func (i *influxClient) setOrg(org string) {
	i.org = org
}

func (i *influxClient) setBucket(bucket string) {
	i.bucket = bucket
}

func (i *influxClient) setClient() {
	i.client = influxdb2.NewClient(i.url, i.authToken)
}

func (i *influxClient) setWriteApi() {
	i.writeApi = i.client.WriteAPI(i.org, i.bucket)
}

func (i *influxClient) setQueryApi() {
	i.queryAPI = i.client.QueryAPI(i.org)
}

func (i *influxClient) WriteLineRecord(line string) {
	i.writeApi.WriteRecord(line)
}

func (i *influxClient) QueryRecord(cxt context.Context, query string) (*api.QueryTableResult, error) {
	result, err := i.queryAPI.Query(cxt, query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (i *influxClient) Close() {
	i.client.Close()
}

func MakeQuery(queryString string, args ...interface{}) string {
	return fmt.Sprintf(queryString, args...)
}

func MakeLine(queryString string, args ...interface{}) string {
	return fmt.Sprintf(queryString, args...)
}

func NewInfluxDBClient(url, authToken, org, bucket string) *influxClient {
	i := &influxClient{}
	i.setURL(url)
	i.setAuthToken(authToken)
	i.setOrg(org)
	i.setBucket(bucket)
	i.setClient()
	i.setWriteApi()
	i.setQueryApi()
	return i
}

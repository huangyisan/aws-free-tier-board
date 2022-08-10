package meTrojan

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"testing"
	"trojan-dashboard/pkg/meInfluxdb"
)

func Test_influxdb(t *testing.T) {
	url := "http://abc/"
	token := "=="
	org := "qz"
	//bucket := "qz"
	client := influxdb2.NewClient(url, token)
	// get non-blocking write client
	//writeAPI := client.WriteAPI("qz", "qz")

	//// write line protocol
	//writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 100.5, 45.0))
	//// Flush writes
	//writeAPI.Flush()

	queryAPI := client.QueryAPI(org)
	query := `import "math"
		last_traffic = from(bucket: "qz") 
			|> range(start: %v)
			|> filter(fn: (r) => r["_measurement"] == "trojan") 
			|> filter(fn: (r) => r["_field"] == "download" or r["_field"] == "upload") 
			|> last(column: "_time")
		// last_traffic
		first_traffic = from(bucket: "qz") 
			|> range(start: %v)
			|> filter(fn: (r) => r["_measurement"] == "trojan") 
			|> filter(fn: (r) => r["_field"] == "download" or r["_field"] == "upload") 
			|> first(column: "_time")

		union(tables: [first_traffic, last_traffic])
		|> difference()
		|> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
		|> map(fn: (r) => ({ r with total: math.abs(x: r.download) + math.abs(x: r.upload)}))`
	newQuery := meInfluxdb.MakeQuery(query, "-1d")
	result, err := queryAPI.Query(context.Background(), newQuery)
	if err != nil {
		panic(err)
	}
	for result.Next() {
		record := result.Record()
		fmt.Printf("%v %v: %v=%v %v %v\n", record.Time(), record.Measurement(), record.Field(), record.Value(), record.ValueByKey("tag"), record.ValueByKey("total"))
	}
	client.Close()

}

//
//func Test_cc(t *testing.T) {
//	i := meInfluxdb.NewInfluxDBClient(
//		"http://abc/",
//		"cc",
//		"qz",
//		"qz")
//	a := QueryAllTraffic(i, common.GetMonthFirstDay())
//	fmt.Printf("%v\n", a)
//}

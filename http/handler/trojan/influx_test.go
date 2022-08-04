package trojan

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"testing"
)

func Test_influxdb(t *testing.T) {
	url := ""
	token := ""
	org := "qz"
	bucket := "qz"
	client := influxdb2.NewClient(url, token)
	// get non-blocking write client
	writeAPI := client.WriteAPI("qz", "qz")

	// write line protocol
	writeAPI.WriteRecord(fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 100.5, 45.0))
	// Flush writes
	writeAPI.Flush()

	queryAPI := client.QueryAPI(org)
	query := fmt.Sprintf(`from(bucket: "%v") |> range(start: -1d)`, bucket)
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		panic(err)
	}
	for result.Next() {
		record := result.Record()
		fmt.Printf("%v %v: %v=%v\n", record.Time(), record.Measurement(), record.Field(), record.Value())
	}
	client.Close()

}

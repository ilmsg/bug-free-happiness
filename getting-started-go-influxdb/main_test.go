package main

import (
	"context"
	"os"
	"reflect"
	"testing"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/domain"
	"github.com/joho/godotenv"
)

func Test_connectToInfluxDB(t *testing.T) {

	godotenv.Load("./test_influxdb.env")

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Successful connection to InfluxDB",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := connectToInfluxDB()

			if (err != nil) != tt.wantErr {
				t.Errorf("connectToInfluxDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			health, err := got.Health(context.Background())
			if (err != nil) && health.Status == domain.HealthCheckStatusPass {
				t.Errorf("connectionToTifluxDB() error. database not healthy")
			}
		})
	}
}

func Test_write_event_with_line_protocol(t *testing.T) {
	tests := []struct {
		name  string
		f     func(influxdb2.Client, []ThermostatSetting)
		datas []ThermostatSetting
	}{
		{
			name: "Write new record with line protocol",
			// Your data Points
			datas: []ThermostatSetting{{user: "foo", avg: 35.5, max: 42}},
			f: func(c influxdb2.Client, datas []ThermostatSetting) {
				// Send all the data to the DB
				for _, data := range datas {
					write_event_with_line_protocol(c, data)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := init_testDB(t)

			// call function to test
			tt.f(client, tt.datas)
			// test can be flicky if the query is done before that data is ready in the database
			time.Sleep(time.Millisecond * 1000)

			// Option one: QueryTableResult
			results := read_events_as_query_table_result(client)
			// convert results to array to compare with data
			resultsArr := []ThermostatSetting{}
			for _, v := range results {
				resultsArr = append(resultsArr, v)
			}

			if eq := reflect.DeepEqual(resultsArr, tt.datas); !eq {
				t.Errorf("want %v, got %v", tt.datas, resultsArr)
			}

			// Option two: query raw data
			// TODO add validation
			read_events_as_raw_string(client)

			client.Close()
		})
	}
}

func init_testDB(t *testing.T) influxdb2.Client {
	t.Helper()                           // Tells `go test` that this is an helper
	godotenv.Load("./test_influxdb.env") //load environement variable
	client, err := connectToInfluxDB()   // create the client

	if err != nil {
		t.Errorf("impossible to connect to DB")
	}

	bucket := os.Getenv("INFLUXDB_BUCKET")
	org := os.Getenv("INFLUXDB_ORG")

	// Clean the database by deleting the bucket
	ctx := context.Background()
	bucketsAPI := client.BucketsAPI()
	dBucket, err := bucketsAPI.FindBucketByName(ctx, bucket)
	if err == nil {
		client.BucketsAPI().DeleteBucketWithID(context.Background(), *dBucket.Id)
	}

	// create new empty bucket
	dOrg, _ := client.OrganizationsAPI().FindOrganizationByName(ctx, org)
	_, err = client.BucketsAPI().CreateBucketWithNameWithID(ctx, *dOrg.Id, bucket)

	if err != nil {
		t.Errorf("impossible to new create bucket")
	}

	return client
}

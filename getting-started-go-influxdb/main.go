package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/joho/godotenv"
)

type ThermostatSetting struct {
	user string
	max  float64 //temperature
	avg  float64 //temperature
}

func init() {
	godotenv.Load("./test_influxdb.env")
}

func main() {
	//
}

func connectToInfluxDB() (influxdb2.Client, error) {
	dbToken := os.Getenv("INFLUXDB_TOKEN")
	if dbToken == "" {
		return nil, errors.New("INFLUXDB_TOKEN must be set")
	}

	dbURL := os.Getenv("INFLUXDB_URL")
	if dbURL == "" {
		return nil, errors.New("INFLUXDB_URL must be set")
	}

	client := influxdb2.NewClient(dbURL, dbToken)

	// validate client connection health
	_, err := client.Health(context.Background())

	return client, err
}

func read_events_as_raw_string(client influxdb2.Client) {
	org := os.Getenv("INFLUXDB_ORG")

	// Get query client
	queryAPI := client.QueryAPI(org)

	// Query
	fluxQuery := fmt.Sprintf(`from(bucket: "users_business_events")
    |> range(start: -1h)
    |> filter(fn: (r) => r["_measurement"] == "thermostat")
    |> yield(name: "mean")`)

	result, err := queryAPI.QueryRaw(context.Background(), fluxQuery, influxdb2.DefaultDialect())
	if err == nil {
		fmt.Println("QueryResult:")
		fmt.Println(result)
	} else {
		panic(err)
	}
}

func read_events_as_query_table_result(client influxdb2.Client) map[time.Time]ThermostatSetting {
	org := os.Getenv("INFLUXDB_ORG")

	// Get query client
	queryAPI := client.QueryAPI(org)

	// Query. You need to change a bit the Query from the Query Builder
	// Otherwise it won't work
	fluxQuery := fmt.Sprintf(`from(bucket: "users_business_events")
    |> range(start: -1h)
    |> filter(fn: (r) => r["_measurement"] == "thermostat")
    |> yield(name: "mean")`)

	result, err := queryAPI.Query(context.Background(), fluxQuery)

	// Putting back the data in share requires a bit of work
	var resultPoints map[time.Time]ThermostatSetting
	resultPoints = make(map[time.Time]ThermostatSetting)

	if err == nil {
		// Iterate over query response
		for result.Next() {
			// Notice when group key has changed
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}

			val, ok := resultPoints[result.Record().Time()]

			if !ok {
				val = ThermostatSetting{
					user: fmt.Sprintf("%v", result.Record().ValueByKey("user")),
				}
			}

			switch field := result.Record().Field(); field {
			case "avg":
				val.avg = result.Record().Value().(float64)
			case "max":
				val.max = result.Record().Value().(float64)
			default:
				fmt.Printf("unrecognized field %s.\n", field)
			}

			resultPoints[result.Record().Time()] = val

		}
		// check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
		}
	} else {
		panic(err)
	}

	return resultPoints

}

func write_event_with_line_protocol(client influxdb2.Client, t ThermostatSetting) {
	bucket := os.Getenv("INFLUXDB_BUCKET")
	org := os.Getenv("INFLUXDB_ORG")

	// get non-blocking write client
	writeAPI := client.WriteAPI(org, bucket)
	// write line protocol
	writeAPI.WriteRecord(fmt.Sprintf("thermostat,unit=temperature,user=%s avg=%f,max=%f", t.user, t.avg, t.max))
	// Flush writes
	writeAPI.Flush()
}

func write_event_with_params_constror(client influxdb2.Client, t ThermostatSetting) {
	bucket := os.Getenv("INFLUXDB_BUCKET")
	org := os.Getenv("INFLUXDB_ORG")

	// Use blocking write client for writes to desired bucket
	writeAPI := client.WriteAPI(org, bucket)
	// Create point using full params constructor
	p := influxdb2.NewPoint("thermostat",
		map[string]string{"unit": "temperature", "user": t.user},
		map[string]interface{}{"avg": t.avg, "max": t.max},
		time.Now())
	writeAPI.WritePoint(p)
	// Flush writes
	writeAPI.Flush()
}

func write_event_with_fluent_Style(client influxdb2.Client, t ThermostatSetting) {
	bucket := os.Getenv("INFLUXDB_BUCKET")
	org := os.Getenv("INFLUXDB_ORG")

	// Use blocking write client for writes to desired bucket
	writeAPI := client.WriteAPI(org, bucket)
	// create point using fluent style
	p := influxdb2.NewPointWithMeasurement("thermostat").
		AddTag("unit", "temperature").
		AddTag("user", t.user).
		AddField("avg", t.avg).
		AddField("max", t.max).
		SetTime(time.Now())
	writeAPI.WritePoint(p)
	// Flush writes
	writeAPI.Flush()
}

func write_event_with_blocking_write(client influxdb2.Client) {
	bucket := os.Getenv("INFLUXDB_BUCKET")
	org := os.Getenv("INFLUXDB_ORG")

	// Get blocking write client
	writeAPI := client.WriteAPIBlocking(org, bucket)

	// write line protocol
	writeAPI.WriteRecord(context.Background(), fmt.Sprintf("stat,unit=temperature1 avg=%f,max=%f", 23.5, 45.0))
}

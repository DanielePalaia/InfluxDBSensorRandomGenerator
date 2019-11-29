package main

import (
	_ "github.com/influxdata/influxdb1-client"
)

func main() {

	influx := MakeNewInflux("http://localhost:8086", "", "", "TEST")

	influx.Connect()

	influx.InsertRandomCondizionatori(5)
	//influx.Query("select value, color::tag from a_year.fromgolangclient")

	/*for i, r := range results {

		log.Println("values: ", i, r)
	}*/

}

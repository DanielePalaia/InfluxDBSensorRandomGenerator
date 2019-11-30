package main

import (
	"sync"

	_ "github.com/influxdata/influxdb1-client"
)

func main() {

	influx := MakeNewInflux("http://localhost:8086", "", "", "TEST")
	influx.Connect()

	scenario := MakeNewScenario(influx)

	var wg sync.WaitGroup
	wg.Add(2)

	go scenario.InsertRandomEnergyStations(15)
	go scenario.InsertRandomPowerMeters(15)
	go scenario.InsertRandomCondizionatori(15)
	go scenario.InsertRandomEnvironmentSensors(15)

	wg.Wait()

}

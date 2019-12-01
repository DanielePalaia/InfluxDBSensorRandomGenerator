package main

import (
	"log"
	"strconv"
	"sync"

	_ "github.com/influxdata/influxdb1-client"
)

func main() {

	/** Read configuration file properties.ini*/
	/* Reading properties from ./properties.ini */
	prop, _ := ReadPropertiesFile("./properties.ini")
	influxURL, _ := prop["influxURL"]
	influxUsername := prop["influxUsername"]
	influxPassword := prop["influxPassword"]
	influxDatabase := prop["database"]
	numsiti, err := strconv.Atoi(prop["numsiti"])
	if err != nil {
		log.Fatalf("error with pollling value")
	}
	influxPolling, err := strconv.Atoi(prop["polling"])
	if err != nil {
		log.Fatalf("error with pollling value")
	}

	log.Println("starting generator for influxdb: " + influxURL + " database: " + influxDatabase + " every " + prop["polling"] + " seconds")

	influx := MakeNewInflux(influxURL, influxUsername, influxPassword, influxDatabase)
	influx.Connect()

	scenario := MakeNewScenario(influx, numsiti, influxPolling)

	var wg sync.WaitGroup
	wg.Add(4)

	go scenario.InsertRandomEnergyStations()
	go scenario.InsertRandomPowerMeters()
	go scenario.InsertRandomCondizionatori()
	go scenario.InsertRandomEnvironmentSensors()

	wg.Wait()

}

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
	prop, err := ReadPropertiesFile("./properties.ini")
	if err != nil {
		log.Fatalf("error trying to open file properties.ini")
	}
	influxURL, _ := prop["influxURL"]
	influxUsername := prop["influxUsername"]
	influxPassword := prop["influxPassword"]
	influxDatabase := prop["database"]

	numApparati, err := strconv.Atoi(prop["numApparati"])
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

	scenario := MakeNewScenario(influx, influxPolling, numApparati)

	var wg sync.WaitGroup
	wg.Add(4)

	appliances, err := ReadPropertiesFile("./config/type_appliances.ini")
	if err != nil {
		log.Fatalf("error trying to open file ./config/type_appliaces.ini")
	}

	for k, v := range appliances {
		go scenario.InsertAppliances(k, v)
	}

	wg.Wait()

}

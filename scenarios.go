package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	_ "github.com/influxdata/influxdb1-client"
	client "github.com/influxdata/influxdb1-client/v2"
)

/* Tye influx (it manages the connections) */
type scenarios struct {
	influxdb *influx
	polling  int
	sites    map[string]string
}

func MakeNewScenario(influxdb *influx, polling int) *scenarios {

	var err error
	myScenario := new(scenarios)
	myScenario.influxdb = influxdb
	myScenario.polling = polling
	myScenario.sites, err = ReadPropertiesFile("./config/site_names.ini")
	if err != nil {
		log.Fatalf("error trying to open property file: ./config/site_names.ini")
	}

	return myScenario

}

func (scenario *scenarios) InsertAppliances(typeApp string, fileName string) {

	for {

		for k, _ := range scenario.sites {
			site := k + ".1." + typeApp

			scenario.InsertFromFile(site, fileName)
		}
		time.Sleep(time.Duration(scenario.polling) * time.Second)
	}

}

func (scenario *scenarios) InsertFromFile(site string, fileName string) {
	properties, err := ReadPropertiesFile(fileName)
	if err != nil {
		log.Fatalf("error trying to open property file: " + fileName)
	}
	mapValue := make(map[string]interface{})

	for k, v := range properties {
		newv, _ := strconv.Atoi(v)
		mapValue[k] = rand.Int31n(int32(newv))

	}

	var (
		sampleSize = 1
		pts        = make([]*client.Point, sampleSize)
	)

	for i := 0; i < sampleSize; i++ {
		pts[i], _ = client.NewPoint(
			site,
			nil,
			mapValue,
			time.Now())

	}

	batchPointConfig := client.BatchPointsConfig{

		Database: scenario.influxdb.dbName,
		//RetentionPolicy: "a_year",
	}

	bps, err := client.NewBatchPoints(batchPointConfig)

	if err != nil {
		log.Fatal("error inserting rows into influx database %v", err)
	}

	bps.AddPoints(pts)

	err = scenario.influxdb.conn.Write(bps)
	if err != nil {
		log.Println("Insert data error:")
		log.Fatal(err)
	}
}

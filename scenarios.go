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
	influxdb    *influx
	polling     int
	sites       map[string]string
	numApparati int
}

func MakeNewScenario(influxdb *influx, polling int, numApparati int) *scenarios {

	var err error
	myScenario := new(scenarios)
	myScenario.influxdb = influxdb
	myScenario.polling = polling
	myScenario.sites, err = ReadPropertiesFile("./config/site_names.ini")
	myScenario.numApparati = numApparati
	if err != nil {
		log.Fatalf("error trying to open property file: ./config/site_names.ini")
	}

	return myScenario

}

func (scenario *scenarios) InsertAppliances(typeApp string, fileName string) {

	numApparati := scenario.numApparati
	for {

		for k, _ := range scenario.sites {

			if typeApp == "EnergyStation" {
				numApparati = 1
			} else {
				numApparati = scenario.numApparati
			}

			for i := 0; i < numApparati; i++ {
				site := k + "_1_" + typeApp + strconv.Itoa(i)
				scenario.InsertFromFile(site, fileName)
			}
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

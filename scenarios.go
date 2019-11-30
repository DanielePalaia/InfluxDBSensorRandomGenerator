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
}

func MakeNewScenario(influxdb *influx) *scenarios {

	myScenario := new(scenarios)
	myScenario.influxdb = influxdb

	return myScenario

}

func (scenario *scenarios) InsertRandomCondizionatori(seconds int) {

	for {

		for i := 0; i < 10; i++ {
			sito := "sito" + strconv.Itoa(i) + ".1." + "condizionatore"

			scenario.InsertCondizionatore(sito)

		}
		time.Sleep(10 * time.Second)
	}

}

func (senario *scenarios) InsertRandomPowerMeters(seconds int) {

	for {

		for i := 0; i < 10; i++ {
			sito := "sito" + strconv.Itoa(i) + ".1." + "powermeter"

			senario.InsertPowerMeter(sito)

		}
		time.Sleep(10 * time.Second)
	}

}

func (scenario *scenarios) InsertPowerMeter(measure string) {
	var (
		sampleSize = 1
		pts        = make([]*client.Point, sampleSize)
	)

	EAttNeg := 50 + rand.Int31n(1000)
	EAttPos := 50 + rand.Int31n(1000)
	P := 50 + rand.Int31n(1000)
	P1 := 50 + rand.Int31n(1000)
	P2 := 50 + rand.Int31n(1000)
	P3 := 50 + rand.Int31n(1000)
	Pavg := 50 + rand.Int31n(1000)

	for i := 0; i < sampleSize; i++ {
		pts[i], _ = client.NewPoint(
			measure,
			nil,
			map[string]interface{}{
				"Energia attiva erogata":   EAttNeg,
				"Energia attiva assorbita": EAttPos,
				"Potenza":                  P,
				"Potenza Attiva L1":        P1,
				"Potenza attiva L2":        P2,
				"Potenza attiva L3":        P3,
				"Potenza media":            Pavg,
			},
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

func (scenario *scenarios) InsertRandomEnergyStations(seconds int) {

	for {

		for i := 0; i < 10; i++ {
			sito := "sito" + strconv.Itoa(i) + ".1." + "energystation"

			scenario.InsertEnergyStation(sito)

		}
		time.Sleep(10 * time.Second)
	}

}

func (scenario *scenarios) InsertEnergyStation(measure string) {
	var (
		sampleSize = 1
		pts        = make([]*client.Point, sampleSize)
	)

	for i := 0; i < sampleSize; i++ {
		pts[i], _ = client.NewPoint(
			measure,
			nil,
			map[string]interface{}{
				"Anomalia stazione di energia": rand.Int31n(2),
				"Grande Allarme":               rand.Int31n(2),
				"Grande allarme elettrico":     rand.Int31n(2),
				"Allarme porta":                rand.Int31n(2),
				"Allarme incendio":             rand.Int31n(2),
				"Allarme alta temperatura":     rand.Int31n(2),
				"Piccolo allarme":              rand.Int31n(2),
				"Minimo batteria":              rand.Int31n(2),
				"Allarme Fase Rete S.E.":       rand.Int31n(2),
				"Pressurizzazione":             rand.Int31n(100),
				"Corrente batterie":            rand.Int31n(10),
				"Tensione batterie":            rand.Int31n(10),
				"Tensione AC Linea A mV":       rand.Int31n(10),
				"Tensione AC Linea B mV":       rand.Int31n(10),
				"Tensione AC Linea C mV":       rand.Int31n(10),
				"Numero di moduli SM AC":       rand.Int31n(10),
				"Numero di moduli SM BAT":      rand.Int31n(10),
				"Numero di moduli SM IO":       rand.Int31n(10),
				"Sonda di temperatura 1 °C":    rand.Int31n(40),
				"Allarme SOV Allarme":          rand.Int31n(2),
				"Corrente di sistema mA":       rand.Int31n(40),
				"Capacità del sistema":         rand.Int31n(40),
				"Tensione di sistema":          rand.Int31n(40),
			},
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

func (scenario *scenarios) InsertRandomEnvironmentSensors(seconds int) {

	for {

		for i := 0; i < 10; i++ {
			sito := "sito" + strconv.Itoa(i) + ".1." + "sensor"

			scenario.InsertEnvironmentSensor(sito)

		}
		time.Sleep(10 * time.Second)
	}

}

func (scenario *scenarios) InsertEnvironmentSensor(measure string) {
	var (
		sampleSize = 1
		pts        = make([]*client.Point, sampleSize)
	)

	randomHumidity := 50 + rand.Int31n(80)
	windSpeed := rand.Int31n(2000)
	randomTemperatureOutDoor := rand.Int31n(40)
	randomTemperatureIndoor := 10 + rand.Int31n(40)

	for i := 0; i < sampleSize; i++ {
		pts[i], _ = client.NewPoint(
			measure,
			nil,
			map[string]interface{}{
				"Umidity":            randomHumidity,
				"IndoorTemperature":  randomTemperatureIndoor,
				"OutdoorTemperature": randomTemperatureOutDoor,
				"WindSpeed":          windSpeed,
			},
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

func (scenario *scenarios) InsertCondizionatore(measure string) {
	var (
		sampleSize = 1
		pts        = make([]*client.Point, sampleSize)
	)

	randomHumidity := 50 + rand.Int31n(80)
	randomPreassure := 10 + rand.Int31n(30)
	randomTemperatureOutDoor := rand.Int31n(40)
	randomTemperatureIndoor := 10 + rand.Int31n(40)

	for i := 0; i < sampleSize; i++ {
		pts[i], _ = client.NewPoint(
			measure,
			nil,
			map[string]interface{}{
				"Umidity":            randomHumidity,
				"Preassure":          randomPreassure,
				"IndoorTemperature":  randomTemperatureIndoor,
				"OutdoorTemperature": randomTemperatureOutDoor,
				"ExitTemperature":    randomTemperatureIndoor,
			},
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

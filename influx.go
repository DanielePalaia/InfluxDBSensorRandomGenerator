package main

import (
	"fmt"
	"log"

	_ "github.com/influxdata/influxdb1-client"
	client "github.com/influxdata/influxdb1-client/v2"
)

/* Tye influx (it manages the connections) */
type influx struct {
	url    string
	user   string
	passwd string
	dbName string
	conn   client.Client
}

func MakeNewInflux(url string, user string, passwd string, dbName string) *influx {

	myInflux := new(influx)
	myInflux.url = url
	myInflux.user = user
	myInflux.passwd = passwd
	myInflux.dbName = dbName

	return myInflux

}

/* Connecting to influxdb server */
func (infl *influx) Connect() {

	var err error
	infl.conn, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     infl.url,
		Username: infl.user,
		Password: infl.passwd,
	})
	if err != nil {
		log.Fatalf("error connecting to influxdb server %v", err)
	}

}

/* Create a new database */
func (infl *influx) Create(dbName string) {
	// Workaround, since daocloud influxdb haven't privision an instance
	// create the db instance here
	q := client.Query{
		Command:  fmt.Sprintf("create database %s", dbName),
		Database: dbName,
	}

	// ignore the error of existing database
	_, err := infl.conn.Query(q)
	if err == nil {
		log.Fatalf("error creating a new database %v", err)
	}
}

func (infl *influx) Query(query string) [][]interface{} {
	q := client.Query{
		Command:  query,
		Database: infl.dbName,
	}

	response, err := infl.conn.Query(q)
	if err != nil {
		log.Println("Error, ", err)
		return nil
	}

	if response.Error() != nil {
		log.Println("Response error, ", response.Error())
		return nil
	}

	result := response.Results[0]
	if result.Err != "" {
		log.Println("Result error, ", result.Err)
		return nil
	}

	serie := result.Series[0]

	for k, v := range serie.Tags {

		log.Println("k, v" + k + v)
	}

	return serie.Values
}

/*
func (infl *influx) Insert() {
	var (
		shapes     = []string{"circle", "rectangle", "square", "triangle"}
		colors     = []string{"red", "blue", "green", "yellow"}
		sampleSize = 4
		pts        = make([]*client.Point, sampleSize)
	)

	for i := 0; i < sampleSize; i++ {
		pts[i], _ = client.NewPoint(
			"fromgolangclient",
			map[string]string{
				"color": colors[i],
				"shape": shapes[i],
			},
			map[string]interface{}{
				"value": i,
			},
			time.Now())

	}

	batchPointConfig := client.BatchPointsConfig{

		Database:        infl.dbName,
		RetentionPolicy: "a_year",
	}

	bps, err := client.NewBatchPoints(batchPointConfig)

	if err != nil {
		log.Fatal("error inserting rows into influx database %v", err)
	}

	bps.AddPoints(pts)

	err = infl.conn.Write(bps)
	if err != nil {
		log.Println("Insert data error:")
		log.Fatal(err)
	}
}*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	 "github.com/influxdata/influxdb1-client"
//	client "github.com/influxdata/influxdb1-client/v2"
)

var ic *client.Client

var host string
var port string
var db string
var user string
var password string


func InitInfluxdb() {
	host = "influxdb"
	port = "8086"
	db = "test"
	user = "test"
	password = "test"
	u, err := url.Parse(fmt.Sprintf("http://%s:%s", host, port))
	if err != nil {
		log.Fatal(err)
	}

	ic, err = client.NewClient(client.Config{URL: *u})
	if err != nil {
		log.Fatal(err)
	}

	if _, _, err := ic.Ping(); err != nil {
		log.Fatal(err)
	}
	ic.SetAuth(user, password)
	q := client.Query{
		Command:  fmt.Sprintf("create database %s", db),
		Database: db,
	}
	fmt.Println("runnign the DB query!!!")
	ic.Query(q)
	fmt.Println("Initilization is Finished!!")
}

func StoreDataInfluxdb(count int) error {
	startTime := time.Now()
	pts := make([]client.Point, count)
	for i := 0; i < count; i++ {
		data := GenerateData()
		pts[i] = client.Point{
			Measurement: "shapes",
			Fields: map[string]interface{}{
				"AccountID": data.AccountID,
				"Name": data.Name,
				"FullName": data.FullName,
				"ProductName": data.ProductName,
				"Email": data.Email,
				"EmailSubject": data.EmailSubject,
				"EmailBody": data.EmailBody,
				"UserAgent": data.UserAgent,
				"Company": data.Company,
				"DomainName": data.DomainName,
				"Gender": data.Gender,
				"Language": data.Language,
				"value":i,
			},
			Time: time.Now(),
		}
	}

	bps := client.BatchPoints{
		Points:          pts,
		Database:        db,
	}
	_, err := ic.Write(bps)
	if err != nil {
		log.Println("Insert data error:")
		log.Fatal(err)
		return err
	}
	endTime := time.Now()
	diff := endTime.Sub(startTime).Seconds()
	fmt.Println("Write Operation Finished for Count::" + "   " + string(count) + "   " + "in Following Seconds")
	fmt.Println("*************")
	fmt.Println(diff)
	fmt.Println("*************")
	return nil
}

func query() map[string]string {
	q := client.Query{
		Command:  "select * from shapes where value = 0",
		Database: db,
	}

	response, err := ic.Query(q)
	if err != nil {
		log.Println("Error, ", err)
		return nil
	}

	if response.Error() != nil {
		log.Println("Response error, ", response.Error())
		return nil
	}

	result := response.Results[0]
	if result.Err != nil {
		log.Println("Result error, ", result.Err)
		return nil
	}

	serie := result.Series[0]
	// if serie.Err != nil {
	// 	log.Println("Serie error, ", serie.Err)
	// 	return nil
	// }

	return serie.Tags
}

func hello(res http.ResponseWriter, req *http.Request) {
	m := query()
	res.Write([]byte(fmt.Sprintf("The first shape is %s %s!", m["color"], m["shape"])))
}

//
// func main() {
	// host = "influx"
	// port = "8086"
	// db = "test"
	// user = "test"
	// password = "test"
	//
	// // workaround, daocloud influxdb have not privision db instance
	// if len(db) == 0 {
	// 	db = "mydb"
	// }
	//
	// connect()
	// log.Println("Successfully connect to influxdb ...")
	//
	// // prepare data
	// create()
	// insert()
	//
	// http.HandleFunc("/", hello)
	//
	// log.Println("Start listening...")
	// if err := http.ListenAndServe(":80", nil); err != nil {
	// 	log.Fatal(err)
	// }
// }

package main

import (
	"encoding/csv"
	"fmt"
	//"log"
	//	"net/http"
	"os"
	"strconv"
	"time"
	//"strings"
)

func main() {

	// prepare data
	// create()
	// insert()

	// go build -o main main.go
	//  go run main.go Cassandra write

	// main will take arguents, based on that It will run various test, DBType=Cassandra, TestType=Read
	//  StoreDataCassandra()

	// flag.StringVar(&DBType,"DBType","","Enter your DB Type")
	// flag.StringVar(&TestName,"TestName","","Enter Your TestName")
	// flag.Parse()
	//
	//	InitIgnite()
	if os.Args[1] == "Cassandra" && os.Args[2] == "Write" {
		fmt.Println(os.Args[1], os.Args[2])
		fmt.Println("Starting Writing Operation For Cassandra")
		InitCassandra()
		count, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println("Cassandra Successfuly Initilized!")
		var timeTaken float64
		var totalTaken []float64
		var CSVData []string
		startTime := time.Now()
		sum := float64(0)
		countTest := 1000
		file, err := os.Create("/etc/test/csv/result_cassandra.csv")
		if err != nil {
			fmt.Println("error", err)
		}
		defer file.Close()
		writer := csv.NewWriter(file)
		defer writer.Flush()
		for queryCount := 0; queryCount < count; queryCount++ {
			timeTaken, err = StoreDataCassandra()
			if err != nil {
				fmt.Println("error occoured!!", err)
			}
			totalTaken = append(totalTaken, timeTaken)
			if queryCount%countTest == 0 {
				//fmt.Println("queryCount%countTest", queryCount%countTest) Should be always Zero
				countTest := countTest + 1000
				fmt.Println("this is count Test", countTest)
				for _, value := range totalTaken {
					sum = sum + value
				}
				value1 := fmt.Sprintf("%f", sum)
				value2 := strconv.Itoa(queryCount)
				CSVData = append(CSVData, value2, value1)
				//	CSVData = append(CSVData,
				err2 := writer.Write(CSVData)
				if err2 != nil {
					fmt.Println("error occoured in CSV", err)
				}
				fmt.Println("this is query count", queryCount)
				fmt.Println("Number of operations", len(totalTaken))
				fmt.Println("Time Taken", sum)
				fmt.Println("CSVData*********", CSVData)
				totalTaken = nil
				sum = 0.0
				CSVData = nil
			}
		}
		totalDiff := time.Now().Sub(startTime).Seconds()
		fmt.Println(totalDiff)
	} else if os.Args[1] == "Cassandra" && os.Args[2] == "WriteShoot" {

		fmt.Println(os.Args[1], os.Args[2])
		fmt.Println("Starting Writing Operation For Cassandra")
		InitCassandra()
		for {
			_, err := StoreDataCassandra()
			if err != nil {
				fmt.Println("error occoured!!", err)
			}
		}
	} else if os.Args[1] == "Influxdb" && os.Args[2] == "Write" {
		fmt.Println(os.Args[1], os.Args[2])
		fmt.Println("Starting Writing Operation For Influxdb")
		InitInfluxdb()
		count, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println("Timescaledb Initilized successfully")
		StoreDataInfluxdb(count)
	} else if os.Args[1] == "Timescaledb" && os.Args[2] == "Write" {
		fmt.Println(os.Args[1], os.Args[2])
		fmt.Println("Starting Writing Operation For TimeScaledb")
		InitTimeScale()
		count, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println("Timescaledb Initilized successfully")
		StoreDataTimeScaleDB(count)
	} else if os.Args[1] == "Scylladb" && os.Args[2] == "Write" {
		fmt.Println(os.Args[1], os.Args[2])
		fmt.Println("Starting Writing Operation For Scylladb")
		InitScyllaDB()
		count, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println("Scylladb Successfuly Initilized!")
		StoreDataSycllaDB(count)
	} else if os.Args[1] == "Scylladb" && os.Args[2] == "Read" {
		fmt.Println(os.Args[1], os.Args[2])
		fmt.Println("Starting Writing Operation For Scylladb")
		InitScyllaDB()
		count, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println("Scylladb Successfuly Initilized!")
		FetchDataSycllaDB(count)
	} else if os.Args[1] == "Influxdb" && os.Args[2] == "Read" {
		fmt.Println(os.Args[1], os.Args[2])
		fmt.Println("Starting Reading Operation For InfluxDB")
		InitInfluxdb()
		count, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println("Influx Successfuly Initilized!")
		FetchDataInflux(count)
	} else if os.Args[1] == "Timescaledb" && os.Args[2] == "Read" {
		fmt.Println(os.Args[1], os.Args[2])
		fmt.Println("Starting Writing Operation For Scylladb")
		InitTimeScale()
		count, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println("Timescaledb Successfuly Initilized!")
		FetchDataTimescaledb(count)
	} else if os.Args[1] == "Timescaledb" && os.Args[2] == "ReadMultiple" {
		fmt.Println(os.Args[1], os.Args[2])
		fmt.Println("Starting Writing Operation For Scylladb")
		InitTimeScale()
		count, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println("Timescaledb Successfuly Initilized!")
		startTime := time.Now()
		for i := 0; i < 100000; i++ {
			fmt.Println("__hello")
			FetchDataTimescaledb(count)
		}
		endTime := time.Now()
		diff := endTime.Sub(startTime).Seconds()
		fmt.Println("Multiple READ Opeartion Finised in Following Seconds")
		fmt.Println("*************")
		fmt.Println(diff)
		fmt.Println("*************")
	} else if os.Args[1] == "Influxdb" && os.Args[2] == "ReadMultiple" {
		fmt.Println(os.Args[1], os.Args[2])
		fmt.Println("Starting Reading Operation For InfluxDB")
		InitInfluxdb()
		count, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println("Influxdb Successfuly Initilized!")
		startTime := time.Now()
		for i := 0; i < 1000; i++ {
			fmt.Println("__hello")
			FetchDataInflux(count)
		}
		endTime := time.Now()
		diff := endTime.Sub(startTime).Seconds()
		fmt.Println("Multiple READ Opeartion Finised in Following Seconds")
		fmt.Println("*************")
		fmt.Println(diff)
		fmt.Println("*************")

	} else if os.Args[1] == "Cassandra" && os.Args[2] == "Read" {
		fmt.Println(os.Args[1], os.Args[2])
		fmt.Println("Starting Read Operation For Cassandra")
		InitCassandra()
		count, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println("Cassandr Successfuly Initilized!")
		FetchDataCassandra(count)
	} else if os.Args[1] == "Cassandra" && os.Args[2] == "ReadComplex" {
		fmt.Println(os.Args[1], os.Args[2])
		fmt.Println("Starting Read Operation For Cassandra")
		InitCassandra()
		count, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println("Cassandr Successfuly Initilized!")
		FetchDataCassandraComplex(count)
		//FetchDataCassandra(count)
	} else if os.Args[1] == "Cassandra" && os.Args[2] == "ReadComplexMultiple" {
		fmt.Println(os.Args[1], os.Args[2])
		fmt.Println("Starting Read Operation For Cassandra")
		InitCassandra()
		count, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		startTime := time.Now()
		fmt.Println("Cassandr Successfuly Initilized!")
		for i := 0; i < 100000; i++ {
			fmt.Println("hello____________")
			FetchDataCassandraComplex(count)
		}
		endTime := time.Now()
		diff := endTime.Sub(startTime).Seconds()
		fmt.Println("Multiple READ Opeartion Finised in Following Seconds")
		fmt.Println("*************")
		fmt.Println(diff)
		fmt.Println("*************")
		//FetchDataCassandra(count)
	} else if os.Args[1] == "Cassandra" && os.Args[2] == "ReadMultiple" {
		fmt.Println(os.Args[1], os.Args[2])
		fmt.Println("Starting Read Operation For Cassandra")
		InitCassandra()
		count, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println(count)
		fmt.Println("Cassandr Successfuly Initilized!")
		startTime := time.Now()
		fmt.Println("Multiple Operatin Started")
		for i := 0; i < 100000; i++ {
			fmt.Println("hello____________")
			FetchDataCassandra(100)
		}
		endTime := time.Now()
		diff := endTime.Sub(startTime).Seconds()
		fmt.Println("Multiple READ Opeartion Finised in Following Seconds")
		fmt.Println("*************")
		fmt.Println(diff)
		fmt.Println("*************")

	} else if os.Args[1] == "Scylladb" && os.Args[2] == "ReadMultiple" {
		fmt.Println(os.Args[1], os.Args[2])
		fmt.Println("Starting Read Operation For Cassandra")
		InitScyllaDB()
		count, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println(count)
		fmt.Println("Cassandr Successfuly Initilized!")
		startTime := time.Now()
		fmt.Println("Multiple Operatin Started")
		for i := 0; i < 100000; i++ {
			fmt.Println("hello____________")
			FetchDataSycllaDB(count)
		}
		endTime := time.Now()
		diff := endTime.Sub(startTime).Seconds()
		fmt.Println("Multiple READ Opeartion Finised in Following Seconds")
		fmt.Println("*************")
		fmt.Println(diff)
		fmt.Println("*************")

	} else if os.Args[1] == "Cassandra" && os.Args[2] == "Read" && os.Args[4] == "Complex" {
		fmt.Println(os.Args[1], os.Args[2])
		fmt.Println("Starting Read Operation For Cassandra")
		InitCassandra()
		count, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println("Cassandr Successfuly Initilized!")
		FetchDataCassandraComplex(count)
	} else {

		fmt.Println("No Arguments Passed!!")
		fmt.Println(os.Args[1], os.Args[2])
	}
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	//"strings"
)

func main() {
	// go build -o main main.go
	//  go run main.go Cassandra write

	// main will take arguents, based on that It will run various test, DBType=Cassandra, TestType=Read
	//  StoreDataCassandra()

	// flag.StringVar(&DBType,"DBType","","Enter your DB Type")
	// flag.StringVar(&TestName,"TestName","","Enter Your TestName")
	// flag.Parse()
	//
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
		StoreDataCassandra(count)
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

package main
import (
  "fmt"
  "os"
  "strconv"
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

if os.Args[1]=="Cassandra" && os.Args[2]=="Write"{
  fmt.Println(os.Args[1],os.Args[2])
  fmt.Println("Starting Writing Operation For Cassandra")
  InitCassandra()
  count, err := strconv.Atoi(os.Args[3])
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
  fmt.Println("Cassandra Successfuly Initilized!")
  StoreDataCassandra(count)
} else if os.Args[1]=="Scylladb" && os.Args[2]=="Write" {

  fmt.Println(os.Args[1],os.Args[2])
  fmt.Println("Starting Writing Operation For Scylladb")
  InitScyllaDB()
  count, err := strconv.Atoi(os.Args[3])
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
  fmt.Println("Scylladb Successfuly Initilized!")
  StoreDataSycllaDB(count)
} else if os.Args[1]=="Scylladb" && os.Args[2]=="Read" {
  fmt.Println(os.Args[1],os.Args[2])
  fmt.Println("Starting Writing Operation For Scylladb")
  InitScyllaDB()
  fmt.Println("Scylladb Successfuly Initilized!")
  FetchDataSycllaDB()
}else if os.Args[1]=="Cassandra" && os.Args[2]=="Read" {
  fmt.Println(os.Args[1],os.Args[2])
  fmt.Println("Starting Read Operation For Cassandra")
  InitCassandra()
  fmt.Println("Cassandr Successfuly Initilized!")
  FetchDataCassandra()
}else {
  fmt.Println("No Arguments Passed!!")
  fmt.Println(os.Args[1],os.Args[2])
}
}

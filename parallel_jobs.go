package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// what is structure of task...

type Job struct {
	jobID    int
	randomNo int
}

// What will be produced after processing on Task!!
type JobResult struct {
	job Job
	err error
}

var Jobs = make(chan Job, 20)          // type of channel which will accept Job{}
var Results = make(chan JobResult, 20) //type of channel to accept Result{}

// Now We have to Create Dispatcher which will start all the workers , depends on How many workers we want to start
func CreateWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go StoreDataCassandraWorker(&wg) // Run the worker to as Go Routine!!
	}
	wg.Wait()
	close(Results)
}

func SubmitJobs(noOfJobs int) { // Send the Jobs to Channel so able to Process those, Think of it liKe Worker Pool!!
	for i := 0; i < noOfJobs; i++ {
		randno := rand.Intn(999)
		job := Job{i, randno}
		Jobs <- job
	}
	close(Jobs)
}

func Result(done chan bool) {
	for result := range Results {
		fmt.Println("Here we are getting results of all the Submitted Jobs!!")
		if result.err != nil {
			done <- false
		}
	}
	done <- true
}

// func main() {
// 	startTime := time.Now()
// 	// First We have to Admit Number of Jobs to Pool.
// 	numberofjobs := 100000
// 	go SubmitJobs(numberofjobs)
// 	// Now start workers
// 	done := make(chan bool)
// 	go result(done)
// 	numberofworkers := 10
// 	go CreateWorkerPool(numberofworkers)
// 	<-done
// 	// Jobs submitted and Worker is also started! Now wait for results and see magic
//
// 	endTime := time.Now()
// 	diff := endTime.Sub(startTime)
// 	fmt.Println("Total Time Taken is::", diff)
// }

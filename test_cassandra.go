package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

var (
	session_cassandra *gocql.Session
)

func InitCassandra() {
	var err error
	cluster := gocql.NewCluster("cassandra-seed") //replace PublicIP with the IP addresses used by your cluster.
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	//	cluster.Keyspace = keySpace
	cluster.ConnectTimeout = time.Second * 10
	//cluster.Authenticator = gocql.PasswordAuthenticator{Username: "Username", Password: "Password"} //replace the username and password fields with their real settings.
	session_cassandra, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	if err := session_cassandra.Query("CREATE KEYSPACE IF NOT EXISTS test WITH REPLICATION = {'class' : 'SimpleStrategy','replication_factor' : 3};").Exec(); err != nil {
		log.Println(context.Background(), "Error for creating a keyspace[InitCassandra]:", err)
		panic(err)
	}
	if err := session_cassandra.Query("CREATE TABLE IF NOT EXISTS test.user ( account_id bigint, name text, full_name text, product_name text,email text, email_subject text, email_body text,user_agent text, company text, domain_name text,gender text,language text, created_at timestamp, updated_at timestamp, PRIMARY KEY (account_id));").Exec(); err != nil {
		log.Println(context.Background(), "Error in Creating user  table:", err)
		panic(err)
	}
}

//
// func StoreDataCassandra(count int) error {
// 	startTime := time.Now()
// 	countTest := 1000
// 	for i := 0; i < count; i++ {
// 		countTimeStart := time.Now()
// 		data := GenerateData()
// 		if err := session_cassandra.Query(`INSERT INTO test.user ( account_id, name, full_name,product_name,email,
//      email_subject, email_body,user_agent, company, domain_name,gender,language,
//     created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?,?,?,?,?,?,?,?)`,
// 			data.AccountID, data.Name, data.FullName, data.ProductName, data.Email, data.EmailSubject,
// 			data.EmailBody, data.UserAgent, data.Company, data.DomainName, data.Gender, data.Language, data.CreatedAt, data.Updatedat).Exec(); err != nil {
// 			log.Println("Error! unable to insert data into cassandra", err)
// 			return err
// 		}
// 		if i%countTest == 0 {
// 			countTest := countTest + 1000
// 			fmt.Println(countTest)
// 			countTime := time.Now()
// 			secDiff := countTime.Sub(startTime).Seconds()
// 			fmt.Println("timke taken for Every 1000 operations", secDiff)
// 		}
// 	}
// 	endTime := time.Now()
// 	diff := endTime.Sub(startTime).Seconds()
// 	fmt.Println("Write Operation Finished for Count::" + "   " + string(count) + "   " + "in Following Seconds")
// 	fmt.Println("*************")
// 	fmt.Println(diff)
// 	fmt.Println("*************")
// 	return nil
// }

func StoreDataCassandra() (float64, error) {
	startTime := time.Now()
	//countTimeStart := time.Now()
	data := GenerateData()
	if err := session_cassandra.Query(`INSERT INTO test.user ( account_id, name, full_name,product_name,email,
     email_subject, email_body,user_agent, company, domain_name,gender,language,
    created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?,?,?,?,?,?,?,?)`,
		data.AccountID, data.Name, data.FullName, data.ProductName, data.Email, data.EmailSubject,
		data.EmailBody, data.UserAgent, data.Company, data.DomainName, data.Gender, data.Language, data.CreatedAt, data.Updatedat).Exec(); err != nil {
		log.Println("Error! unable to insert data into cassandra", err)
		return 2.2, err
	}
	endTime := time.Now()
	diff := endTime.Sub(startTime).Seconds()
	return diff, nil
}

func FetchDataCassandra(accountid int) (float64, error){
	startTime := time.Now()
	query := fmt.Sprintf("SELECT account_id, name, full_name,product_name,email,email_subject, email_body,user_agent, company, domain_name,gender,language, created_at, updatedat from test.user_table WHERE account_id = ?")
	iter := session_cassandra.Query(query, accountid).Iter()
	for iter.Scan(&accountid) {
		fmt.Println("account_ID", accountid)
	}
	iter.Close()
	endTime := time.Now()
	diff := endTime.Sub(startTime).Seconds()
	return diff, nil
}

func FetchDataCassandraComplex(accountid int) {
	startTime := time.Now()
	query := fmt.Sprintf("SELECT COUNT(account_id,email,company) from test.user_table WHERE account_id = ? and created_at>? and created_at<? GROUP BY account_id, created_at;")
	created_at := time.Now()
	iter := session_cassandra.Query(query, accountid, created_at, created_at, accountid, created_at).Iter()
	endTime := time.Now()
	diff := endTime.Sub(startTime).Seconds()
	fmt.Println("Read Complex Operation Finished in Following Seconds")
	fmt.Println("*************")
	fmt.Println(diff)
	fmt.Println("*************")
	for iter.Scan(&accountid) {
		fmt.Println("account_ID", accountid)
	}
	iter.Close()
}

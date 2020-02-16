package main

import (
	"context"
	"log"
	"time"
	"github.com/gocql/gocql"
"fmt"
)

var (
	session_cassandra *gocql.Session
)

func InitCassandra() {
	var err error
	cluster := gocql.NewCluster("cassandra") //replace PublicIP with the IP addresses used by your cluster.
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
//	cluster.Keyspace = keySpace
	cluster.ConnectTimeout = time.Second * 10
	//cluster.Authenticator = gocql.PasswordAuthenticator{Username: "Username", Password: "Password"} //replace the username and password fields with their real settings.
	session_cassandra, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	if err := session_cassandra.Query("CREATE KEYSPACE IF NOT EXISTS test WITH REPLICATION = {'class' : 'SimpleStrategy','replication_factor' : 1};").Exec(); err != nil {
		log.Println(context.Background(), "Error for creating a keyspace[InitCassandra]:", err)
		panic(err)
  }
	if err := session_cassandra.Query("CREATE TABLE IF NOT EXISTS test.user ( account_id bigint, name text, full_name text, product_name text,email text, email_subject text, email_body text,user_agent text, company text, domain_name text,gender text,language text, created_at timestamp, updated_at timestamp, PRIMARY KEY (account_id));").Exec(); err != nil {
		log.Println(context.Background(), "Error in Creating user  table:", err)
		panic(err)
	}
}

func StoreDataCassandra(count int) error {
	startTime:=time.Now()
	for i:=0;i<count;i++{
  data:=GenerateData()
  if err := session_cassandra.Query(`INSERT INTO test.user ( account_id, name, full_name,product_name,email,
     email_subject, email_body,user_agent, company, domain_name,gender,language,
    created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?,?,?,?,?,?,?,?)`,
		data.AccountID, data.Name,data.FullName, data.ProductName, data.Email, data.EmailSubject,
    data.EmailBody, data.UserAgent,data.Company,data.DomainName,data.Gender,data.Language,data.CreatedAt,data.Updatedat).Exec(); err != nil {
		log.Println("Error! unable to insert data into cassandra", err)
		return err
  }
}
  endTime:=time.Now()
	diff:=endTime.Sub(startTime).Seconds()
	fmt.Println("Write Operation Finished for Count::"+"   "+string(count)+"   "+"in Following Seconds")
  fmt.Println("*************")
	fmt.Println(diff)
	fmt.Println("*************")
  return nil
}

func FetchDataCassandra(accountid int) {
	startTime:=time.Now()
	query := fmt.Sprintf("SELECT account_id, name, full_name,product_name,email,email_subject, email_body,user_agent, company, domain_name,gender,language, created_at, updatedat from test.user_table WHERE account_id = ?")
	iter := session_cassandra.Query(query, accountid).Iter()
	endTime:=time.Now()
	diff:=endTime.Sub(startTime).Seconds()
	fmt.Println("Write Operation Finished in Following Seconds")
  fmt.Println("*************")
	fmt.Println(diff)
	fmt.Println("*************")
	for iter.Scan(&accountid) {
		fmt.Println("account_ID",accountid)
	}
	iter.Close()
}

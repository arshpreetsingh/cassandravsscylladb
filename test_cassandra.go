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
	// cluster.Keyspace = ServiceConfig.Cassandra.KeySpace
	cluster.ConnectTimeout = time.Second * 10
	//cluster.Authenticator = gocql.PasswordAuthenticator{Username: "Username", Password: "Password"} //replace the username and password fields with their real settings.
	session_cassandra, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	keySpace := "test"

	if err := session_cassandra.Query("CREATE KEYSPACE IF NOT EXISTS " + keySpace + " WITH REPLICATION = {'class' : 'SimpleStrategy','replication_factor' : 1};").Exec(); err != nil {
		log.Println(context.Background(), "Error for creating a keyspace[InitCassandra]:", err)
  }
		return

	if err := session_cassandra.Query("CREATE TABLE IF NOT EXISTS " + keySpace + ".user_table(account_id uuid, name text, full_name text,product_name text,email text, email_subject text, email_body string,user_agent string, company string, domain_name string,gender string,language string, created_at timestamp, updatedat timestamp PRIMARY KEY(account_id, email, company))").Exec(); err != nil {
		log.Println(context.Background(), "Error in Creating user bounces table:", err)
		panic(err)
	}
}

func StoreDataCassandra() error {
  data:=GenerateData()
  if err := session_cassandra.Query(`INSERT INTO user_table(account_id, name, full_name,product_name,email,
     email_subject, email_body,user_agent, company, domain_name,gender,language,
    created_at, updatedat) VALUES (?, ?, ?, ?, ?, ?, ?,?,?,?,?,?,?,?)`,
		data.AccountID, data.FullName, data.ProductName, data.Email, data.EmailSubject,
    data.EmailBody, data.UserAgent,data.Company,data.DomainName,data.Gender,data.Language,data.CreatedAt,data.Updatedat).Exec(); err != nil {
		log.Println("Error! unable to insert data into cassandra[StoreBounceData]:", err)
		return err
  }
  return nil

}

func FetchDataCassandra() {
	accountid:="1234"
	query := fmt.Sprintf("SELECT account_id, name, full_name,product_name,email,email_subject, email_body,user_agent, company, domain_name,gender,language, created_at, updatedat from test.user_table WHERE account_id = ?")
	iter := session_cassandra.Query(query, accountid).Iter()
	for iter.Scan(&accountid) {
		fmt.Println("account_ID",accountid)
	}
	iter.Close()
}

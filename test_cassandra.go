package main

import (
	"context"
	"log"
	"time"
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func InitCassandra() {
	var err error
	cluster := gocql.NewCluster("cassandra") //replace PublicIP with the IP addresses used by your cluster.
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	// cluster.Keyspace = ServiceConfig.Cassandra.KeySpace
	cluster.ConnectTimeout = time.Second * 10
	//cluster.Authenticator = gocql.PasswordAuthenticator{Username: "Username", Password: "Password"} //replace the username and password fields with their real settings.
	session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	keySpace := "test"

	if err := session.Query("CREATE KEYSPACE IF NOT EXISTS " + keySpace + " WITH REPLICATION = {'class' : 'SimpleStrategy','replication_factor' : 1};").Exec(); err != nil {
		log.Println(context.Background(), "Error for creating a keyspace[InitCassandra]:", err)
  }
		return

	if err := session.Query("CREATE TABLE IF NOT EXISTS " + keySpace + ".user_table(account_id uuid, name text, full_name text,product_name text,email text, email_subject text, email_body string,user_agent string, company string, domain_name string,gender string,language string, created_at timestamp, updatedat timestamp PRIMARY KEY(account_id, email, company))").Exec(); err != nil {
		log.Println(context.Background(), "Error in Creating user bounces table:", err)
		panic(err)
	}
}

func StoreData() error {
  data:=GenerateData()
  if err := session.Query(`INSERT INTO user_table(account_id, name, full_name,product_name,email,
     email_subject, email_body,user_agent, company, domain_name,gender,language,
    created_at, updatedat) VALUES (?, ?, ?, ?, ?, ?, ?,?,?,?,?,?,?,?)`,
		data.AccountID, data.FullName, data.ProductName, data.Email, data.EmailSubject,
    data.EmailBody, data.UserAgent,data.Company,data.DomainName,data.Gender,data.Language,data.CreatedAt,data.Updatedat).Exec(); err != nil {
		log.Println("Error! unable to insert data into cassandra[StoreBounceData]:", err)
		return err
  }
  return nil

}



//
// func StoreBounceData(jsonData []byte) error {
// 	var data Data
// 	var err error
// 	err = json.Unmarshal(jsonData, &data)
// 	if err != nil {
// 		return err
// 	}
//
// 	if err := session.Query(`INSERT INTO bounces (account_id, email_uuid, email_address, diagnostic, hard_bounce, spreader_ip, created_at) VALUES (?, ?, ?, ?, ?, ?, ?)`,
// 		data.AccountID, data.EmailUUID, data.Rcpt, data.Bounce.Diagnostic, data.Bounce.HardBounce, data.SendingIP, data.TimeStamp).Exec(); err != nil {
// 		log.Println("Error! unable to insert data into cassandra[StoreBounceData]:", err)
// 		return err
// 	}
//
// 	return nil
// }

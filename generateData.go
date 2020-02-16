package main

import (
"github.com/icrowley/fake"
"time"
"math/rand"
)

type TestData struct {
AccountID int
Name string
FullName string
ProductName string
Email string
EmailSubject string
EmailBody string
UserAgent string
Company string
DomainName string
Gender string
Language  string
CreatedAt time.Time
Updatedat time.Time
}

func GenerateData() TestData{
  test_data:=TestData{AccountID:rand.Intn(10000000),Name:fake.FirstName(),FullName:fake.FullName(),
    ProductName:fake.ProductName(),Email:fake.EmailAddress(),EmailSubject:fake.EmailSubject(),
    EmailBody:fake.EmailBody(),UserAgent:fake.UserAgent(),Company:fake.Company(),DomainName:fake.DomainName(),Gender:fake.Gender(),Language:fake.Language(),
    CreatedAt:time.Now(),Updatedat:time.Now()}
    return test_data
}

func GenerateDataBulk(count int) []TestData{
  test_data_bulk:=[]TestData{}
  for i:=0;i<count;i++ {
    test_data:=TestData{AccountID:rand.Intn(10000000),Name:fake.FirstName(),FullName:fake.FullName(),
      ProductName:fake.ProductName(),Email:fake.EmailAddress(),EmailSubject:fake.EmailSubject(),
      EmailBody:fake.EmailBody(),UserAgent:fake.UserAgent(),Company:fake.Company(),DomainName:fake.DomainName(),Gender:fake.Gender(),Language:fake.Language(),
      CreatedAt:time.Now(),Updatedat:time.Now()}
    test_data_bulk=append(test_data_bulk,test_data)
  }
  return test_data_bulk
}
//
// func main() {
//   data:=GenerateData()
//   fmt.Println(data)
//   data_bulk:=GenerateDataBulk(10)
//   fmt.Println(data_bulk)
//
// }

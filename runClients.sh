#!/bin/sh
END=1000
for i in $(seq 1 $END)
do
 docker-compose run -d test_cassandra Cassandra WriteShoot
done
docker-compose run test_cassandra Cassandra Write 500000
docker logs -f test_cassandra

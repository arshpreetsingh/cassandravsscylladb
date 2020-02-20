# cassandravsscylladb
# Build all the Binaries
```
docker-compose build
```
# Run Required test

```
sudo docker-compose run test_cassandra Cassandra Write 100000
sudo docker-compose run test_scylladb Scylladb Write 100000
sudo docker-compose run test_scylladb Scylladb Write 100000000000
sudo docker-compose run test_scylladb Scylladb Read 1000000000000
sudo docker-compose run test_timescaledb Timescaledb Write 1000000000000
docker-compose run test_timescaledb Timescaledb Read 100000
```

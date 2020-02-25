# cassandravsscylladb

# Build Base images for all Go_lang clients
```
docker-compose build base
```
# Build all the Binaries
```
docker-compose build
```
# Run Required test

```
# Say you want to Read Multiple Benchmarks on Cassandra
docker-compose run test_cassandra Cassandra Write 100000
docker-compose run test_cassandra Cassandra WriteShoot
docker-compose run test_cassandra Cassandra Read 100
docker-compose run test_cassandra Cassandra ReadMultiple 100
docker-compose run test_cassandra Cassandra ReadComplex 100
docker-compose run test_influxdb Influxdb Write 100
sudo docker-compose run test_influxdb Influxdb Read 10000
sudo docker-compose run test_influxdb Influxdb ReadMultiple 100

sudo docker-compose run test_scylladb Scylladb Write 100000
sudo docker-compose run test_scylladb Scylladb Write 100000000000
sudo docker-compose run test_scylladb Scylladb Read 1000000000000
sudo docker-compose run test_timescaledb Timescaledb Write 1000000000000
docker-compose run test_timescaledb Timescaledb Read 100000
docker-compose run test_timescaledb Timescaledb ReadMultiple 100
docker-compose run test_timescaledb Scylladb ReadMultiple 100
```

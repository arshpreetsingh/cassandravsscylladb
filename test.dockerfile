FROM cassandravsscylladb_base:latest as builder
Add . /etc/test/
WORKDIR /etc/test
#RUN tail -f /dev/null
# Build Service as Executable:
RUN go get -v
RUN go build -o testService -ldflags="-w -s"
ENTRYPOINT ["./testService"]
#RUN tail -f /dev/null
